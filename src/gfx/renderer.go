// Package gfx is the graphics part of the game engine I'm developing.
package gfx

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"GopherGL/src/camera"
)

// InitRenderer currently only sets up the shaders.
func InitRenderer() {
	pointShader = createShader("../shaders/point.glsl")
	directionalShader = createShader("../shaders/directional.glsl")
	ambientShader = createShader("../shaders/ambient.glsl")
	basicShader = createShader("../shaders/basic.glsl")
}

// BeginFrame clears the screen, do this before rendering.
func BeginFrame() {
	// The background color.
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.Clear(gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(0.2, 0.3, 0.3, 1.0)
}

// Render takes in an Entity and draws it to the framebuffer.
func Render(c *camera.Camera, e *Entity, dl *DirectionalLight) {
	// Set the texture and specular lighting map.
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, e.mat.texID)
	gl.ActiveTexture(gl.TEXTURE1)
	gl.BindTexture(gl.TEXTURE_2D, e.mat.specID)
	basicShader.SetUniformFloat("mat.shininess", e.mat.Shininess)
	basicShader.SetUniformInt32("mat.diffTex", 0)
	basicShader.SetUniformInt32("mat.specTex", 1)
	
	basicShader.SetUniformMat4("model", e.Trans)
	basicShader.SetUniformMat4("view", c.View)
	basicShader.SetUniformVec3("viewPos", c.Pos)
	basicShader.SetUniformMat4("projection", c.Proj)
	
	basicShader.SetUniformFloat("sun.intensity", dl.intensity)
	basicShader.SetUniformVec3("sun.direction", dl.dir)
	
	gl.UseProgram(basicShader.program)
	gl.BindVertexArray(e.vao)
	gl.DrawElements(gl.TRIANGLES, e.size, gl.UNSIGNED_INT, gl.Ptr(nil))
	
	/*
	// TODO: There should be deferred shading instead of this mess.
	ambientShader.SetUniformFloat("mat.shininess", e.mat.Shininess)
	ambientShader.SetUniformMat4("model", e.Trans)
	ambientShader.SetUniformMat4("view", c.View)
	ambientShader.SetUniformMat4("projection", c.Proj)
	
	gl.UseProgram(ambientShader.program)
	gl.BindVertexArray(e.vao)
	gl.DrawElements(gl.TRIANGLES, e.size, gl.UNSIGNED_INT, gl.Ptr(nil))
	
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.ONE, gl.ONE)
	gl.DepthMask(false)
	gl.DepthFunc(gl.EQUAL)
	
	directionalShader.SetUniformFloat("mat.shininess", e.mat.Shininess)
	directionalShader.SetUniformInt32("mat.diffTex", 0)
	directionalShader.SetUniformInt32("mat.specTex", 1)
	
	directionalShader.SetUniformMat4("model", e.Trans)
	directionalShader.SetUniformMat4("view", c.View)
	directionalShader.SetUniformVec3("viewPos", c.Pos)
	directionalShader.SetUniformMat4("projection", c.Proj)
	
	directionalShader.SetUniformFloat("dl.intensity", dl.intensity)
	directionalShader.SetUniformVec3("dl.direction", dl.dir)
	
	gl.UseProgram(directionalShader.program)
	gl.BindVertexArray(e.vao)
	gl.DrawElements(gl.TRIANGLES, e.size, gl.UNSIGNED_INT, gl.Ptr(nil))
	
	gl.DepthFunc(gl.LESS)
	gl.DepthMask(true)
	gl.Disable(gl.BLEND)
	*/
}
