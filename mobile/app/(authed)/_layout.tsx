import { Redirect, Stack } from "expo-router";

export default function AppLayout() {
    // check from context if use if logged in

    const isLoggedIn = false;

    if (!isLoggedIn) {
        return <Redirect href="/login" />;
    }

    return <Stack screenOptions={{ headerShown: false }} />;
}