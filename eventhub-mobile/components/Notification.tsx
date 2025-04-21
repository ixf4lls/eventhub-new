import { colors } from "@/constants/colors"
import { fonts } from "@/constants/fonts"
import { Text } from "react-native"
import { Image, StyleSheet, View } from "react-native"

type CustomNotificationProps = {
    type: 'success' | 'error',
    message: string
}

const CustomNotification = ({type, message}: CustomNotificationProps) => {
    return (
        <View style={styles.notification}>
            <Image source={require('../assets/icons/success.png')} style={{width: 24, height: 24}}></Image>
            <Text style={styles.message}>{message}</Text>
        </View>
    )
}

const styles = StyleSheet.create({
    notification:{
        backgroundColor: colors.success,
        paddingVertical: 12,
        paddingHorizontal: 16,
        display: "flex",
        flexDirection: "row",
        alignItems:"center",
        gap: 10,
        marginHorizontal: 16,
        borderRadius: 12
    },
    message:{
        color: "#ffffff",
        fontFamily: fonts.Unbounded,
        fontWeight: 500,
        fontSize: 12
    }
})

export default CustomNotification