import { User } from "@/types/user";
import AsyncStorage from "@react-native-async-storage/async-storage";
import { router } from "expo-router";
import { createContext, PropsWithChildren, useContext, useEffect, useState } from "react";


interface AuthContextProps {
    isLoggedIn: boolean;
    isLoadingAuth: boolean;
    authentication: (authMode: "login" | "register", email: string, password: string) => Promise<void>;
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

    const authentication = async (authMode: "login" | "register", email: string, password: string) => {
        // TODO: implement
    };

    const logout = () => {
        // TODO: implement
    };

    return (
        <AuthContext.Provider
            value={{ isLoggedIn, isLoadingAuth, user, authentication, logout }}
        >
            {children}
        </AuthContext.Provider>
    )
}