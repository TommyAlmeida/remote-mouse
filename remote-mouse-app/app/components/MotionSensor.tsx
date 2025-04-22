import React, { memo } from 'react';
import { View, Text, TouchableOpacity, StyleSheet } from 'react-native';
import { MotionSensorProps } from '../types';

const MotionSensor = memo(({
    config,
    connected,
    sensorData,
    mouseDelta,
    resetTracking
}: MotionSensorProps) => {
    return (
        <View style={styles.touchPadContainer}>
            <Text style={styles.touchPadText}>
                {config.useMotionSensor
                    ? 'Move your device to control the mouse'
                    : 'Motion sensor disabled'}
            </Text>

            {config.useMotionSensor && connected && (
                <View style={styles.sensorDataContainer}>
                    <Text style={styles.sensorTitle}>Gyroscope Data:</Text>
                    <Text>
                        X: {sensorData.x.toFixed(2)},
                        Y: {sensorData.y.toFixed(2)},
                        Z: {sensorData.z.toFixed(2)}
                    </Text>

                    <Text style={styles.sensorTitle}>Mouse Movement:</Text>
                    <Text>DeltaX: {mouseDelta.x}, DeltaY: {mouseDelta.y}</Text>

                    <TouchableOpacity
                        style={styles.resetButton}
                        onPress={resetTracking}
                    >
                        <Text style={styles.resetButtonText}>Reset Tracking</Text>
                    </TouchableOpacity>
                </View>
            )}

            {!connected && config.useMotionSensor && (
                <View style={styles.notConnectedContainer}>
                    <Text style={styles.notConnectedText}>
                        Connect to the server to enable motion tracking
                    </Text>
                </View>
            )}
        </View>
    );
});

const styles = StyleSheet.create({
    touchPadContainer: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: '#e0e0e0',
        margin: 15,
        borderRadius: 10,
    },
    touchPadText: {
        fontSize: 16,
        color: '#757575',
        textAlign: 'center',
    },
    sensorDataContainer: {
        marginTop: 20,
        padding: 15,
        backgroundColor: 'rgba(255, 255, 255, 0.7)',
        borderRadius: 5,
        alignItems: 'center',
        width: '90%',
    },
    sensorTitle: {
        fontWeight: 'bold',
        marginTop: 10,
        marginBottom: 5,
    },
    instructionText: {
        marginTop: 15,
        color: '#555',
        textAlign: 'center',
        lineHeight: 22,
    },
    resetButton: {
        backgroundColor: '#FF9800',
        padding: 10,
        borderRadius: 5,
        marginTop: 15,
    },
    resetButtonText: {
        color: 'white',
        fontWeight: 'bold',
    },
    notConnectedContainer: {
        marginTop: 20,
        padding: 15,
        backgroundColor: 'rgba(244, 67, 54, 0.1)',
        borderRadius: 5,
        alignItems: 'center',
        width: '90%',
    },
    notConnectedText: {
        color: '#F44336',
        textAlign: 'center',
    },
});

export default MotionSensor; 