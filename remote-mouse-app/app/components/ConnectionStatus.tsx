import React, { memo } from 'react';
import { View, Text, TouchableOpacity, StyleSheet } from 'react-native';
import { ConnectionProps } from '../types';

const ConnectionStatus = memo(({ connected, connect, disconnect }: ConnectionProps) => {
    return (
        <View style={styles.statusContainer}>
            <View style={[
                styles.statusIndicator,
                { backgroundColor: connected ? '#4CAF50' : '#F44336' }
            ]} />

            <Text style={styles.statusText}>
                {connected ? 'Connected' : 'Disconnected'}
            </Text>

            <TouchableOpacity
                style={[
                    styles.connectionButton,
                    { backgroundColor: connected ? '#F44336' : '#4CAF50' }
                ]}
                onPress={connected ? disconnect : connect}
            >
                <Text style={styles.buttonText}>
                    {connected ? 'Disconnect' : 'Connect'}
                </Text>
            </TouchableOpacity>
        </View>
    );
});

const styles = StyleSheet.create({
    statusContainer: {
        flexDirection: 'row',
        alignItems: 'center',
        padding: 15,
        backgroundColor: 'white',
        borderBottomWidth: 1,
        borderBottomColor: '#e0e0e0',
    },
    statusIndicator: {
        width: 12,
        height: 12,
        borderRadius: 6,
        marginRight: 10,
    },
    statusText: {
        flex: 1,
        fontSize: 16,
    },
    connectionButton: {
        paddingVertical: 8,
        paddingHorizontal: 15,
        borderRadius: 5,
    },
    buttonText: {
        color: 'white',
        fontWeight: 'bold',
    },
});

export default ConnectionStatus; 