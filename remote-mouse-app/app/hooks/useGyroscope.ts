import { useState, useCallback, useEffect, useRef } from 'react';
import { Gyroscope } from 'expo-sensors';
import { SensorData, MouseDelta, Config } from '../types';

interface UseGyroscopeProps {
    enabled: boolean;
    config: Config;
    connected: boolean;
    sendMessage: (message: string) => boolean;
}

export function useGyroscope({ enabled, config, connected, sendMessage }: UseGyroscopeProps) {
    const [sensorData, setSensorData] = useState<SensorData>({ x: 0, y: 0, z: 0 });
    const [mouseDelta, setMouseDelta] = useState<MouseDelta>({ x: 0, y: 0 });
    const [subscription, setSubscription] = useState<any>(null);

    const cumulativeRotation = useRef<SensorData>({ x: 0, y: 0, z: 0 });
    const lastUpdate = useRef<number>(0);
    const updateInterval = 30;

    const movingBuffer = useRef<SensorData[]>([]);
    const bufferSize = 3; // Reduced from 5 to 3 for less delay

    const deadZone = 0.003; // Reduced from 0.005 for more sensitivity

    const handleGyroscopeData = useCallback((gyroscopeData: SensorData) => {
        setSensorData(gyroscopeData);

        const now = Date.now();
        if (now - lastUpdate.current < updateInterval) return;
        lastUpdate.current = now;

        if (connected && enabled) {
            const filteredData = {
                x: Math.abs(gyroscopeData.x) < deadZone ? 0 : gyroscopeData.x,
                y: gyroscopeData.y,
                z: Math.abs(gyroscopeData.z) < deadZone ? 0 : gyroscopeData.z
            };

            movingBuffer.current.push(filteredData);
            if (movingBuffer.current.length > bufferSize) {
                movingBuffer.current.shift();
            }

            const avgData = movingBuffer.current.reduce(
                (acc, curr) => ({
                    x: acc.x + curr.x / movingBuffer.current.length,
                    y: acc.y + curr.y / movingBuffer.current.length,
                    z: acc.z + curr.z / movingBuffer.current.length
                }),
                { x: 0, y: 0, z: 0 }
            );

            const smoothFactor = 0.8;
            const smoothedX = avgData.x * smoothFactor + cumulativeRotation.current.x * (1 - smoothFactor);
            const smoothedZ = avgData.z * smoothFactor + cumulativeRotation.current.z * (1 - smoothFactor);

            const decayFactor = 0.95;

            cumulativeRotation.current = {
                x: smoothedX * decayFactor,
                y: gyroscopeData.y,
                z: smoothedZ * decayFactor
            };

            const multiplier = 15 * config.speedFactor;

            const jitterThreshold = 0.2 / config.speedFactor;

            const deltaX = Math.abs(smoothedZ) < jitterThreshold ? 0 : Math.round(smoothedZ * -multiplier);
            const deltaY = Math.abs(smoothedX) < jitterThreshold ? 0 : Math.round(smoothedX * -multiplier);

            const significantChange = Math.abs(deltaX) > 0 || Math.abs(deltaY) > 0;
            const newDelta = significantChange ? { x: deltaX, y: deltaY } : mouseDelta;

            if (newDelta.x !== mouseDelta.x || newDelta.y !== mouseDelta.y) {
                setMouseDelta(newDelta);

                if (significantChange) {
                    const message = `${deltaX},${deltaY}`;
                    sendMessage(message);
                }
            }
        }
    }, [connected, enabled, config.speedFactor, mouseDelta, sendMessage]);

    const subscribe = useCallback(() => {
        console.log('[Gyroscope] Subscribing to sensor');
        Gyroscope.setUpdateInterval(updateInterval);

        const subscription = Gyroscope.addListener(handleGyroscopeData);
        setSubscription(subscription);

        return () => {
            console.log('[Gyroscope] Unsubscribing from sensor');
            subscription.remove();
            setSubscription(null);
        };
    }, [handleGyroscopeData]);

    useEffect(() => {
        let cleanup: (() => void) | undefined;

        if (enabled && connected) {
            cleanup = subscribe();
        }

        return () => {
            if (cleanup) cleanup();
        };
    }, [enabled, connected, subscribe]);

    const resetTracking = useCallback(() => {
        cumulativeRotation.current = { x: 0, y: 0, z: 0 };
        movingBuffer.current = [];
        setMouseDelta({ x: 0, y: 0 });
        console.log('[Gyroscope] Tracking reset');
    }, []);

    return {
        sensorData,
        mouseDelta,
        resetTracking,
        isSubscribed: subscription !== null
    };
} 