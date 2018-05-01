package main

import (
	"github.com/go-gl/mathgl/mgl32"

	"GopherGL/src/camera"
	"GopherGL/src/gfx"
	"GopherGL/src/window"
)

// check is used to check errors.
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	window, err := window.CreateWindow(800, 600, "GopherGL", true)
	check(err)

	renderer := gfx.Renderer{}
	cam := camera.CreateCamera(mgl32.Vec3{0.0, 0.0, 3.0}, float32(window.X)/float32(window.Y), 90.0)

	// Set the correct shader and texture.
	shader := gfx.CreateShader("../shaders/basic.glsl")

	// Position of the light.
	lightPos := mgl32.Vec3{-1.0, 0.0, 1.0}

	cubeMat := gfx.CreateMaterial("../res/containerTex.png", "../res/containerSpec.png", 1.0)
	shader.SetUniformFloat("mat.shininess", cubeMat.Shininess)
	cube := gfx.CreateCube(0.0, 0.0, 0.0, 0.0, 0.0, 0.0, cubeMat)

	// TODO: This should be handled differently. Most of it can be done when creating the objects.
	// Set uniform.
	shader.SetUniformMat4("model", cube.Trans)
	shader.SetUniformVec3("lightPos", lightPos)

	for window.IsOpen() {
		window.HandleInput(3.0, cam)

		// TODO: For better performance, do NOT recalculate the projection matrix each frame, but only when changed.
		// Keeps the aspect ratio correct
		if window.AspectChanged() {
			cam.SetProjection(float32(window.X)/float32(window.Y), 90.0)
		}

		cam.Update(shader)

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
