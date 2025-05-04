import { Redirect, SplashScreen } from "expo-router";
import { useAuth } from "../src/context/AuthContext";
import useAppFonts from "@/utils/useAppFonts";
import { ActionSheetProvider } from '@expo/react-native-action-sheet';
import { hideAsync } from "expo-splash-screen";

const Home = () => {
  const { user } = useAuth();
  useAppFonts();
  SplashScreen.preventAutoHideAsync()
  hideAsync()
  // prod
  return (
    <ActionSheetProvider>
      <Redirect href={user ? "/(root)/(tabs)/home" : "/(auth)/welcome"} />
    </ActionSheetProvider>
  )

  // dev
  // return <Redirect href={"/(root)/(tabs)/home"} />;
};

export default Home;
