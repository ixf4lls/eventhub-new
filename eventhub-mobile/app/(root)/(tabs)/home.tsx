import { AuthContext } from '@/src/context/AuthContext'
import EventCard from '@/components/EventCard'
import { colors } from '@/constants/colors'
import { fonts } from '@/constants/fonts'
import { router } from 'expo-router'
import { useContext, useEffect, useRef, useState } from 'react'
import {
  Alert,
  Button,
  Image,
  RefreshControl,
  ScrollView,
  StyleSheet,
  Text,
  View,
} from 'react-native'
import { SafeAreaView } from 'react-native-safe-area-context'
import { format, parseISO } from 'date-fns'
import { ru } from 'date-fns/locale'
import { ADDRESS } from '@/constants/address'
import { fetchWithToken } from '@/utils/tokenInterceptor'
import AsyncStorage from '@react-native-async-storage/async-storage'
import EmptySpace from '@/components/EmptyState'
import RenderCategory from '@/components/RenderCategory'
import Swiper from 'react-native-swiper'
import RenderCategoryFull from '@/components/RenderCategoryFull'

type Event = {
  id: number
  title: string
  description: string
  category: string
  is_public: boolean
  status: string
  date: string
  start_time: string
  end_time: string
  location: string
  creator_id: number
}

export default function Home() {
  const [joinedEvents, setJoinedEvents] = useState<Event[]>([])
  const [openEvents, setOpenEvents] = useState<Event[]>([])

  const [loading, setLoading] = useState(false)

  const [refreshing, setRefreshing] = useState(false)

  const [activeIndex, setActiveIndex] = useState(0)

  const swiperRef = useRef<Swiper>(null)

  const LoadEvents = async () => {
    setRefreshing(true)
    const timeout = setTimeout(() => setLoading(true), 300)
    try {
      const token = await AsyncStorage.getItem('accessToken')
      if (!token) {
        router.replace('/(auth)/login')
        return
      }
      const response = await fetchWithToken(
        'http://' + ADDRESS + '/api/events',
        {
          method: 'GET',
        }
      ) as Response

      if (!response.ok) {
        if (response.status === 401) {
          router.replace('/(auth)/login')
          return
        }
        if (response.status === 404) {
          setJoinedEvents([])
          setOpenEvents([])
          return
        }
        try {
          const errorData = await response.json()
          throw new Error(`Ошибка запроса: ${errorData.message || response.status}`)
        } catch (e) {
          throw new Error(`Ошибка запроса: ${response.status}`)
        }
      }

      const data = await response.json()

      const joined = Array.isArray(data.joined_events) ? data.joined_events : []
      const open = Array.isArray(data.open_events) ? data.open_events : []

      setJoinedEvents(joined)

      const joinedEventIds = joined.map((event: Event) => event.id)
      setOpenEvents(open.filter((event: Event) => !joinedEventIds.includes(event.id)))
    } catch (error) {
      console.error('Ошибка запроса:', error)
    } finally {
      clearTimeout(timeout)
      setLoading(false)
      setRefreshing(false)
    }
  }

  useEffect(() => {
    LoadEvents()
    const interval = setInterval(() => {
      LoadEvents()
    }, 30000)

    return () => clearInterval(interval)
  }, [])

  return (
    <SafeAreaView style={{ backgroundColor: '#ffffff' }}>
      <View style={styles.header}>
        <View style={styles.header__content}>
          <Image
            source={require('../../../assets/icons/logo.png')}
            style={{ width: 20, height: 20 }}
            resizeMode="contain"
          />
          <Text style={styles.header__title}>EVENTHUB</Text>
        </View>
      </View>
      <View style={styles.swiper}>
        <View style={styles.swiper__header}>
          <Text style={[styles.swiper__text, activeIndex == 0 ? { color: colors.primary, fontWeight: 600 } : { color: colors.grey_text, fontWeight: 400 }]}>Актуальные</Text>
          <Text style={[styles.swiper__text, activeIndex == 1 ? { color: colors.primary, fontWeight: 600 } : { color: colors.grey_text, fontWeight: 400 }]}>Вы участвуете</Text>
        </View>
        <Swiper
          ref={swiperRef}
          loop={false}
          scrollEnabled={true}
          onIndexChanged={(index) => setActiveIndex(index)}
        >
          <ScrollView
            style={styles.categories}
            refreshControl={
              <RefreshControl refreshing={refreshing} onRefresh={LoadEvents} />
            }
          >
            {
              RenderCategoryFull({ events: openEvents, emptyMessage: 'Актуальных мероприятий пока нет', emptyInfo: "Следите за обновлениями" })
            }
          </ScrollView>

          <ScrollView
            style={styles.categories}
            refreshControl={
              <RefreshControl refreshing={refreshing} onRefresh={LoadEvents} />
            }
          >
            {
              RenderCategoryFull({ events: joinedEvents, emptyMessage: 'Вы пока не участвуете ни в одном мероприятии', emptyInfo: '' })
            }
          </ScrollView>
        </Swiper>
      </View>
    </SafeAreaView>
  )
}

const styles = StyleSheet.create({
  header: {
    position: 'fixed',
    display: 'flex',
    alignItems: 'center',
    marginTop: 8,
    marginBottom: 16,
    boxSizing: 'border-box',
  },
  header__content: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'center',
    alignContent: 'center',
  },
  header__title: {
    fontFamily: fonts.Unbounded,
    fontWeight: 800,
    color: colors.primary,
    marginLeft: 4,
    fontSize: 16,
  },
  categories: {
    width: '100%',
    height: '100%',
  },
  category: {
    marginBottom: 32,
  },
  category__title: {
    color: colors.grey_text,
    fontFamily: fonts.Unbounded,
    fontSize: 14,
    fontWeight: 600,
    marginLeft: 16,
    marginBottom: 8,
  },
  category__content: {
    paddingHorizontal: 16,
  },
  category__item: {
    marginBottom: 12,
  },
  swiper: {
    height: "100%"
  },
  swiper__header: {
    width: "100%",
    display: "flex",
    flexDirection: "row",
    justifyContent: "center",
    alignItems: "center",
    gap: 24,
    paddingTop: 4,
    paddingBottom: 16
  },
  swiper__text: {
    fontFamily: fonts.Unbounded,
    fontSize: 12,
    fontWeight: 500
  }
})
