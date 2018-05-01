package gfx

import "github.com/go-gl/gl/v3.3-core/gl"

// Renderer is an OpenGL renderer.
type Renderer struct {
}

// BeginFrame clears the screen, do this before rendering.
func (r *Renderer) BeginFrame() {
	// The background color.
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.Clear(gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(0.2, 0.3, 0.3, 1.0)
}

// Render takes in an Entity and draws it to the framebuffer.
func (r *Renderer) Render(e *Entity) {
	// Set the texture and specular lighting map.
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, e.mat.texID)
	gl.ActiveTexture(gl.TEXTURE1)
	gl.BindTexture(gl.TEXTURE_2D, e.mat.specID)

	gl.BindVertexArray(e.vao)
	gl.DrawElements(gl.TRIANGLES, e.size, gl.UNSIGNED_INT, gl.Ptr(nil))
}
