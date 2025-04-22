import { useState, useCallback, useEffect, useRef } from 'react';
import { Alert } from 'react-native';
import { Config } from '../types';

export function useWebSocket(config: Config) {
    const [connected, setConnected] = useState(false);
    const [ws, setWs] = useState<WebSocket | null>(null);
    const reconnectAttempts = useRef(0);
    const maxReconnectAttempts = 3;
    const lastConfigUpdateTime = useRef<Record<string, number>>({});
    const configUpdateDebounceTime = 300;

    const connect = useCallback(() => {
        try {
            console.log(`[WebSocket] Connecting to ${config.serverUrl}...`);
            const socket = new WebSocket(config.serverUrl);

            socket.onopen = () => {
                console.log('[WebSocket] Connection established');
                setConnected(true);
                setWs(socket);
                reconnectAttempts.current = 0;

                try {
                    socket.send(`config:speed=${config.speedFactor}`);
                    console.log(`[WebSocket] Sent config:speed=${config.speedFactor}`);
                    socket.send(`config:bounds=${config.enforceBounds}`);
                    console.log(`[WebSocket] Sent config:bounds=${config.enforceBounds}`);
                } catch (error) {
                    console.error('[WebSocket] Error sending initial config:', error);
                }
            };

            socket.onclose = (event) => {
                console.log(`[WebSocket] Connection closed: ${event.code} ${event.reason}`);
                setConnected(false);
                setWs(null);
            };

            socket.onerror = (error) => {
                console.error('[WebSocket] Connection error:', error);
                if (!connected) {
                    Alert.alert('Connection Error', 'Failed to connect to server');
                }
                setConnected(false);
                setWs(null);
            };
        } catch (error) {
            console.error('[WebSocket] Error creating connection:', error);
            Alert.alert('Connection Error', 'Failed to connect to server');
        }
    }, [config.serverUrl, config.speedFactor, config.enforceBounds, connected]);

    // Disconnect from WebSocket server
    const disconnect = useCallback(() => {
        if (ws) {
            ws.close();
            setWs(null);
            setConnected(false);
        }
    }, [ws]);

    // Send message to server
    const sendMessage = useCallback((message: string) => {
        if (connected && ws) {
            try {
                ws.send(message);
                return true;
            } catch (error) {
                console.error('[WebSocket] Error sending message:', error);
                return false;
            }
        }
        return false;
    }, [connected, ws]);

    const updateServerConfig = useCallback((key: string, value: any) => {
        const now = Date.now();
        const lastUpdate = lastConfigUpdateTime.current[key] || 0;

        if (now - lastUpdate > configUpdateDebounceTime) {
            let message = '';
            if (key === 'speedFactor') {
                message = `config:speed=${value}`;
            } else if (key === 'enforceBounds') {
                message = `config:bounds=${value}`;
            } else {
                return; // Not a server-supported config
            }

            try {
                const success = sendMessage(message);
                if (success) {
                    lastConfigUpdateTime.current[key] = now;
                    console.log(`[WebSocket] Sent ${message}`);
                }
            } catch (error) {
                console.error(`[WebSocket] Error sending config update for ${key}:`, error);
            }
        }
    }, [sendMessage]);

    // Auto-reconnect on server errors, with backoff
    useEffect(() => {
        const handleReconnect = () => {
            if (!connected && reconnectAttempts.current < maxReconnectAttempts) {
                const timeout = Math.pow(2, reconnectAttempts.current) * 1000;
                console.log(`[WebSocket] Attempting to reconnect in ${timeout}ms (attempt ${reconnectAttempts.current + 1}/${maxReconnectAttempts})`);

                setTimeout(() => {
                    reconnectAttempts.current += 1;
                    connect();
                }, timeout);
            }
        };

        // If we lose connection unexpectedly (ws was set but then unset), try to reconnect
        const wasConnected = ws !== null;
        const isConnected = connected && ws !== null;

        if (wasConnected && !isConnected) {
            handleReconnect();
        }

        return () => {
            if (ws) {
                ws.close();
            }
        };
    }, [ws, connected, connect]);

    return {
        connected,
        connect,
        disconnect,
        sendMessage,
        updateServerConfig
    };
} 