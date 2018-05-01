package main

import (
	"github.com/go-gl/mathgl/mgl32"

	"GopherGL/src/camera"
	"GopherGL/src/gfx"
	"GopherGL/src/window"
)

var (
	projChanged = false
	// Create a camera.
	cam                   = camera.CreateCamera(mgl32.Vec3{0.0, 0.0, 3.0}, 70)
	movementSpeed float32 = 3.0
)

// check is used to check errors.
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	window := window.CreateWindow(800, 600, "GopherGL", true)
	renderer := gfx.Renderer{}

	// Set the correct shader and texture.
	shader := gfx.CreateShader("../shaders/basic.glsl")

	// Create a projection.
	proj := mgl32.Perspective(mgl32.DegToRad(cam.Fov), float32(window.X)/float32(window.Y), 0.1, 1000.0)
	// Position of the light.
	lightPos := mgl32.Vec3{-1.0, 0.0, 1.0}

	cubeMat := gfx.CreateMaterial("../res/containerTex.png", "../res/containerSpec.png", 1.0)
	shader.SetUniformFloat("mat.shininess", cubeMat.Shininess)
	cube := gfx.CreateCube(0.0, 0.0, 0.0, 0.0, 0.0, 0.0, cubeMat)

	// TODO: This should be handled differently. Most of it can be done when creating the objects.
	// Set uniform.
	shader.SetUniformMat4("model", cube.Trans)
	shader.SetUniformMat4("projection", proj)

	shader.SetUniformVec3("lightPos", lightPos)
	shader.SetUniformVec3("viewPos", cam.Pos)

	shader.SetUniformInt32("mat.diffTex", 0)
	shader.SetUniformInt32("mat.specTex", 1)

	// TODO: Make some sort of time package for this delta time stuff.

	for window.IsOpen() {
		window.HandleInput(movementSpeed, cam)
		cam.UpdateCamera(shader)

		// TODO: For better performance, do NOT recalculate the projection matrix each frame, but only when changed.
		// Keeps the aspect ratio correct
		proj := mgl32.Perspective(mgl32.DegToRad(cam.Fov), float32(window.X)/float32(window.Y), 0.1, 1000.0)
		shader.SetUniformMat4("projection", proj)

		// Rotate dirt cube.
		cube.SetRot(window.Time(), 0.0, window.Time())
		shader.SetUniformMat4("model", cube.Trans)

		// OpenGL stuff.
		renderer.BeginFrame()
		renderer.Render(cube)

		window.Update()
	}

	window.Close()
}
