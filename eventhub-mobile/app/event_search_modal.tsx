import EventCard from "@/components/EventCard";
import { ADDRESS_NO_PORT } from "@/constants/address";
import { colors } from "@/constants/colors";
import { fonts } from "@/constants/fonts";
import { router } from "expo-router";
import { useCallback, useEffect, useRef, useState } from "react";
import {
  Image,
  SafeAreaView,
  ScrollView,
  StyleSheet,
  Text,
  TextInput,
  TouchableOpacity,
  View,
} from "react-native";

const ModalScreen = () => {
  const [query, setQuery] = useState("");
  const [results, setResults] = useState([]);

  const searchbarRef = useRef<TextInput>(null);

  useEffect(() => {
    if (searchbarRef.current) {
      searchbarRef.current.focus();
    }
  }, []);

  const search = useCallback(async (text: string) => {
    try {
      const res = await fetch(`http://${ADDRESS_NO_PORT}:9200/events/_search`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          query: {
            multi_match: {
              query: text,
              type: "best_fields",
              fields: ["title^3", "category^2", "description"],
              fuzziness: "AUTO",
              operator: "and",
            },
          },
        }),
      });

      const data = await res.json();
      const hits = data.hits?.hits || [];
      setResults(hits);
    } catch (err) {
      console.error("Ошибка поиска:", err);
    }
  }, []);

  const handleChange = (text: string) => {
    setQuery(text);
    if (text.length >= 2) {
      search(text); // Запускаем поиск только если длина текста >= 2
    } else {
      setResults([]); // Очищаем результаты
    }
  };

  return (
    <SafeAreaView style={{ backgroundColor: "#ffffff" }}>
      <View style={{ height: "100%" }}>
        <View style={styles.search}>
          <TextInput
            ref={searchbarRef}
            style={styles.search_field}
            placeholder="Поиск"
            value={query}
            onChangeText={handleChange}
          ></TextInput>
          <TouchableOpacity onPress={() => router.back()}>
            <Text style={styles.cancel_text}>Отменить</Text>
          </TouchableOpacity>
        </View>

        <ScrollView style={styles.results}>
          {results.map((event) => {
            return (
              <View key={event._source.id} style={{ marginBottom: 12 }}>
                <EventCard
                  id={event._source.id}
                  title={event._source.title}
                  category={event._source.category}
                  date={event._source.date}
                  start_time={event._source.start_time}
                  end_time={event._source.end_time}
                  location={event._source.location}
                  isCompleted={event._source.is_completed}
                />
              </View>
            );
          })}
        </ScrollView>
      </View>
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  search: {
    display: "flex",
    flexDirection: "row",
    alignItems: "center",
    gap: 16,
    marginBottom: 8,
    paddingHorizontal: 16,
  },
  search_field: {
    backgroundColor: "#EBEDF2",
    paddingTop: 10,
    paddingBottom: 8,
    paddingHorizontal: 8,
    borderRadius: 8,
    fontFamily: fonts.Montserrat,
    fontSize: 14,
    color: "#898D94",
    fontWeight: 400,
    flexGrow: 1,
  },
  cancel_text: {
    color: colors.primary,
  },
  results: {
    paddingTop: 8,
    paddingHorizontal: 16,
    display: "flex",
    flexDirection: "column",
  },
});

export default ModalScreen;
