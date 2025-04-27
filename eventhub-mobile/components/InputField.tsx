import { colors } from "@/constants/colors";
import { memo, useRef, useState } from "react";
import {
  StyleSheet,
  TextInput,
  View,
  Image,
  TouchableOpacity,
} from "react-native";
import * as Haptics from "expo-haptics";
import { fonts } from "@/constants/fonts";

type InputFieldProps = {
  value: string;
  onChangeText: (text: string) => void;
  placeholder: string;
  isIncorrect?: boolean;
  secureTextEntry?: boolean;
  props?: any;
};

const InputField = memo(({
  value,
  onChangeText,
  placeholder,
  isIncorrect,
  secureTextEntry,
  ...props
}: InputFieldProps) => {
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
        secureTextEntry={secureTextEntry}
        autoCapitalize="none"
        autoCorrect={false}
        {...props}
      />
      {secureTextEntry && (
        <TouchableOpacity
          onPress={() => {
            Haptics.impactAsync(Haptics.ImpactFeedbackStyle.Soft);
          }}
          style={styles.passwordVisibility}
        >
          <Image
            source={
              secureTextEntry
                ? require("../assets/icons/password_hide.png")
                : require("../assets/icons/password_show.png")
            }
            style={{ width: 16, height: 16 }}
          />
        </TouchableOpacity>
      )}
    </View>
  );
});

const styles = StyleSheet.create({
  container: {
    width: "100%",
  },
  input: {
    width: "100%",
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
  passwordVisibility: {
    position: "absolute",
    marginTop: 16,
    right: 16,
  },
});

export default InputField;
