import { AuthContext } from '@/src/context/AuthContext'
import { ADDRESS } from '@/constants/address'
import AsyncStorage from '@react-native-async-storage/async-storage'
import { router } from 'expo-router'
import { useCallback, useContext, useEffect, useState } from 'react'
import { Button, Text } from 'react-native'
import { SafeAreaView } from 'react-native-safe-area-context'
import { fetchWithToken } from '@/utils/tokenInterceptor'

const Profile = () => {
  // const LoadUserData = async () => {
  //   try {
  //     const response = await fetchWithToken(
  //       `http://${ADDRESS}/`
  //     )
  //   }
  // }

  return (
    <SafeAreaView>
      {/* <Text>first name: {firstName}</Text>
      <Text>last name: {lastName}</Text>
      <Text>username: {username}</Text> */}

      <Button
        title="Logout"
        onPress={() => {
          AsyncStorage.removeItem('token')
          router.replace('/(auth)/login')
        }}
      />
    </SafeAreaView>
  )
}

export default Profile
