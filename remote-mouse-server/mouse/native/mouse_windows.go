//go:build windows
// +build windows

package native

import (
	"sync"
	"syscall"
	"unsafe"
)

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	procGetSystemMetrics = user32.NewProc("GetSystemMetrics")
	procSetCursorPos     = user32.NewProc("SetCursorPos")
	procGetCursorPos     = user32.NewProc("GetCursorPos")
	procMouseEvent       = user32.NewProc("mouse_event")
	
	mouseMutex sync.Mutex
)

const (
	smCxScreen = 0
	smCyScreen = 1

	mouseeventLeftdown   = 0x0002
	mouseeventLeftup     = 0x0004
	mouseeventRightdown  = 0x0008
	mouseeventRightup    = 0x0010
	mouseeventAbsolute   = 0x8000
	mouseeventMove       = 0x0001
)

// POINT represents a point structure from Win32 API
type POINT struct {
	X int32
	Y int32
}

// MoveAbsolute moves the mouse cursor to the specified absolute coordinates
func MoveAbsolute(x, y int) {
	mouseMutex.Lock()
	defer mouseMutex.Unlock()
	
	procSetCursorPos.Call(uintptr(x), uintptr(y))
}

// MoveRelative moves the mouse cursor by the specified delta values
func MoveRelative(deltaX, deltaY int) {
	mouseMutex.Lock()
	defer mouseMutex.Unlock()
	
	var point POINT
	procGetCursorPos.Call(uintptr(unsafe.Pointer(&point)))
	
	newX := int(point.X) + deltaX
	newY := int(point.Y) + deltaY
	
	procSetCursorPos.Call(uintptr(newX), uintptr(newY))
}

// LeftClick performs a left mouse button click
func LeftClick() {
	mouseMutex.Lock()
	defer mouseMutex.Unlock()
	
	procMouseEvent.Call(
		uintptr(mouseeventLeftdown),
		0,
		0,
		0,
		0,
	)
	procMouseEvent.Call(
		uintptr(mouseeventLeftup),
		0,
		0,
		0,
		0,
	)
}

// LeftDown performs a left mouse button press
func LeftDown() {
	mouseMutex.Lock()
	defer mouseMutex.Unlock()
	
	procMouseEvent.Call(
		uintptr(mouseeventLeftdown),
		0,
		0,
		0,
		0,
	)
}

// LeftUp performs a left mouse button release
func LeftUp() {
	mouseMutex.Lock()
	defer mouseMutex.Unlock()
	
	procMouseEvent.Call(
		uintptr(mouseeventLeftup),
		0,
		0,
		0,
		0,
	)
}

// RightClick performs a right mouse button click
func RightClick() {
	mouseMutex.Lock()
	defer mouseMutex.Unlock()
	
	procMouseEvent.Call(
		uintptr(mouseeventRightdown),
		0,
		0,
		0,
		0,
	)
	procMouseEvent.Call(
		uintptr(mouseeventRightup),
		0,
		0,
		0,
		0,
	)
}

// RightDown performs a right mouse button press
func RightDown() {
	mouseMutex.Lock()
	defer mouseMutex.Unlock()
	
	procMouseEvent.Call(
		uintptr(mouseeventRightdown),
		0,
		0,
		0,
		0,
	)
}

// RightUp performs a right mouse button release
func RightUp() {
	mouseMutex.Lock()
	defer mouseMutex.Unlock()
	
	procMouseEvent.Call(
		uintptr(mouseeventRightup),
		0,
		0,
		0,
		0,
	)
}

// DoubleClick performs a double click with the left mouse button
func DoubleClick() {
	mouseMutex.Lock()
	defer mouseMutex.Unlock()
	
	procMouseEvent.Call(
		uintptr(mouseeventLeftdown),
		0,
		0,
		0,
		0,
	)
	procMouseEvent.Call(
		uintptr(mouseeventLeftup),
		0,
		0,
		0,
		0,
	)
	
	procMouseEvent.Call(
		uintptr(mouseeventLeftdown),
		0,
		0,
		0,
		0,
	)
	procMouseEvent.Call(
		uintptr(mouseeventLeftup),
		0,
		0,
		0,
		0,
	)
}

// GetScreenSize returns the primary display dimensions
func GetScreenSize() (width, height int) {
	w, _, _ := procGetSystemMetrics.Call(uintptr(smCxScreen))
	h, _, _ := procGetSystemMetrics.Call(uintptr(smCyScreen))
	return int(w), int(h)
}

// GetMousePosition returns the current mouse cursor position
func GetMousePosition() (x, y int) {
	var point POINT
	procGetCursorPos.Call(uintptr(unsafe.Pointer(&point)))
	return int(point.X), int(point.Y)
} 