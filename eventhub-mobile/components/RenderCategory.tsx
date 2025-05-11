import React, { useState } from "react"
import { StyleSheet, Text, TouchableOpacity, View } from "react-native"
import EventCard from "./EventCard"
import { format, parseISO } from "date-fns"
import { ru } from "date-fns/locale"
import { colors } from "@/constants/colors"
import { fonts } from "@/constants/fonts"
import EmptyState from "./EmptyState"

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

type RenderCategoryProps = {
    events: Event[],
    title: string,
    emptyMessage: string,
    emptyInfo: string,
    isCompleted: boolean
}

const formatDate = (date: string) => {
    const parsedDate = parseISO(date)
    const formattedDate = format(parsedDate, 'd MMMM', {
        locale: ru,
    }).toUpperCase()

    return formattedDate
}

const formatTime = (time: string) => {
    if (!time) return '00:00'
    return time.slice(0, 5)
}

const renderEvents = (events: Event[], eventsToShow: number, isCompleted: boolean) => {
    return events.slice(0, eventsToShow).map((event) => {
        return (
            <View style={styles.category__item} key={'event_' + event.id}>
                <EventCard
                    id={event.id}
                    title={event.title}
                    category={event.category}
                    date={formatDate(event.date)}
                    start_time={formatTime(event.start_time)}
                    end_time={formatTime(event.end_time)}
                    location={event.location}
                    isCompleted={isCompleted}
                />
            </View>
        )
    })
}

const RenderCategory = ({ events, title, emptyMessage, emptyInfo, isCompleted }: RenderCategoryProps) => {
    const step = 2
    const [eventsToShow, setEventsToShow] = useState(step)

    return (
        <View style={styles.category}>
            {
                events.length == 0
                    ? emptyMessage == ''
                        ? null
                        : <EmptyState message={emptyMessage} info={emptyInfo} />
                    : <View>
                        <Text style={styles.category__title}>{title}</Text>
                        <View style={styles.category__content}>
                            {renderEvents(events, eventsToShow, isCompleted)}
                        </View>
                        {
                            eventsToShow >= events.length
                                ? null
                                : <TouchableOpacity onPress={() => setEventsToShow(eventsToShow + step)}>
                                    <View style={styles.button}>
                                        <Text style={styles.button__title}>Загрузить еще</Text>
                                    </View>
                                </TouchableOpacity>
                        }
                    </View>
            }
        </View>
    )
}

const styles = StyleSheet.create({
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
    button: {
        borderRadius: 12,
        width: "100%",
        padding: 12,
        display: "flex",
        justifyContent: "center",
        alignItems: "center"
    },
    button__title: {
        fontFamily: fonts.Unbounded,
        fontSize: 10,
        color: colors.primary,
        fontWeight: 400
    }
})

export default RenderCategory
