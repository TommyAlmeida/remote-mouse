import React, { useState, useCallback, useMemo } from 'react';
import { View, StyleSheet } from 'react-native';
import { StatusBar } from 'expo-status-bar';

import { useWebSocket } from './hooks/useWebSocket';
import { useGyroscope } from './hooks/useGyroscope';

import Header from './components/Header';
import ConnectionStatus from './components/ConnectionStatus';
import Settings from './components/Settings';
import MotionSensor from './components/MotionSensor';
import Controls from './components/Controls';

import { Config } from './types';

export default function Index() {
  const [showSettings, setShowSettings] = useState(false);

  const serverUrl = process.env.WS_SERVER_URL || 'ws://192.168.1.94:8080/ws';
  const [config, setConfig] = useState<Config>({
    serverUrl,
    speedFactor: 2.0,
    enforceBounds: true,
    useMotionSensor: true,
    showControlPad: true,
  });

  const { connected, connect, disconnect, sendMessage, updateServerConfig } = useWebSocket(config);

  const updateConfig = useCallback((key: keyof Config, value: any) => {
    setConfig(prev => ({ ...prev, [key]: value }));

    if (connected && (key === 'speedFactor' || key === 'enforceBounds')) {
      updateServerConfig(key, value);
    }
  }, [connected, updateServerConfig]);

  const { sensorData, mouseDelta, resetTracking } = useGyroscope({
    enabled: config.useMotionSensor,
    config,
    connected,
    sendMessage
  });

  const connectionProps = useMemo(() => ({
    connected,
    connect,
    disconnect
  }), [connected, connect, disconnect]);

  const settingsProps = useMemo(() => ({
    config,
    updateConfig,
    visible: showSettings
  }), [config, updateConfig, showSettings]);

  const motionSensorProps = useMemo(() => ({
    config,
    connected,
    sensorData,
    mouseDelta,
    resetTracking
  }), [config, connected, sensorData, mouseDelta, resetTracking]);

  const controlsProps = useMemo(() => ({
    connected,
    sendMessage,
    visible: config.showControlPad
  }), [connected, sendMessage, config.showControlPad]);

  return (
    <View style={styles.container}>
      <StatusBar style="auto" />

      <Header
        title="Remote Mouse"
        onSettingsPress={() => setShowSettings(!showSettings)}
      />

      <ConnectionStatus {...connectionProps} />

      <Settings {...settingsProps} />

      <MotionSensor {...motionSensorProps} />

      <Controls {...controlsProps} />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#f5f5f5',
  },
});
