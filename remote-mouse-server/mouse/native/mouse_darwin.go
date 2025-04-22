//go:build darwin
// +build darwin

package native

import (
	"sync"
)

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#include "mouse_darwin.h"
*/
import "C"

var (
	mouseMutex sync.Mutex
)

func MoveAbsolute(x, y int) {
	mouseMutex.Lock()
	defer mouseMutex.Unlock()
	
	C.MoveMouseAbsolute(C.int(x), C.int(y))
}

func MoveRelative(deltaX, deltaY int) {
	mouseMutex.Lock()
	defer mouseMutex.Unlock()
	
	C.MoveMouseRelative(C.int(deltaX), C.int(deltaY))
}

func LeftClick() {
	mouseMutex.Lock()
	defer mouseMutex.Unlock()
	
	C.LeftClick(C.bool(false))
}

func LeftDown() {
	mouseMutex.Lock()
	defer mouseMutex.Unlock()
	
	pos := C.GetMousePosition()
	event := C.CreateMouseEvent(C.kCGMouseButtonLeft, C.kCGEventLeftMouseDown, pos)
	C.CGEventPost(C.kCGHIDEventTap, event)
	C.ReleaseEvent(event)
}

func LeftUp() {
	mouseMutex.Lock()
	defer mouseMutex.Unlock()
	
	pos := C.GetMousePosition()
	event := C.CreateMouseEvent(C.kCGMouseButtonLeft, C.kCGEventLeftMouseUp, pos)
	C.CGEventPost(C.kCGHIDEventTap, event)
	C.ReleaseEvent(event)
}

func RightClick() {
	mouseMutex.Lock()
	defer mouseMutex.Unlock()
	
	C.RightClick()
}

func RightDown() {
	mouseMutex.Lock()
	defer mouseMutex.Unlock()
	
	pos := C.GetMousePosition()
	event := C.CreateMouseEvent(C.kCGMouseButtonRight, C.kCGEventRightMouseDown, pos)
	C.CGEventPost(C.kCGHIDEventTap, event)
	C.ReleaseEvent(event)
}

func RightUp() {
	mouseMutex.Lock()
	defer mouseMutex.Unlock()
	
	pos := C.GetMousePosition()
	event := C.CreateMouseEvent(C.kCGMouseButtonRight, C.kCGEventRightMouseUp, pos)
	C.CGEventPost(C.kCGHIDEventTap, event)
	C.ReleaseEvent(event)
}

func DoubleClick() {
	mouseMutex.Lock()
	defer mouseMutex.Unlock()
	
	C.LeftClick(C.bool(true))
}

func GetScreenSize() (width, height int) {
	var w, h C.int
	C.GetScreenSize(&w, &h)
	return int(w), int(h)
}

func GetMousePosition() (x, y int) {
	pos := C.GetMousePosition()
	return int(pos.x), int(pos.y)
} 