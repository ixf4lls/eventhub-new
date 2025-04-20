import CustomButton from "@/components/CustomButton";
import InputField from "@/components/InputField";
import { ADDRESS } from "@/constants/address";
import { colors } from "@/constants/colors";
import { fonts } from "@/constants/fonts";
import { fetchWithToken } from "@/utils/tokenInterceptor";
import AsyncStorage from "@react-native-async-storage/async-storage";
import { router } from "expo-router";
import { useState } from "react";
import { StyleSheet, Text, View, Dimensions, Modal, Alert } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import * as Haptics from 'expo-haptics'

type Organization = {
  id: number,
  name: string
  founder_id: number,
}

const Organizations = () => {
  const [joined, setJoined] = useState<Organization[]>([])
  const [founded, setFounded] = useState<Organization[]>([])

  const [joinByCodeMenuActive, setJoinByCodeMenuActive] = useState(false)
  const [tempOrgName, setTempOrgName] = useState("")

  const LoadOrganizations = async () =>{
    try {
      const token = await AsyncStorage.getItem('accessToken')
      if (!token) {
        router.replace('/(auth)/login')
        return
      }
      const response = await fetchWithToken(
        'http://' + ADDRESS + '/api/organizations',
        {
          method: 'GET',
        }
      ) as Response

      if (!response.ok){
        if (response.status === 401) {
          router.replace('/(auth)/login')
        }
        throw new Error(`Ошибка запроса: ${response.status}`)
      }

      const data = await response.json()

      const joined = Array.isArray(data.joined) ? data.joined : []
      const founded = Array.isArray(data.founded) ? data.founded : []

      setJoined(joined)
      setFounded(founded)
    } catch(error){
      console.error('Ошибка запроса:', error)      
    }
  }

  const createOrg = async () => {
    if (tempOrgName == ''){
      Alert.alert('Ошибка', "Укажите название организации")
      Haptics.notificationAsync(Haptics.NotificationFeedbackType.Error)
      return
    }
    try {
      const token = await AsyncStorage.getItem('accessToken')
      if (!token) {
        router.replace('/(auth)/login')
        return
      }
      const response = await fetchWithToken(
        'http://' + ADDRESS + '/api/organizations/create',
        {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ "organization_name": tempOrgName }),
        }
      ) as Response

      if (!response.ok){
        const errorData = await response.json()
        if (response.status === 401) {
          router.replace('/(auth)/login')
          Haptics.notificationAsync(Haptics.NotificationFeedbackType.Error)
        } else{
          Alert.alert('Ошибка', errorData.message)
          Haptics.notificationAsync(Haptics.NotificationFeedbackType.Error)
        }
        return
      }
      
      setJoinByCodeMenuActive(false)
      LoadOrganizations()
    } catch(error){
      Alert.alert('Ошибка сервера', 'Проверьте подключение.')  
    }
  }

  const renderCreateOrg = () => {
    return (
      <Modal
        visible={joinByCodeMenuActive}
        transparent={true}
        animationType="fade"
        onRequestClose={() => setJoinByCodeMenuActive(false)}
      >
        <View style={styles.menu}>
          <Text style={styles.menu__title}>Как будет называться ваша организация?</Text>
          <InputField value={tempOrgName} onChangeText={(text) => setTempOrgName(text)} placeholder="Введите название организации"  secureTextEntry={false}/>
          <View style={{width:"100%", marginBottom:8, marginTop: 16}}>
            <CustomButton onPress={() => createOrg()} title="Создать организацию" type="action" fill="solid" />
          </View>
          <CustomButton onPress={() => setJoinByCodeMenuActive(false)} title="Назад" type="action" fill="bordered" />
        </View>
        <View style={styles.overlay} />
      </Modal>
    )
  }

  const renderNoOrgs = () =>{
    return (
      <View style={styles.welcome}>
        <Text style={styles.welcome__text}>👋 Привет! Ты пока не в организации, но это легко исправить!</Text>
        <View style={{marginBottom: 8}}><CustomButton onPress={undefined} title={"Присоединиться по коду"} type='action' fill='solid'/></View>
        <CustomButton onPress={() => setJoinByCodeMenuActive(true)} title={"Создать организацию"} type='action' fill='bordered'/>
      </View>
    )
  }

  const renderOrgs = () => {
    return <Text>asd</Text>
  }

  return (
    <SafeAreaView style={{backgroundColor: "#ffffff", flex:1}}>
      {joinByCodeMenuActive ? renderCreateOrg() : null}
      <View style={styles.header}><Text style={styles.header__text}>Организации</Text></View>
      {(joined.length == 0 && founded.length)== 0 ? renderNoOrgs() : renderOrgs()}
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  overlay:{
    position: "absolute",
    top: 0,
    left:0,
    width: "100%",
    height: "100%",
    backgroundColor:"#1F2024",
    opacity: 0.85,
    zIndex: 5
  },
  menu:{
    marginHorizontal:24,
    backgroundColor: "#ffffff",
    display:"flex",
    justifyContent: "center",
    alignItems:"center",
    marginTop:"auto",
    marginBottom: "auto",
    borderRadius: 16,
    paddingVertical: 16,
    paddingHorizontal:20,
    zIndex: 6
  },
  menu__title:{
    fontFamily: fonts.Unbounded,
    color: colors.black,
    fontWeight: 700,
    fontSize: 18,
    textAlign: "center",
    marginBottom: 16,
    paddingVertical: 8,
    paddingHorizontal: 8
  },
  menu__input:{},
  header:{
    display:"flex",
    position:"fixed",
    width: "100%",
    marginTop: 16,
    marginBottom:32,
    justifyContent: "center",
    alignItems:"center"
  },
  header__text:{
    fontFamily: fonts.Unbounded,
    color: colors.black,
    fontWeight: 600,
    fontSize:16
  },
  welcome:{
    backgroundColor: "#EAF2FF",
    marginHorizontal: 16,
    paddingHorizontal: 16,
    paddingVertical: 16,
    borderRadius: 16
  },
  welcome__text:{
    fontWeight: 400,
    fontFamily: fonts.Unbounded,
    fontSize:16,
    color: colors.black,
    marginBottom: 24
  }
})


export default Organizations;
