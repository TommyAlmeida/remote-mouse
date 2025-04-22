import React, { memo, useCallback, useRef, useState } from 'react';
import { View, Text, TextInput, Switch, StyleSheet } from 'react-native';
import Slider from '@react-native-community/slider';
import { SettingsProps } from '../types';

const Settings = memo(({ config, updateConfig, visible }: SettingsProps) => {
    if (!visible) return null;

    const [localSpeed, setLocalSpeed] = useState(config.speedFactor);
    const sliderMoving = useRef(false);

    const handleSliderChange = useCallback((value: number) => {
        sliderMoving.current = true;
        setLocalSpeed(value);
    }, []);

    const handleSliderComplete = useCallback((value: number) => {
        sliderMoving.current = false;
        updateConfig('speedFactor', value);
    }, [updateConfig]);

    const handleSpeedInputChange = useCallback((text: string) => {
        const value = parseFloat(text);
        if (!isNaN(value) && value >= 0.5 && value <= 5.0) {
            setLocalSpeed(value);
            updateConfig('speedFactor', value);
        }
    }, [updateConfig]);

    const handleTextChange = useCallback((value: string) => {
        updateConfig('serverUrl', value);
    }, [updateConfig]);

    return (
        <View style={styles.settingsPanel}>
            <Text style={styles.settingTitle}>Server URL</Text>
            <TextInput
                style={styles.input}
                value={config.serverUrl}
                onChangeText={handleTextChange}
                placeholder="ws://server-ip:port/ws"
                autoCapitalize="none"
                autoCorrect={false}
            />

            <Text style={styles.settingTitle}>
                Mouse Speed: {localSpeed.toFixed(1)}
            </Text>
            <View style={styles.sliderContainer}>
                <Text>0.5</Text>
                <Slider
                    style={styles.slider}
                    minimumValue={0.5}
                    maximumValue={5.0}
                    step={0.1}
                    value={localSpeed}
                    onValueChange={handleSliderChange}
                    onSlidingComplete={handleSliderComplete}
                    minimumTrackTintColor="#2196F3"
                    maximumTrackTintColor="#BDBDBD"
                    thumbTintColor="#2196F3"
                />
                <Text>5.0</Text>
            </View>
            <TextInput
                style={styles.sliderInput}
                value={localSpeed.toString()}
                onChangeText={handleSpeedInputChange}
                keyboardType="numeric"
            />

            <View style={styles.switchRow}>
                <Text style={styles.settingTitle}>Enforce Screen Bounds</Text>
                <Switch
                    value={config.enforceBounds}
                    onValueChange={(value) => updateConfig('enforceBounds', value)}
                />
            </View>

            <View style={styles.switchRow}>
                <Text style={styles.settingTitle}>Use Motion Sensor</Text>
                <Switch
                    value={config.useMotionSensor}
                    onValueChange={(value) => updateConfig('useMotionSensor', value)}
                />
            </View>

            <View style={styles.switchRow}>
                <Text style={styles.settingTitle}>Show Control Pad</Text>
                <Switch
                    value={config.showControlPad}
                    onValueChange={(value) => updateConfig('showControlPad', value)}
                />
            </View>
        </View>
    );
});

const styles = StyleSheet.create({
    settingsPanel: {
        backgroundColor: 'white',
        padding: 15,
        borderBottomWidth: 1,
        borderBottomColor: '#e0e0e0',
    },
    settingTitle: {
        fontSize: 16,
        marginBottom: 5,
        marginTop: 10,
    },
    input: {
        borderWidth: 1,
        borderColor: '#e0e0e0',
        borderRadius: 5,
        padding: 10,
        fontSize: 16,
    },
    sliderContainer: {
        flexDirection: 'row',
        alignItems: 'center',
        justifyContent: 'space-between',
        marginBottom: 10,
    },
    slider: {
        flex: 1,
        marginHorizontal: 10,
        height: 40,
    },
    sliderInput: {
        borderWidth: 1,
        borderColor: '#e0e0e0',
        borderRadius: 5,
        padding: 10,
        width: 80,
        textAlign: 'center',
        alignSelf: 'center',
        marginBottom: 10,
    },
    switchRow: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: 'center',
        marginVertical: 5,
    },
});

export default Settings; 