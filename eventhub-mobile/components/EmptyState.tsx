import { colors } from "@/constants/colors"
import { fonts } from "@/constants/fonts"
import { StyleSheet, Text, View } from "react-native"

type EmptySpaceProps = {
    message: string,
    info: string
}

const EmptyState = ({ message, info }: EmptySpaceProps) => {
    return (
        <View style={styles.container}>
            <Text style={styles.icon}>😔</Text>
            <View style={styles.text}>
                <Text style={styles.text__message}>{message}</Text>
                <Text style={styles.text__info}>{info}</Text>
            </View>
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        width: "100%",
        backgroundColor: "#F8F9FE",
        borderRadius: 16,
        padding: 24,
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        flexDirection: "column",
        gap: 8
    },
    icon: {
        fontSize: 32,
    },
    text: {
        display: "flex",
        flexDirection: "column",
        gap: 4,
        alignItems: "center"
    },
    text__message: {
        fontFamily: fonts.Unbounded,
        fontSize: 14,
        fontWeight: 500,
        color: colors.black,
        textAlign: "center"
    },
    text__info: {
        fontFamily: fonts.Unbounded,
        fontSize: 12,
        fontWeight: 400,
        color: colors.grey_text,
        textAlign: "center"
    }
})

export default EmptyState