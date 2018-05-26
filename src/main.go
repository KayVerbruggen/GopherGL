package main

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/glfw/v3.2/glfw"

	"GopherGL/src/camera"
	"GopherGL/src/gfx"
	"GopherGL/src/window"
	"GopherGL/src/input"
)

// check is used to check errors.
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// HandleInput handles all the input. This function will move into a seperate input package later.
func handleInput(w *window.Window, movSpd float32, cam *camera.Camera) {
	// Close the window.
	if input.KeyPressed(glfw.KeyEscape) {
		w.GlfwHandle().SetShouldClose(true)
	}

	// Move forwards.
	if input.KeyPressed(glfw.KeyW) {
		cam.Pos = mgl32.Vec3{cam.Pos.X(), cam.Pos.Y(), cam.Pos.Z() - movSpd*w.DeltaTime()}
	}

	// Move backwards.
	if input.KeyPressed(glfw.KeyS) {
		cam.Pos = mgl32.Vec3{cam.Pos.X(), cam.Pos.Y(), cam.Pos.Z() + movSpd*w.DeltaTime()}
	}

	// Move left.
	if input.KeyPressed(glfw.KeyA) {
		cam.Pos = mgl32.Vec3{cam.Pos.X() - movSpd*w.DeltaTime(), cam.Pos.Y(), cam.Pos.Z()}
	}

	// Move right.
	if input.KeyPressed(glfw.KeyD) {
		cam.Pos = mgl32.Vec3{cam.Pos.X() + movSpd*w.DeltaTime(), cam.Pos.Y(), cam.Pos.Z()}
	}
}

func main() {
	window, err := window.CreateWindow(800, 600, "GopherGL", true)
	check(err)

	gfx.InitRenderer()

	input.Init(window)
	cam := camera.CreateCamera(mgl32.Vec3{0.0, 0.0, 3.0}, float32(window.X)/float32(window.Y), 90.0)

	// This is the sun of the scene.
	sun := gfx.CreateDirectionalLight(mgl32.Vec3{0.5, -0.5, 0.0}, 1.0)

	cubeMat := gfx.CreateMaterial("../res/containerTex.png", "../res/containerSpec.png", 1.0)
	cube := gfx.CreateCube(0.0, 0.0, 0.0, 0.0, 0.0, 0.0, cubeMat)

	// TODO: This should be handled differently. Most of it can be done when creating the objects.
	// Set uniform.

	for window.IsOpen() {
		handleInput(window, 3.0, cam)

		// Keeps the aspect ratio correct
		if window.AspectChanged() {
			cam.SetProjection(float32(window.X)/float32(window.Y), 90.0)
		}

		cam.Update()
		
		// Rotate dirt cube.
		cube.SetRot(window.Time(), 0.0, window.Time())
		
		// OpenGL stuff.
		gfx.BeginFrame()
		gfx.Render(cam, cube, sun)

		window.Update()
	}

	window.Close()
}
