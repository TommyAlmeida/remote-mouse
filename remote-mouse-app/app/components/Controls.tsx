import React, { memo, useCallback } from 'react';
import { View, Text, TouchableOpacity, StyleSheet } from 'react-native';
import { Ionicons } from '@expo/vector-icons';
import { ControlsProps } from '../types';

const Controls = memo(({ connected, sendMessage, visible }: ControlsProps) => {
    const handleManualMove = useCallback((dx: number, dy: number) => {
        if (connected) {
            const message = `${dx},${dy}`;
            sendMessage(message);
            console.log(`[WebSocket] Manual movement: ${message}`);
        }
    }, [connected, sendMessage]);

    const handleClick = useCallback((type: string) => {
        const message = `click:${type}`;
        if (sendMessage(message)) {
            console.log(`[WebSocket] Sent click: ${message}`);
        } else {
            console.log('[WebSocket] Cannot send click, not connected');
        }
    }, [sendMessage]);

    if (!visible) {
        return (
            <View style={styles.mouseButtonsContainer}>
                <TouchableOpacity
                    style={[styles.mouseButton, !connected && styles.disabledButton]}
                    onPress={() => handleClick('left')}
                    disabled={!connected}
                >
                    <Text style={styles.buttonText}>Left Click</Text>
                </TouchableOpacity>

                <TouchableOpacity
                    style={[styles.mouseButton, !connected && styles.disabledButton]}
                    onPress={() => handleClick('right')}
                    disabled={!connected}
                >
                    <Text style={styles.buttonText}>Right Click</Text>
                </TouchableOpacity>

                <TouchableOpacity
                    style={[styles.mouseButton, !connected && styles.disabledButton]}
                    onPress={() => handleClick('double')}
                    disabled={!connected}
                >
                    <Text style={styles.buttonText}>Double Click</Text>
                </TouchableOpacity>
            </View>
        );
    }

    return (
        <>
            <View style={styles.manualControlsContainer}>
                <TouchableOpacity
                    style={[styles.arrowButton, !connected && styles.disabledButton]}
                    onPress={() => handleManualMove(0, -10)}
                    disabled={!connected}
                >
                    <Ionicons name="chevron-up" size={24} color="white" />
                </TouchableOpacity>

                <View style={styles.horizontalControls}>
                    <TouchableOpacity
                        style={[styles.arrowButton, !connected && styles.disabledButton]}
                        onPress={() => handleManualMove(-10, 0)}
                        disabled={!connected}
                    >
                        <Ionicons name="chevron-back" size={24} color="white" />
                    </TouchableOpacity>

                    <View style={{ width: 50 }} />

                    <TouchableOpacity
                        style={[styles.arrowButton, !connected && styles.disabledButton]}
                        onPress={() => handleManualMove(10, 0)}
                        disabled={!connected}
                    >
                        <Ionicons name="chevron-forward" size={24} color="white" />
                    </TouchableOpacity>
                </View>

                <TouchableOpacity
                    style={[styles.arrowButton, !connected && styles.disabledButton]}
                    onPress={() => handleManualMove(0, 10)}
                    disabled={!connected}
                >
                    <Ionicons name="chevron-down" size={24} color="white" />
                </TouchableOpacity>
            </View>

            <View style={styles.mouseButtonsContainer}>
                <TouchableOpacity
                    style={[styles.mouseButton, !connected && styles.disabledButton]}
                    onPress={() => handleClick('left')}
                    disabled={!connected}
                >
                    <Text style={styles.buttonText}>Left Click</Text>
                </TouchableOpacity>

                <TouchableOpacity
                    style={[styles.mouseButton, !connected && styles.disabledButton]}
                    onPress={() => handleClick('right')}
                    disabled={!connected}
                >
                    <Text style={styles.buttonText}>Right Click</Text>
                </TouchableOpacity>

                <TouchableOpacity
                    style={[styles.mouseButton, !connected && styles.disabledButton]}
                    onPress={() => handleClick('double')}
                    disabled={!connected}
                >
                    <Text style={styles.buttonText}>Double Click</Text>
                </TouchableOpacity>
            </View>
        </>
    );
});

const styles = StyleSheet.create({
    manualControlsContainer: {
        alignItems: 'center',
        justifyContent: 'center',
        padding: 10,
        backgroundColor: 'rgba(33, 150, 243, 0.1)',
        borderRadius: 10,
        marginHorizontal: 15,
    },
    horizontalControls: {
        flexDirection: 'row',
        alignItems: 'center',
        justifyContent: 'center',
        width: '100%',
    },
    arrowButton: {
        backgroundColor: '#2196F3',
        width: 50,
        height: 50,
        borderRadius: 25,
        alignItems: 'center',
        justifyContent: 'center',
        margin: 5,
    },
    mouseButtonsContainer: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        padding: 15,
    },
    mouseButton: {
        flex: 1,
        backgroundColor: '#2196F3',
        padding: 15,
        margin: 5,
        borderRadius: 5,
        alignItems: 'center',
        justifyContent: 'center',
        minHeight: 50,
    },
    buttonText: {
        color: 'white',
        fontWeight: 'bold',
        textAlign: 'center',
    },
    disabledButton: {
        backgroundColor: '#B0BEC5',
        opacity: 0.7,
    },
});

export default Controls; 