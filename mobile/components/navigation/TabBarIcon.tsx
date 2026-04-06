import { Ionicons } from '@expo/vector-icons';
import { ComponentProps } from 'react';

export function TabBarIcon(props: {
    name: ComponentProps<typeof Ionicons>["name"];
    color: string;
}) {
    return <Ionicons size={24} style={{ marginBottom: -3 }} {...props} />;
}
