import axios, { AxiosError, AxiosInstance, AxiosResponse } from "axios";
import AsyncStorage from "@react-native-async-storage/async-storage";
import Constants from "expo-constants";

// Get the dev server host IP so physical devices can reach the backend.
// In dev, Expo exposes the host machine's IP via debuggerHost (e.g. "192.168.100.226:8081").
// This works for both emulators and physical devices on the same Wi-Fi network.
const devServerHost = Constants.expoConfig?.hostUri?.split(":")[0] 
  ?? Constants.manifest2?.extra?.expoGo?.debuggerHost?.split(":")[0]
  ?? "127.0.0.1";

const url = `http://${devServerHost}:3000`

const Api: AxiosInstance = axios.create({baseURL: url + "/api"})

Api.interceptors.request.use(async config => {
    const token = await AsyncStorage.getItem("token")

    if (token) config.headers.set("Authorization", `Bearer ${token}`)

    return config
})

Api.interceptors.response.use(
    async (res: AxiosResponse) => res.data,
    async (err: AxiosError) => Promise.reject(err)
)

export { Api }