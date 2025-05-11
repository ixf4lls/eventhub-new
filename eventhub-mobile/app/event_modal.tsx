import {
  View,
  Text,
  Button,
  Platform,
  Image,
  StyleSheet,
  TouchableOpacity,
  ScrollView,
  Alert,
} from "react-native";
import { StatusBar } from "expo-status-bar";
import { useLocalSearchParams, useRouter } from "expo-router";
import { fonts } from "@/constants/fonts";
import { colors } from "@/constants/colors";
import * as Haptics from "expo-haptics";
import CustomButton from "@/components/CustomButton";
import { ADDRESS } from "@/constants/address";
import { useEffect, useState } from "react";
import { fetchWithToken } from "@/utils/tokenInterceptor";
import { format, parseISO } from "date-fns";
import { ru } from "date-fns/locale";

type Event = {
  id: number;
  title: string;
  description: string;
  category: string;
  is_public: boolean;
  status: string;
  date: string;
  start_time: string;
  end_time: string;
  location: string;
  creator_id: number;
  organization_id: number;
};

type EventResponse = {
  event: Event;
  is_joined: boolean;
  is_creator: boolean;
};

const formatDate = (date: string) => {
  const parsedDate = parseISO(date);
  const formattedDate = format(parsedDate, "d MMMM", {
    locale: ru,
  }).toUpperCase();

  return formattedDate;
};

const formatTime = (time: string) => {
  if (!time) return "00:00";
  return time.slice(0, 5);
};

const ModalScreen = () => {
  const router = useRouter();

  const [event, setEvent] = useState<Event>();
  const [userJoined, setUserJoined] = useState<boolean>(false);
  const [isCreator, setIsCreator] = useState<boolean>(false);

  const { id, isCompleted } = useLocalSearchParams();

  const LoadData = async (id: number) => {
    try {
      const response = (await fetchWithToken(
        "http://" + ADDRESS + "/api/events/" + id,
        {
          method: "GET",
        },
      )) as Response;

      if (!response.ok) {
        if (response.status === 401) {
          router.replace("/(auth)/login");
        }
        throw new Error(`Ошибка запроса: ${response.status}`);
      }

      const data = (await response.json()) as EventResponse;
      setEvent(data.event);
      setUserJoined(data.is_joined);
      setIsCreator(data.is_creator);
    } catch (error) {
      Alert.alert("Ошибка сервера", "Проверьте подключение.");
      Haptics.notificationAsync(Haptics.NotificationFeedbackType.Error);
    }
  };

  useEffect(() => {
    LoadData(Number(id));
    const interval = setInterval(() => {
      LoadData(Number(id));
    }, 30000);

    return () => clearInterval(interval);
  }, []);

  const JoinEvent = async (event_id: number) => {
    try {
      const response = (await fetchWithToken(
        "http://" + ADDRESS + "/api/events/" + event_id + "/join",
        {
          method: "POST",
        },
      )) as Response;

      if (!response.ok) {
        if (response.status === 401) {
          router.replace("/(auth)/login");
        }
        throw new Error(`Ошибка запроса: ${response.status}`);
      }
      Haptics.notificationAsync(Haptics.NotificationFeedbackType.Success);
      setUserJoined(true);
    } catch (error) {
      Alert.alert("Ошибка сервера", "Проверьте подключение.");
      Haptics.notificationAsync(Haptics.NotificationFeedbackType.Error);
    }
  };

  const QuitEvent = async (event_id: number) => {
    try {
      const response = (await fetchWithToken(
        "http://" + ADDRESS + "/api/events/" + event_id + "/quit",
        {
          method: "DELETE",
        },
      )) as Response;

      if (!response.ok) {
        const errorData = await response.json();
        if (response.status === 401) {
          router.replace("/(auth)/login");
        }
        throw new Error(`Ошибка запроса: ${response.status}`);
      }
      Haptics.notificationAsync(Haptics.NotificationFeedbackType.Success);
      setUserJoined(false);
    } catch (error) {
      Alert.alert("Ошибка сервера", "Проверьте подключение.");
      Haptics.notificationAsync(Haptics.NotificationFeedbackType.Error);
    }
  };

  const handleQuitEvent = (event_id: number) => {
    Alert.alert("Вы уверены, что хотите отменить участие?", "", [
      {
        text: "Нет, оставить все как есть",
        onPress: () => {},
        style: "cancel",
      },
      {
        text: "Да, отменить",
        onPress: () => QuitEvent(event_id),
        style: "destructive",
      },
    ]);
  };

  const handleDeleteEvent = (event_id: number) => {
    Alert.alert("Вы уверены, что хотите удалить это мероприятие", "", [
      {
        text: "Нет, оставить все как есть",
        onPress: () => {},
        style: "cancel",
      },
      {
        text: "Да, удалить",
        onPress: () => DeleteEvent(event_id),
        style: "destructive",
      },
    ]);
  };

  const DeleteEvent = async (event_id: number) => {
    try {
      const response = (await fetchWithToken(
        "http://" + ADDRESS + "/api/events/" + event_id + "/delete",
        {
          method: "DELETE",
        },
      )) as Response;

      if (!response.ok) {
        if (response.status === 401) {
          router.replace("/(auth)/login");
        }
        throw new Error(`Ошибка запроса: ${response.status}`);
      }
      Haptics.notificationAsync(Haptics.NotificationFeedbackType.Success);
      Alert.alert("Мероприятие успешно удалено.");
      router.back();
      router.replace("/(root)/(tabs)/home");
    } catch (error) {
      Alert.alert("Ошибка сервера", "Проверьте подключение." + error);
      Haptics.notificationAsync(Haptics.NotificationFeedbackType.Error);
    }
  };

  return (
    <View style={{ flex: 1, backgroundColor: "#ffffff" }}>
      <TouchableOpacity onPress={() => router.back()} style={styles.close}>
        <Image
          source={require("../assets/icons/close.png")}
          style={{ height: 16, width: 16 }}
        />
      </TouchableOpacity>
      <ScrollView style={styles.container}>
        <Image
          source={require("../assets/images/blank_image_modal.png")}
          style={{ height: "50%", width: "100%" }}
        />

        <View style={styles.content}>
          <View style={styles.tags}>
            {isCompleted == "false" ? (
              isCreator ? (
                <View style={styles.creator}>
                  <Text style={styles.creator__text}>ВЫ СОЗДАТЕЛЬ</Text>
                </View>
              ) : userJoined ? (
                <View style={styles.join}>
                  <Text style={styles.join__text}>ВЫ УЧАСТВУЕТЕ</Text>
                </View>
              ) : null
            ) : null}

            <View style={styles.public}>
              <Text style={styles.public__text}>
                {event?.is_public ? "ОТКРЫТОЕ" : "ДЛЯ ОРГАНИЗАЦИИ"}
              </Text>
            </View>
          </View>
          <View style={styles.mainText}>
            <Text style={styles.title}>{event?.title}</Text>
            <Text style={styles.category}>{event?.category}</Text>
            <Text style={styles.description}>{event?.description}</Text>
          </View>

          <View style={styles.date}>
            <View style={styles.date__item}>
              <Text style={styles.item__text}>
                {event?.date ? formatDate(event.date) : ""}
              </Text>
            </View>
            <View style={styles.date__item}>
              <Text style={styles.item__text}>
                {event?.start_time ? formatTime(event?.start_time) : ""} -{" "}
                {event?.end_time ? formatTime(event?.end_time) : ""}
              </Text>
            </View>
          </View>
          <View style={styles.location}>
            <Image
              source={require("../assets/icons/location.png")}
              style={{ height: 18, width: 14 }}
            />
            <Text style={styles.location__text}>{event?.location}</Text>
          </View>
        </View>

        <Text>{event?.creator_id}</Text>
      </ScrollView>

      {isCompleted == "false" ? (
        isCreator ? (
          <View style={styles.edit_buttons}>
            <View style={{ flexGrow: 3 }}>
              <CustomButton
                onPress={() => {
                  if (!event) return;
                  router.push({
                    pathname: "/event_edit_modal",
                    params: {
                      org_id: event.organization_id,
                      event_id: event.id,
                    },
                  });
                }}
                title="Изменить данные"
                type={"action"}
                fill={"bordered"}
              />
            </View>
            <View style={{ flexGrow: 1 }}>
              <CustomButton
                onPress={() => handleDeleteEvent(event?.id || 0)}
                title="Удалить"
                type={"cancel"}
                fill={"bordered"}
              />
            </View>
          </View>
        ) : (
          <View style={styles.button}>
            {userJoined ? (
              <CustomButton
                onPress={() => handleQuitEvent(Number(id))}
                title="Отменить участие"
                type={"cancel"}
                fill={"bordered"}
              />
            ) : (
              <CustomButton
                onPress={() => JoinEvent(Number(id))}
                title="Я иду!"
                type={"action"}
                fill={"solid"}
              />
            )}
          </View>
        )
      ) : (
        <View style={styles.completed}>
          <Text style={styles.completed__text}>Мероприятие завершено</Text>
        </View>
      )}

      <StatusBar style={Platform.OS === "ios" ? "light" : "auto"} />
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    height: "100%",
    display: "flex",
    width: "100%",
    marginBottom: 32,
  },
  close: {
    position: "absolute",
    top: 16,
    right: 16,
    zIndex: 5,
  },
  content: {
    backgroundColor: "#ffffff",
    height: 555,
    width: "100%",
    paddingHorizontal: 24,
    paddingVertical: 16,
  },
  tags: {
    display: "flex",
    flexDirection: "row",
    gap: 8,
  },
  join: {
    borderRadius: 6,
    paddingVertical: 6,
    paddingHorizontal: 8,
    borderColor: colors.success,
    borderWidth: 1,
    width: "auto",
  },
  join__text: {
    fontFamily: fonts.Unbounded,
    color: colors.success,
    fontWeight: 600,
    fontSize: 8,
    textTransform: "uppercase",
  },
  public: {
    borderRadius: 6,
    paddingVertical: 6,
    paddingHorizontal: 8,
    borderColor: colors.primary,
    borderWidth: 1,
  },
  public__text: {
    fontFamily: fonts.Unbounded,
    color: colors.primary,
    fontWeight: 600,
    fontSize: 8,
    textTransform: "uppercase",
  },
  creator: {
    borderRadius: 6,
    paddingVertical: 6,
    paddingHorizontal: 8,
    backgroundColor: "#CA3CFF",
    justifyContent: "center",
    alignItems: "center",
    width: "auto",
  },
  creator__text: {
    fontFamily: fonts.Unbounded,
    color: "#ffffff",
    fontWeight: 600,
    fontSize: 8,
    textTransform: "uppercase",
  },
  mainText: {
    height: 350,
  },
  title: {
    marginTop: 8,
    color: colors.black,
    fontFamily: fonts.Unbounded,
    fontWeight: 800,
    fontSize: 20,
  },
  category: {
    color: colors.secondary,
    fontFamily: fonts.Montserrat,
    fontWeight: 600,
    fontSize: 12,
    marginTop: 2,
  },
  description: {
    color: colors.secondary,
    fontFamily: fonts.Montserrat,
    fontWeight: 400,
    fontSize: 12,
    marginTop: 16,
  },
  date: {
    display: "flex",
    flexDirection: "row",
    gap: 8,
  },
  date__item: {
    borderRadius: 6,
    backgroundColor: colors.primary,
    display: "flex",
    justifyContent: "center",
    alignContent: "center",
    paddingHorizontal: 8,
    paddingVertical: 6,
  },
  item__text: {
    fontFamily: fonts.Unbounded,
    fontWeight: 600,
    fontSize: 10,
    color: "#ffffff",
  },
  location: {
    marginTop: 8,
    display: "flex",
    flexDirection: "row",
    gap: 8,
    alignItems: "flex-start",
  },
  location__text: {
    fontFamily: fonts.Montserrat,
    fontWeight: 500,
    fontSize: 12,
    color: colors.secondary,
  },
  button: {
    bottom: 24,
    paddingHorizontal: 24,
  },
  edit_buttons: {
    marginHorizontal: 24,
    display: "flex",
    flexDirection: "row",
    justifyContent: "space-between",
    marginBottom: 24,
    gap: 8,
  },
  completed: {
    width: "100%",
    height: 48,
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    marginBottom: 32,
  },
  completed__text: {
    fontFamily: fonts.Unbounded,
    fontSize: 12,
    fontWeight: 600,
    color: colors.grey_text,
  },
});

export default ModalScreen;
