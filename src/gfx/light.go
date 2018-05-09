package gfx

import (
	"github.com/go-gl/mathgl/mgl32"
)

// DirectionalLight can be used to represents light sources like suns, where only the direction matters.
type DirectionalLight struct {
	color, dir mgl32.Vec3
	intensity float32
}

// CreateDirectionalLight returns a pointer to the light and sets the uniforms in the shader.
func CreateDirectionalLight(s *Shader, dir mgl32.Vec3, i float32) *DirectionalLight {
	s.SetUniformFloat("sun.intensity", i)
	s.SetUniformVec3("sun.direction", dir)

	return &DirectionalLight {
		mgl32.Vec3{1.0, 1.0, 1.0},
		dir,
		i,
	}
}