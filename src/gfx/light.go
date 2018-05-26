package gfx

import (
	"github.com/go-gl/mathgl/mgl32"
)

var (
	pointShader *Shader
	directionalShader *Shader
	ambientShader *Shader
	basicShader *Shader
)

// DirectionalLight can be used to represents light sources like suns, where only the direction matters.
type DirectionalLight struct {
	color, dir mgl32.Vec3
	intensity float32
}

// CreateDirectionalLight returns a pointer to the light and sets the uniforms in the shader.
func CreateDirectionalLight(dir mgl32.Vec3, i float32) *DirectionalLight {
	//directionalShader.SetUniformFloat("sun.intensity", i)
	//directionalShader.SetUniformVec3("sun.direction", dir)

	return &DirectionalLight {
		mgl32.Vec3{1.0, 1.0, 1.0},
		dir,
		i,
	}
}

// PointLight is a type of light were position matters, it will shine in all directions.
type PointLight struct {
	color, position mgl32.Vec3
	constant, linear, quadratic float32
}
/*
func CreatePointLight(s *Shader, pos mgl32.Vec3) *PointLight {
	
}
*/