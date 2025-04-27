declare module 'react-native-calendar' {
    import { ViewStyle, TextStyle } from 'react-native';

    interface CalendarTheme {
        backgroundColor?: string;
        calendarBackground?: string;
        textSectionTitleColor?: string;
        selectedDayBackgroundColor?: string;
        selectedDayTextColor?: string;
        todayTextColor?: string;
        dayTextColor?: string;
        textDisabledColor?: string;
        dotColor?: string;
        selectedDotColor?: string;
        arrowColor?: string;
        monthTextColor?: string;
        indicatorColor?: string;
    }

    interface CalendarProps {
        current?: Date;
        onDateSelect?: (date: Date) => void;
        mode?: 'date' | 'time';
        theme?: CalendarTheme;
        style?: ViewStyle;
    }

    const Calendar: React.FC<CalendarProps>;
    export default Calendar;
} 