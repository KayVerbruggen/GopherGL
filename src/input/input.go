package input

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"GopherGL/src/window"
)

var hndl *glfw.Window

// Init takes in the GLFW Window handle, this is needed for handling input.
func Init(w *window.Window) {
	hndl = w.GlfwHandle()
}

// KeyPressed takes in a key and returns whether it is pressed at that moment.
func KeyPressed(k glfw.Key) bool {
	return hndl.GetKey(k) == glfw.Press
}

// KeyReleased takes in a key and returns whether it is released at that moment.
func KeyReleased(k glfw.Key) bool {
	return hndl.GetKey(k) == glfw.Release
}

// MousePos returns the x and y position of the mouse.
func MousePos() (uint32, uint32) {
	x, y := hndl.GetCursorPos()
	return uint32(x), uint32(y)
}

// SetMousePos can be used to change the position of the cursor.
func SetMousePos(x, y uint32) {
	hndl.SetCursorPos(float64(x), float64(y))
}

/*
// MouseScroll returns the amount the mouse has scrolled, this can be negative.
func MouseScroll() int32 {
	hndl.Scroll
}
*/