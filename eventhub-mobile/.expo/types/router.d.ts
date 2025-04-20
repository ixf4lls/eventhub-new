/* eslint-disable */
import * as Router from 'expo-router';

export * from 'expo-router';

declare module 'expo-router' {
  export namespace ExpoRouter {
    export interface __routes<T extends string = string> extends Record<string, unknown> {
      StaticRoutes: `/` | `/(auth)` | `/(auth)/AuthContext` | `/(auth)/login` | `/(auth)/registration` | `/(auth)/welcome` | `/(root)` | `/(root)/(tabs)` | `/(root)/(tabs)/home` | `/(root)/(tabs)/notifications` | `/(root)/(tabs)/organizations` | `/(root)/(tabs)/profile` | `/(root)/home` | `/(root)/notifications` | `/(root)/organizations` | `/(root)/profile` | `/(tabs)` | `/(tabs)/home` | `/(tabs)/notifications` | `/(tabs)/organizations` | `/(tabs)/profile` | `/AuthContext` | `/_sitemap` | `/home` | `/login` | `/modal` | `/notifications` | `/organizations` | `/profile` | `/registration` | `/welcome`;
      DynamicRoutes: never;
      DynamicRouteTemplate: never;
    }
  }
}
