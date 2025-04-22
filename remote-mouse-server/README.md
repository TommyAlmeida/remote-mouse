# Remote Mouse

## Usage

```sh
go run main.go
```

Run the mobile server and you good to go.

## WebSocket API

Connect to the WebSocket endpoint at `/ws` to control the mouse. Send the following text messages:

### Basic Mouse Movement

```sh
"10,5"     // Move 10px right, 5px down
"-5,0"     // Move 5px left
"0,-10"    // Move 10px up
```

### Mouse Clicks

```sh
"click:left"     // Left click
"click:right"    // Right click
"click:double"   // Double click
```

### Mouse Button Press/Release (for dragging)

```sh
"leftbutton:down"   // Press and hold left button
"leftbutton:up"     // Release left button
"rightbutton:down"  // Press and hold right button
"rightbutton:up"    // Release right button
```

### Configuration Settings

```sh
"config:speed=1.5"     // Set speed multiplier to 1.5
"config:bounds=true"   // Enable screen bounds checking
"config:silent=false"  // Disable silent mode (enable logging)
```

### Stabilization Settings (for drift/jiggle control)

```sh
"stabilize:enable=true"       // Enable stabilization
"stabilize:enable=false"      // Disable stabilization
"stabilize:deadzone=3"        // Set dead zone to 3px 
"stabilize:smoothing=0.5"     // Set smoothing level (0-1)
"stabilize:jiggle=true"       // Enable jiggle filtering
"stabilize:drift=true"        // Enable drift compensation
```
