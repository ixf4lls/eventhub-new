import { colors } from '@/constants/colors'
import { fonts } from '@/constants/fonts'
import { router } from 'expo-router'
import { Image, StyleSheet, Text, TouchableOpacity, View } from 'react-native'

type EventCardProps = {
  id: number
  title: string
  category: string
  date: string
  start_time: string
  end_time: string
  location: string
}

const EventCard = ({
  id,
  title,
  category,
  date,
  start_time,
  end_time,
  location,
}: EventCardProps) => {
  return (
    <TouchableOpacity
      onPress={() =>
        router.push({
          pathname: '/event_modal',
          params: {
            id: id,
          },
        })
      }
      style={styles.card}
    >

      <View style={styles.image_container}>
        <Image
          source={require('../assets/images/blank_image_card.png')}
          resizeMode="cover"
          style={styles.image}
        />
      </View>
      <View style={styles.card__content}>
        <View style={styles.content__date}>
          <View style={styles.date__day}>
            <Text style={styles.date__text}>{date}</Text>
          </View>
          <View style={styles.date__time}>
            <Text style={styles.date__text}>
              {start_time} - {end_time}
            </Text>
          </View>
        </View>
        <View style={{width: "100%"}}>
          <Text style={styles.content__title}>{title}</Text>
          <Text style={styles.content__category}>{category}</Text>
        </View>
        <View style={styles.content__place}>
          <Image
            source={require('../assets/icons/location.png')}
            style={{ width: 10, height: 13 }}
            resizeMode="contain"
          />
          <Text style={styles.place__text}>{location}</Text>
        </View>
      </View>
    </TouchableOpacity>
  )
}

const styles = StyleSheet.create({
  card: {
    width: "100%",
    borderRadius: 16,
    backgroundColor: '#F8F9FE',
    overflow: 'hidden',
    display: "flex",
    flexDirection: "row",
    maxHeight: 200,
    position: "relative"
    // boxShadow:
    // 'rgba(14, 63, 126, 0.03) 0px 0px 0px 1px, rgba(42, 51, 69, 0.03) 0px 1px 1px -0.5px, rgba(42, 51, 70, 0.03) 0px 3px 3px -1.5px, rgba(42, 51, 70, 0.03) 0px 6px 6px -3px, rgba(14, 63, 126, 0.03) 0px 12px 12px -6px, rgba(14, 63, 126, 0.03) 0px 24px 24px -12px',
  },
  image_container: {
    width: 135,
    position: 'relative', 
  },
  image: {
    position: 'absolute',
    top: 0,
    left: 0,
    right: 0,
    bottom: 0,
    width: '100%',
    height: '100%',
  },
  card__content: {
    padding: 16,
    boxSizing: 'border-box',
    flexGrow: 1,
    flexShrink: 1
  },
  content__date: {
    display: 'flex',
    flexDirection: 'row',
  },
  date__day: {
    paddingHorizontal: 8,
    paddingVertical: 6,
    borderStyle: 'solid',
    borderColor: colors.primary,
    borderWidth: 1,
    borderRadius: 6,
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
  },
  date__time: {
    paddingHorizontal: 8,
    paddingVertical: 6,
    borderStyle: 'solid',
    borderColor: colors.primary,
    borderWidth: 1,
    borderRadius: 6,
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    marginLeft: 8,
  },
  date__text: {
    fontWeight: 600,
    fontFamily: fonts.Unbounded,
    fontSize: 8,
    textTransform: 'uppercase',
    color: colors.primary,
  },
  content__title: {
    fontWeight: 700,
    fontFamily: fonts.Unbounded,
    marginTop: 8,
    color: colors.black,
    fontSize: 14,
    flexShrink: 1,
    flexWrap: "wrap"
  },
  content__category: {
    fontWeight: 600,
    fontFamily: fonts.Montserrat,
    fontSize: 12,
    color: colors.secondary,
    marginTop: 4,
  },
  content__place: {
    marginTop: 16,
    display: 'flex',
    flexDirection: 'row',
    alignItems: 'center',
  },
  place__text: {
    marginLeft: 8,
    fontFamily: fonts.Montserrat,
    fontSize: 12,
    color: colors.secondary,
    fontWeight: 500,
  },
})

export default EventCard
