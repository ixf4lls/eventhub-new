import { Stack, useLocalSearchParams } from "expo-router";
import * as SplashScreen from "expo-splash-screen";
import { AuthProvider } from "../src/context/AuthContext";
import { enableScreens } from "react-native-screens";

// Prevent the splash screen from auto-hiding before asset loading is complete.
SplashScreen.preventAutoHideAsync();

export default function RootLayout() {
  enableScreens();

  return (
    <AuthProvider>
      <Stack>
        <Stack.Screen name="index" options={{ headerShown: false }} />
        <Stack.Screen name="(auth)" options={{ headerShown: false }} />
        <Stack.Screen name="(root)" options={{ headerShown: false }} />
        <Stack.Screen
          name="event_modal"
          options={{ presentation: "modal", headerShown: false }}
        />
        <Stack.Screen
          name="organization_modal" 
          options={{
            headerShown: false,
          }}
        />
        <Stack.Screen
          name="event_creation_modal" 
          options={{
            headerShown: false,
          }}
        />
      </Stack>
    </AuthProvider>
  );
}
