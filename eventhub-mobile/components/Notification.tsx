import { colors } from "@/constants/colors";
import { fonts } from "@/constants/fonts";
import { format, parseISO } from "date-fns";
import { ru } from "date-fns/locale";
import { StyleSheet, Text, View } from "react-native";

type NotificationProps = {
  message: string;
  info: string;
  createdAt: string;
};

const formatDate = (dateString: string) => {
  const cleaned = dateString.split(" ")[0] + "T" + dateString.split(" ")[1];
  const date = new Date(cleaned);

  const now = new Date();
  const isToday =
    date.getDate() === now.getDate() &&
    date.getMonth() === now.getMonth() &&
    date.getFullYear() === now.getFullYear();

  const yesterday = new Date(now);
  yesterday.setDate(yesterday.getDate() - 1);
  const isYesterday =
    date.getDate() === yesterday.getDate() &&
    date.getMonth() === yesterday.getMonth() &&
    date.getFullYear() === yesterday.getFullYear();

  if (isToday) return `Сегодня, ${format(date, "HH:mm")}`;
  if (isYesterday) return `Вчера, ${format(date, "HH:mm")}`;

  return format(date, "d MMMM, HH:mm", { locale: ru });
};

const NotificationCard = ({ message, info, createdAt }: NotificationProps) => {
  return (
    <View style={styles.notification}>
      <Text style={styles.time}>{formatDate(createdAt)}</Text>
      <Text style={styles.message}>{message}</Text>
      <Text style={styles.info}>{info}</Text>
    </View>
  );
};

const styles = StyleSheet.create({
  notification: {
    backgroundColor: "#F8F9FE",
    padding: 16,
    display: "flex",
    flexDirection: "column",
    gap: 4,
    borderRadius: 16,
  },
  time: {
    fontFamily: fonts.Montserrat,
    fontWeight: 600,
    fontSize: 10,
    color: "#71727A",
  },
  message: {
    fontFamily: fonts.Montserrat,
    fontWeight: 700,
    fontSize: 14,
    color: colors.black,
  },
  info: {
    fontFamily: fonts.Montserrat,
    fontWeight: 400,
    fontSize: 12,
    color: "#71727A",
  },
});

export default NotificationCard;
