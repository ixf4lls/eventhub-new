import { colors } from "@/constants/colors"
import { fonts } from "@/constants/fonts"
import { Image, StyleSheet, Text, View } from "react-native"

type OrgCardProps = {
    id: number,
    name: string
    founder_id: number,
    invite_code: string,
}

const OrganizationCard = ({id, name, founder_id, invite_code}: OrgCardProps) => {
    return( 
        <View style={styles.card}>
            <Text style={styles.card__title}>{name}</Text>
            <Image source={require('../assets/icons/arrow_right.png')} style={{width: 12, height: 12}}></Image>
        </View>
    )
}

const styles = StyleSheet.create({
    card: {
        backgroundColor: "#EAF2FF",
        borderRadius: 16,
        paddingVertical: 16, 
        paddingHorizontal: 16,
        display: "flex",
        flexDirection: "row",
        justifyContent: "space-between",
        alignItems: "center",
        marginHorizontal: 16,
        marginBottom: 8
    },
    card__title:{
        fontFamily: fonts.Unbounded,
        fontWeight: 500,
        fontSize:16,
        color: colors.black
    }
})

export default OrganizationCard