package gfx

import (
	"github.com/go-gl/mathgl/mgl32"
)

type directionalLight struct {
	color, dir, intensity mgl32.Vec3
}

