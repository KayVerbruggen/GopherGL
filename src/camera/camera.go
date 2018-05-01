package camera

import (
	"github.com/go-gl/mathgl/mgl32"
	"GopherGL/src/gfx"
)

// Camera is a FPP camera.
type Camera struct {
	Pos, Target mgl32.Vec3
	view        mgl32.Mat4
	Fov         float32
}

// CreateCamera creates a FPP camera and sets up the view matrix.
func CreateCamera(pos mgl32.Vec3, fov float32) *Camera {
	c := Camera{}

	c.Fov = fov
	c.Pos = pos
	c.Target = mgl32.Vec3{pos.X(), pos.Y(), pos.Z() - 1.0}
	c.view = mgl32.LookAt(c.Pos.X(), c.Pos.Y(), c.Pos.Z(),
		c.Target.X(), c.Target.Y(), c.Target.Z(),
		0.0, 1.0, 0.0)

	return &c
}

// UpdateCamera update the shader variables after the camera has moved.
func (c *Camera) UpdateCamera(s *gfx.Shader) {
	c.Target = mgl32.Vec3{c.Pos.X(), c.Pos.Y(), c.Pos.Z() - 1.0}
	c.view = mgl32.LookAt(c.Pos.X(), c.Pos.Y(), c.Pos.Z(),
		c.Target.X(), c.Target.Y(), c.Target.Z(),
		0.0, 1.0, 0.0)

	s.SetUniformMat4("view", c.view)
}
