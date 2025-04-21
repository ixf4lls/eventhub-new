import CustomButton from "@/components/CustomButton";
import InputField from "@/components/InputField";
import { ADDRESS } from "@/constants/address";
import { colors } from "@/constants/colors";
import { fonts } from "@/constants/fonts";
import { fetchWithToken } from "@/utils/tokenInterceptor";
import AsyncStorage from "@react-native-async-storage/async-storage";
import { router } from "expo-router";
import { useEffect, useState } from "react";
import { StyleSheet, Text, View, Dimensions, Modal, Alert, ScrollView, RefreshControl } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import * as Haptics from 'expo-haptics'
import OrganizationCard from "@/components/OrganizationCard";
import CustomNotification from "@/components/Notification";

type Organization = {
  id: number,
  name: string
  founder_id: number,
  invite_code: string,
}

type Notification = {
  type: 'success' | 'error',
  message: string
}

const Organizations = () => {
  const [joined, setJoined] = useState<Organization[]>([])
  const [founded, setFounded] = useState<Organization[]>([])

  const [joinByCodeMenuActive, setJoinByCodeMenuActive] = useState(false)
  const [createOrgMenuActive, setCreateOrgMenuActive] = useState(false)

  const [tempOrgName, setTempOrgName] = useState("")
  const [tempInviteCode, setTempInviteCode] = useState("")

  const [refreshing, setRefreshing] = useState(false)

  const [notifications, setNotifications] = useState<Notification[]>([])

  const LoadOrganizations = async () => {
    setRefreshing(true)
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
    } finally{
      setRefreshing(false)
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
      
      setCreateOrgMenuActive(false)
      LoadOrganizations()
    } catch(error){
      Alert.alert('Ошибка сервера', 'Проверьте подключение.')  
    }
  }

  const joinByCode = async () => {
    if (tempInviteCode == ''){
      Alert.alert('Ошибка', "Укажите код приглашения")
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
        'http://' + ADDRESS + '/api/organizations/join/' + tempInviteCode.toUpperCase(),
        {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' }
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

  const renderCreateOrgMenu = () => {
    return (
      <Modal
        visible={createOrgMenuActive}
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
          <CustomButton onPress={() => setCreateOrgMenuActive(false)} title="Назад" type="action" fill="bordered" />
        </View>
        <View style={styles.overlay} />
      </Modal>
    )
  }

  const renderJoinOrgMenu = () => {
    return (
      <Modal
        visible={joinByCodeMenuActive}
        transparent={true}
        animationType="fade"
        onRequestClose={() => setJoinByCodeMenuActive(false)}
      >
        <View style={styles.menu}>
          <Text style={styles.menu__title}>Присоединиться по коду</Text>
          <InputField value={tempInviteCode} onChangeText={(text) => setTempInviteCode(text.toUpperCase())} placeholder="Введите код приглашения"  secureTextEntry={false}/>
          <View style={{width:"100%", marginBottom:8, marginTop: 16}}>
            <CustomButton onPress={() => joinByCode()} title="Присоединиться" type="action" fill="solid" />
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
        <View style={{marginBottom: 8}}><CustomButton onPress={() => setJoinByCodeMenuActive(true)} title={"Присоединиться по коду"} type='action' fill='solid'/></View>
        <CustomButton onPress={() => setCreateOrgMenuActive(true)} title={"Создать организацию"} type='action' fill='bordered'/>
      </View>
    )
  }

  const renderOrgs = () => {
    return (
      <View>
        <View>
          {founded.length > 0 && (
            <View style={{marginBottom: 16}}>
              <Text style={styles.sectionTitle}>Вы управляете</Text>
              {founded.map(org => (
                <OrganizationCard 
                  key={org.id}
                  id={org.id} 
                  name={org.name} 
                  founder_id={org.founder_id} 
                  invite_code={org.invite_code}
                />
              ))}
            </View>
          )}
          {joined.length > 0 && (
            <View>
              <Text style={styles.sectionTitle}>Вы участник</Text>
              {joined.map(org => (
                <OrganizationCard 
                  key={org.id}
                  id={org.id} 
                  name={org.name} 
                  founder_id={org.founder_id} 
                  invite_code={org.invite_code}
                />
              ))}
            </View>
          )}
        </View>
        <View style={styles.orgsActions}>
          <View style={{marginBottom: 8}}><CustomButton onPress={() => setJoinByCodeMenuActive(true)} title={"Присоединиться по коду"} type='action' fill='solid'/></View>
          <CustomButton onPress={() => setCreateOrgMenuActive(true)} title={"Создать организацию"} type='action' fill='bordered'/>
        </View>
      </View>
    )
  }

  useEffect(() => {
    LoadOrganizations()
    const interval = setInterval(() => {
      LoadOrganizations()
    }, 30000)

    return () => clearInterval(interval)
  }, [])

  return (
    <SafeAreaView style={{backgroundColor: "#ffffff", flex:1}}>
      {createOrgMenuActive ? renderCreateOrgMenu() : null}
      {joinByCodeMenuActive ? renderJoinOrgMenu() : null}
      <View style={styles.header}><Text style={styles.header__text}>Организации</Text></View>
      <ScrollView
        refreshControl={
          <RefreshControl refreshing={refreshing} onRefresh={LoadOrganizations} />
      }
      >
        {(joined.length === 0 && founded.length === 0) ? renderNoOrgs() : renderOrgs()}
        {/* <CustomNotification type='success' message="Вы присоединились к организации" /> */}
      </ScrollView>
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
    marginTop: 8,
    marginBottom: 16,
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
    borderRadius: 16,
    marginBottom: 32
  },
  welcome__text:{
    fontWeight: 400,
    fontFamily: fonts.Unbounded,
    fontSize:16,
    color: colors.black,
    marginBottom: 16
  },
  sectionTitle: {
    fontFamily: fonts.Unbounded,
    color: colors.grey_text,
    fontWeight: 600,
    fontSize: 14,
    marginLeft: 16,
    marginBottom: 8
  },
  orgsActions: {
    marginHorizontal: 16,
    paddingVertical: 16,
    marginBottom: 32
  }
})


export default Organizations;
