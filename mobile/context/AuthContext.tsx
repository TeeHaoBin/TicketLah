import { userService } from "@/services/user";
import { User } from "@/types/user";
import AsyncStorage from "@react-native-async-storage/async-storage";
import { router } from "expo-router";
import { createContext, PropsWithChildren, useContext, useEffect, useState } from "react";


interface AuthContextProps {
    isLoggedIn: boolean;
    isLoadingAuth: boolean;
    authenticate: (authMode: "login" | "register", email: string, password: string) => Promise<void>;
    logout: VoidFunction; // VoidFunction same as () => void
    user: User | null;
}

const AuthContext = createContext({} as AuthContextProps);

// Create custom useAuth hook
export function useAuth() {
    return useContext(AuthContext);
}

export function AuthenticationProvider({ children }: PropsWithChildren) {
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [isLoadingAuth, setIsLoadingAuth] = useState(false);
    const [user, setUser] = useState<User | null>(null);

    useEffect(() => {
        async function checkIfLoggedIn() {
            const token = await AsyncStorage.getItem("token")
            const user = await AsyncStorage.getItem("user")

            if (token && user) {
                setIsLoggedIn(true);
                setUser(JSON.parse(user));
                router.replace("/(authed)/(tabs)/(events)");
            } else {
                setIsLoggedIn(false)
            }
        }

        checkIfLoggedIn();
    }, [])

    const authenticate = async (authMode: "login" | "register", email: string, password: string) => {
        try {
            setIsLoadingAuth(true);
            const res = authMode === "login" 
                ? await userService.login({ email, password })
                : await userService.register({ email, password });
            
            if (res.data?.token && res.data?.user) {
                await AsyncStorage.setItem("token", res.data.token);
                await AsyncStorage.setItem("user", JSON.stringify(res.data.user));
                setIsLoggedIn(true);
                setUser(res.data.user);
                router.replace("/(authed)/(tabs)/(events)");
            }
        } catch (error) {
            console.error("Auth error:", error);
        } finally {
            setIsLoadingAuth(false);
        }
    };

    const logout = async () => {
        await AsyncStorage.removeItem("token");
        await AsyncStorage.removeItem("user");
        setIsLoggedIn(false);
        setUser(null);
        router.replace("/login");
    };

    return (
        <AuthContext.Provider
            value={{ isLoggedIn, isLoadingAuth, user, authenticate, logout }}
        >
            {children}
        </AuthContext.Provider>
    )
}