import { Ionicons } from '@expo/vector-icons';
import { ComponentProps } from 'react';

export function TabBarIcon({ style, ...rest }: ComponentProps<typeof Ionicons>) {
    return <Ionicons size={24} style={[{ marginBottom: -3 }, style]} {...rest} />;
}
