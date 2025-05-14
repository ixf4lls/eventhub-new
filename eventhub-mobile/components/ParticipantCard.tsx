import { colors } from "@/constants/colors";
import { fonts } from "@/constants/fonts";
import { Image, StyleSheet, Text, View } from "react-native";

type ParticipantCardProps = {
  username: string;
  first_name: string;
  last_name: string;
};

const ParticipantCard = ({
  username,
  first_name,
  last_name,
}: ParticipantCardProps) => {
  return (
    <View style={styles.card}>
      <Image
        source={require("../assets/icons/avatar.png")}
        style={{ width: 35, height: 35 }}
      />
      <View style={styles.text}>
        <Text style={styles.name}>
          {first_name} {last_name}
        </Text>
        <Text style={styles.username}>@{username}</Text>
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  card: {
    padding: 16,
    borderRadius: 16,
    backgroundColor: "#F8F9FE",
    display: "flex",
    flexDirection: "row",
  },
  text: {
    marginLeft: 16,
  },
  name: {
    fontFamily: fonts.Montserrat,
    fontSize: 14,
    fontWeight: 700,
    color: colors.black,
  },
  username: {
    fontFamily: fonts.Montserrat,
    fontSize: 12,
    fontWeight: 600,
    color: "#71727A",
    marginTop: 4,
  },
});

export default ParticipantCard;
