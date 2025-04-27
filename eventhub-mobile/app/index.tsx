import { Redirect } from "expo-router";
import { useAuth } from "./(auth)/AuthContext";
import useAppFonts from "@/utils/useAppFonts";
import { ActionSheetProvider } from '@expo/react-native-action-sheet';

const Home = () => {
  const { user } = useAuth();
  useAppFonts();

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
