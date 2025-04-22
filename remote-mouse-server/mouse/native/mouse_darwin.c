#include <Cocoa/Cocoa.h>

CGEventRef CreateMouseEvent(CGMouseButton button, CGEventType type, CGPoint position) {
    CGEventRef event = CGEventCreateMouseEvent(NULL, type, position, button);
    return event;
}

void PerformClick(CGMouseButton button, bool doubleClick) {
    CGEventRef event;
    CGPoint currentPos = CGEventGetLocation(CGEventCreate(NULL));
    
    event = CreateMouseEvent(button, kCGEventLeftMouseDown, currentPos);
    CGEventPost(kCGHIDEventTap, event);
    CFRelease(event);
    
    event = CreateMouseEvent(button, kCGEventLeftMouseUp, currentPos);
    CGEventPost(kCGHIDEventTap, event);
    CFRelease(event);
    
    if (doubleClick) {
        event = CreateMouseEvent(button, kCGEventLeftMouseDown, currentPos);
        CGEventPost(kCGHIDEventTap, event);
        CFRelease(event);
        
        event = CreateMouseEvent(button, kCGEventLeftMouseUp, currentPos);
        CGEventPost(kCGHIDEventTap, event);
        CFRelease(event);
    }
}

void LeftClick(bool doubleClick) {
    PerformClick(kCGMouseButtonLeft, doubleClick);
}

void LeftDown() {
    CGPoint currentPos = CGEventGetLocation(CGEventCreate(NULL));
    CGEventRef event = CreateMouseEvent(kCGMouseButtonLeft, kCGEventLeftMouseDown, currentPos);
    CGEventPost(kCGHIDEventTap, event);
    CFRelease(event);
}

void LeftUp() {
    CGPoint currentPos = CGEventGetLocation(CGEventCreate(NULL));
    CGEventRef event = CreateMouseEvent(kCGMouseButtonLeft, kCGEventLeftMouseUp, currentPos);
    CGEventPost(kCGHIDEventTap, event);
    CFRelease(event);
}

void RightClick() {
    CGPoint currentPos = CGEventGetLocation(CGEventCreate(NULL));
    
    CGEventRef event = CreateMouseEvent(kCGMouseButtonRight, kCGEventRightMouseDown, currentPos);
    CGEventPost(kCGHIDEventTap, event);
    CFRelease(event);
    
    event = CreateMouseEvent(kCGMouseButtonRight, kCGEventRightMouseUp, currentPos);
    CGEventPost(kCGHIDEventTap, event);
    CFRelease(event);
}

void RightDown() {
    CGPoint currentPos = CGEventGetLocation(CGEventCreate(NULL));
    CGEventRef event = CreateMouseEvent(kCGMouseButtonRight, kCGEventRightMouseDown, currentPos);
    CGEventPost(kCGHIDEventTap, event);
    CFRelease(event);
}

void RightUp() {
    CGPoint currentPos = CGEventGetLocation(CGEventCreate(NULL));
    CGEventRef event = CreateMouseEvent(kCGMouseButtonRight, kCGEventRightMouseUp, currentPos);
    CGEventPost(kCGHIDEventTap, event);
    CFRelease(event);
}

void MoveMouseAbsolute(int x, int y) {
    CGPoint point = CGPointMake(x, y);
    CGEventRef event = CGEventCreateMouseEvent(NULL, kCGEventMouseMoved, point, kCGMouseButtonLeft);
    CGEventPost(kCGHIDEventTap, event);
    CFRelease(event);
}

void MoveMouseRelative(int deltaX, int deltaY) {
    CGEventRef event = CGEventCreate(NULL);
    CGPoint currentPos = CGEventGetLocation(event);
    CFRelease(event);
    
    CGPoint newPos = CGPointMake(currentPos.x + deltaX, currentPos.y + deltaY);
    event = CGEventCreateMouseEvent(NULL, kCGEventMouseMoved, newPos, kCGMouseButtonLeft);
    CGEventPost(kCGHIDEventTap, event);
    CFRelease(event);
}

void GetScreenSize(int* width, int* height) {
    CGDirectDisplayID displayID = CGMainDisplayID();
    *width = (int)CGDisplayPixelsWide(displayID);
    *height = (int)CGDisplayPixelsHigh(displayID);
}

CGPoint GetMousePosition() {
    CGEventRef event = CGEventCreate(NULL);
    CGPoint point = CGEventGetLocation(event);
    CFRelease(event);
    return point;
}

// Helper function to release a CGEventRef, handling the cast to CFTypeRef properly
void ReleaseEvent(CGEventRef event) {
    if (event) {
        CFRelease((CFTypeRef)event);
    }
} 