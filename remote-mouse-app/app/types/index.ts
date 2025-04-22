// Config interface for app settings
export interface Config {
    serverUrl: string;
    speedFactor: number;
    enforceBounds: boolean;
    useMotionSensor: boolean;
    showControlPad: boolean;
}

// Sensor data from gyroscope
export interface SensorData {
    x: number;
    y: number;
    z: number;
}

// Mouse movement deltas
export interface MouseDelta {
    x: number;
    y: number;
}

// Props for component sharing
export interface ConnectionProps {
    connected: boolean;
    connect: () => void;
    disconnect: () => void;
}

export interface SettingsProps {
    config: Config;
    updateConfig: (key: keyof Config, value: any) => void;
    visible: boolean;
}

export interface MotionSensorProps {
    config: Config;
    connected: boolean;
    sensorData: SensorData;
    mouseDelta: MouseDelta;
    resetTracking: () => void;
}

export interface ControlsProps {
    connected: boolean;
    sendMessage: (message: string) => boolean;
    visible: boolean;
} 