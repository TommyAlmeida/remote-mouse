#ifndef MOUSE_DARWIN_H
#define MOUSE_DARWIN_H

#include <Cocoa/Cocoa.h>

CGEventRef CreateMouseEvent(CGMouseButton button, CGEventType type, CGPoint position);
void PerformClick(CGMouseButton button, bool doubleClick);
void LeftClick(bool doubleClick);
void RightClick();
void LeftDown();
void LeftUp();
void RightDown();
void RightUp();
void MoveMouseAbsolute(int x, int y);
void MoveMouseRelative(int deltaX, int deltaY);
void GetScreenSize(int* width, int* height);
CGPoint GetMousePosition();

void ReleaseEvent(CGEventRef event);

#endif 