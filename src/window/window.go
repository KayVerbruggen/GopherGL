// Package window is pretty much just a wrapper around GLFW.
package window

import (
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"

	"GopherGL/src/camera"
)

var (
	currTime, prevTime, deltaTime float32
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Window stores a glfw handle, x and y dimensions and if vsync is enables.
type Window struct {
	X, Y          uint32
	vsync         bool
	handle        *glfw.Window
	aspectChanged bool
}

func (w *Window) resizeCallback(glfwWin *glfw.Window, x, y int) {
	gl.Viewport(0, 0, int32(x), int32(y))

	// This will keep the aspect ratio correct.
	w.X = uint32(x)
	w.Y = uint32(y)
	w.aspectChanged = true
}

// AspectChanged returns whether or not the aspect ratio has been changed.
func (w *Window) AspectChanged() bool {
	return w.aspectChanged
}

// GlfwHandle returns the exactly that.
func (w *Window) GlfwHandle() *glfw.Window {
	return w.handle
}

// SetScrollCallback takes in a function which will run when scrolling.
func (w *Window) SetScrollCallback(scb glfw.ScrollCallback) {
	w.handle.SetScrollCallback(scb)
}

// CreateWindow initializes GLFW, creates a OpenGL context and opens a window.
func CreateWindow(x, y uint32, name string, vsync bool) (*Window, error) {
	// Lock the goroutine to one thread, because OpenGL doesn't support multithreading nor switching threads.
	runtime.LockOSThread()

	w := &Window{}
	w.X = x
	w.Y = y
	w.vsync = vsync

	err := glfw.Init()
	if err != nil {
		return nil, err
	}

	// Set the correct window hints.
	glfw.WindowHint(glfw.ContextVersionMajor, 3)                // OpenGL version 3.3
	glfw.WindowHint(glfw.ContextVersionMinor, 3)                // OpenGL version 3.3
	glfw.WindowHint(glfw.Resizable, glfw.True)                  // It is resizable.
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)    // Needs to be true to work on macOS.
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile) // Core profile to prevent using deprecated stuff.

	w.handle, err = glfw.CreateWindow(int(x), int(y), name, nil, nil)
	if err != nil {
		return nil, err
	}
	w.handle.MakeContextCurrent()

	// Enable V-Sync to not use 18% CPU for a simple Gopher.
	if vsync {
		glfw.SwapInterval(1)
	}

	err = gl.Init()
	if err != nil {
		return nil, err
	}
	gl.Viewport(0, 0, int32(w.X), int32(w.Y))
	// This allows us to use transparency with PNG files.
	gl.Enable(gl.BLEND)
	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.CULL_FACE)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	// Just some information for the users.
	fmt.Println("OS:", runtime.GOOS, "\nArchitecture:", runtime.GOARCH)
	fmt.Println("OpenGL:", gl.GoStr(gl.GetString(gl.VERSION)))

	// Set the GLFW callbacks.
	w.handle.SetFramebufferSizeCallback(w.resizeCallback)

	return w, nil
}

// Update swaps the buffers and checks for glfw events.
func (w *Window) Update() {
	prevTime = currTime
	currTime = float32(glfw.GetTime())
	deltaTime = currTime - prevTime

	w.handle.SwapBuffers()
	glfw.PollEvents()
}

// IsOpen returns if the window is still opened.
func (w *Window) IsOpen() bool {
	return !w.handle.ShouldClose()
}

// Close closes the window and terminates GLFW.
func (w *Window) Close() {
	w.handle.SetShouldClose(true)
	glfw.Terminate()
}

// DeltaTime returns the time it took to render the frame. This is 1/FPS.
func (w *Window) DeltaTime() float32 {
	return deltaTime
}

// FPS returns the frames per second.
func (w *Window) FPS() float32 {
	return 1.0 / deltaTime
}

// Time returns the time since GLFW has been initialized.
func (w *Window) Time() float32 {
	return float32(glfw.GetTime())
}

// HandleInput handles all the input. This function will move into a seperate input package later.
func (w *Window) HandleInput(movSpd float32, cam *camera.Camera) {
	// Close the window.
	if w.handle.GetKey(glfw.KeyEscape) == glfw.Press {
		w.handle.SetShouldClose(true)
	}

	if w.handle.GetKey(glfw.KeyEscape) == glfw.Press {
		w.handle.SetShouldClose(true)
	}

	// Move forwards.
	if w.handle.GetKey(glfw.KeyW) == glfw.Press {
		cam.Pos = mgl32.Vec3{cam.Pos.X(), cam.Pos.Y(), cam.Pos.Z() - movSpd*deltaTime}
	}

	// Move backwards.
	if w.handle.GetKey(glfw.KeyS) == glfw.Press {
		cam.Pos = mgl32.Vec3{cam.Pos.X(), cam.Pos.Y(), cam.Pos.Z() + movSpd*deltaTime}
	}

	// Move left.
	if w.handle.GetKey(glfw.KeyA) == glfw.Press {
		cam.Pos = mgl32.Vec3{cam.Pos.X() - movSpd*deltaTime, cam.Pos.Y(), cam.Pos.Z()}
	}

	// Move right.
	if w.handle.GetKey(glfw.KeyD) == glfw.Press {
		cam.Pos = mgl32.Vec3{cam.Pos.X() + movSpd*deltaTime, cam.Pos.Y(), cam.Pos.Z()}
	}
}
