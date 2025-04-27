import { colors } from "@/constants/colors";
import { memo, useState } from "react";
import {
  StyleSheet,
  TextInput,
  View,
} from "react-native";
import { fonts } from "@/constants/fonts";

type TextAreaProps = {
  value: string;
  onChangeText: (text: string) => void;
  placeholder: string;
  isIncorrect?: boolean;
  props?: any;
};

const TextArea = memo(({
  value,
  onChangeText,
  placeholder,
  isIncorrect,
  ...props
}: TextAreaProps) => {
  const [localValue, setLocalValue] = useState(value);

  const handleChange = (text: string) => {
    setLocalValue(text);
    onChangeText(text);
  };

  return (
    <View style={styles.container}>
      <TextInput
        style={[styles.input, isIncorrect && { borderColor: colors.error }]}
        value={localValue}
        onChangeText={handleChange}
        placeholder={placeholder}
        placeholderTextColor={colors.form_text}
        multiline={true}
        numberOfLines={6}
        textAlignVertical="top"
        {...props}
      />
    </View>
  );
});

const styles = StyleSheet.create({
  container: {
    width: "100%",
  },
  input: {
    width: "100%",
    height: 250,
    color: colors.form_text,
    borderColor: colors.form_border,
    borderWidth: 1,
    borderRadius: 12,
    paddingHorizontal: 16,
    paddingVertical: 14,
    boxSizing: "border-box",
    fontFamily: fonts.Montserrat,
    fontSize: 14,
  },
});

export default TextArea; 