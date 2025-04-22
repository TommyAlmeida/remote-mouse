# Native Mouse Control

- macOS: Uses Cocoa/CoreGraphics via CGO - this shit is awful to work with
- Windows: Uses Win32 API via syscall

## API

```go
MoveAbsolute(x, y int)
MoveRelative(deltaX, deltaY int)
LeftClick()
RightClick()
DoubleClick()
width, height := GetScreenSize()
x, y := GetMousePosition()
```

## Implementation

- macOS: Implementation in mouse_darwin.c with header file mouse_darwin.h
- Windows: Implementation directly in mouse_windows.go

## Building

The code uses build constraints to select the right implementation for each platform:

- `mouse_darwin.go` - For macOS (requires Cocoa framework)
- `mouse_windows.go` - For Windows (uses Win32 API)
- linux? maybe once i have more time of this lol
