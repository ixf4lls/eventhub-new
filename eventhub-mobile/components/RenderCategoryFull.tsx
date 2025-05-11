import React, { useState } from "react";
import { StyleSheet, Text, TouchableOpacity, View } from "react-native";
import EventCard from "./EventCard";
import { format, parseISO } from "date-fns";
import { ru } from "date-fns/locale";
import { colors } from "@/constants/colors";
import { fonts } from "@/constants/fonts";
import EmptyState from "./EmptyState";

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
};

type RenderCategoryProps = {
  events: Event[];
  emptyMessage: string;
  emptyInfo: string;
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

const renderEvents = (events: Event[], isCompleted: boolean) => {
  return events.map((event) => {
    return (
      <View style={styles.category__item} key={"event_" + event.id}>
        <EventCard
          id={event.id}
          title={event.title}
          category={event.category}
          date={formatDate(event.date)}
          start_time={formatTime(event.start_time)}
          end_time={formatTime(event.end_time)}
          location={event.location}
          isCompleted={isCompleted}
        />
      </View>
    );
  });
};

const RenderCategoryFull = ({
  events,
  emptyMessage,
  emptyInfo,
}: RenderCategoryProps) => {
  const activeEvents = events.filter(
    (event: Event) => event.status == "active",
  );
  const completedEvents = events.filter(
    (event: Event) => event.status == "completed",
  );

  return (
    <>
      <View style={styles.category}>
        <View>
          <View style={styles.category__content}>
            {activeEvents.length == 0 ? (
              <EmptyState message={emptyMessage} info={emptyInfo} />
            ) : (
              <View>
                <Text style={styles.category__title}>Актуальные</Text>
                {renderEvents(activeEvents, false)}
              </View>
            )}
          </View>
        </View>
      </View>

      <View style={styles.category}>
        <View>
          <View style={styles.category__content}>
            {completedEvents.length == 0 ? (
              null
            ) : (
              <View>
                <Text style={styles.category__title}>Завершенные</Text>
                {renderEvents(completedEvents, true)}
              </View>
            )}
          </View>
        </View>
      </View>
    </>
  );
};

const styles = StyleSheet.create({
  category: {
    marginBottom: 32,
  },
  category__title: {
    color: colors.grey_text,
    fontFamily: fonts.Unbounded,
    fontSize: 14,
    fontWeight: 600,
    marginLeft: 8,
    marginBottom: 8,
  },
  category__content: {
    paddingHorizontal: 16,
  },
  category__item: {
    marginBottom: 12,
  },
  button: {
    borderRadius: 12,
    width: "100%",
    padding: 12,
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
  },
  button__title: {
    fontFamily: fonts.Unbounded,
    fontSize: 10,
    color: colors.primary,
    fontWeight: 400,
  },
});

export default RenderCategoryFull;
