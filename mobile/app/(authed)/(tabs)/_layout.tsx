import { Tabs } from "expo-router";
import { ComponentProps } from "react";
import { Text } from "react-native";
import { Ionicons } from "@expo/vector-icons";
import { TabBarIcon } from "../../../components/navigation/TabBarIcon";

export default function TabsLayout() {
    const tabs = [
        {
            showFor: [],
            name: "(events)",
            displayedName: "Events",
            icon: "calendar",
            options: {
                headerShown: false,
            }
        },
        {
            showFor: [],
            name: "(tickets)",
            displayedName: "My Tickets",
            icon: "ticket",
            options: {
                headerShown: false,
            }
        },
        {
            showFor: [],
            name: "scan-ticket",
            displayedName: "Scan Ticket",
            icon: "scan",
            options: {
                headerShown: true,
            }
        },
        {
            showFor: [],
            name: "settings",
            displayedName: "Settings",
            icon: "cog",
            options: {
                headerShown: true,
            }
        },
    ];

    return (
        <Tabs>
            {tabs.map((tab => (
                <Tabs.Screen
                    key={tab.name}
                    name={tab.name}
                    options={{
                        ...tab.options,
                        headerTitle: tab.displayedName,
                        href: tab.name as any,
                        tabBarLabel: ({ focused }) => (
                            <Text style={{ color: focused ? "black" : "gray", fontSize: 12 }}>
                                {tab.displayedName}
                            </Text>
                        ),
                        tabBarIcon: ({ focused }) => (
                            <TabBarIcon
                                name={tab.icon as ComponentProps<typeof Ionicons>["name"]}
                                color={focused ? "black" : "gray"}
                            />
                        )
                    }}
                />
            )))}
        </Tabs>
    )
}