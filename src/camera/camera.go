// Package camera can be used to create a perspective camera, more types will be added later.
package camera

import (
	"github.com/go-gl/mathgl/mgl32"
)

// Camera is a FPP camera.
type Camera struct {
	Pos, Target mgl32.Vec3
	View, Proj  mgl32.Mat4
	Fov         float32
}

// CreateCamera creates a FPP camera and sets up the view matrix.
func CreateCamera(pos mgl32.Vec3, aspect, fov float32) *Camera {
	c := Camera{}

	c.Fov = fov
	c.Pos = pos
	c.Target = mgl32.Vec3{pos.X(), pos.Y(), pos.Z() - 1.0}
	c.View = mgl32.LookAt(c.Pos.X(), c.Pos.Y(), c.Pos.Z(),
		c.Target.X(), c.Target.Y(), c.Target.Z(),
		0.0, 1.0, 0.0)

	c.Proj = mgl32.Perspective(mgl32.DegToRad(fov), aspect, 0.1, 1000.0)

	return &c
}

// Update updates the camera's target and viewmmatrix. It also updates shader variables after the camera has moved.
func (c *Camera) Update() {
	c.Target = mgl32.Vec3{c.Pos.X(), c.Pos.Y(), c.Pos.Z() - 1.0}
	c.View = mgl32.LookAt(c.Pos.X(), c.Pos.Y(), c.Pos.Z(),
		c.Target.X(), c.Target.Y(), c.Target.Z(),
		0.0, 1.0, 0.0)
}

// SetProjection takes the new fov and aspect ratio and creates a new projection matrix.
func (c *Camera) SetProjection(aspect, fov float32) {
	c.Proj = mgl32.Perspective(mgl32.DegToRad(fov), aspect, 0.1, 1000.0)
}