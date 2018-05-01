package gfx

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// Entity represents a mesh with material. It also contains the position, rotation and later maybe the scale.
type Entity struct {
	vao              uint32
	size             int32
	PosX, PosY, PosZ float32
	rotX, rotY, rotZ float32
	Trans            mgl32.Mat4
	mat              *Material
}

// CreateCube returns a pointer to an Entity which is a cube.
func CreateCube(posX, posY, posZ, rotX, rotY, rotZ float32, mat *Material) *Entity {
	c := &Entity{}

	c.PosX, c.PosY, c.PosZ = posX, posY, posZ
	c.rotX, c.rotY, c.rotZ = rotX, rotY, rotZ
	c.Trans = mgl32.HomogRotate3DX(c.rotX).Mul4(mgl32.HomogRotate3DY(c.rotY)).Mul4(mgl32.HomogRotate3DZ(c.rotZ))
	c.Trans = c.Trans.Mul4(mgl32.Translate3D(posX, posY, posZ))
	c.mat = mat

	vertices := []float32{
		// positions      tex coords normals
		// Back quad.
		0.5, -0.5, -0.5, 0.0, 0.0, 0.0, 0.0, -1.0,
		-0.5, -0.5, -0.5, 1.0, 0.0, 0.0, 0.0, -1.0,
		-0.5, 0.5, -0.5, 1.0, 1.0, 0.0, 0.0, -1.0,
		0.5, 0.5, -0.5, 0.0, 1.0, 0.0, 0.0, -1.0,

		// Front quad.
		-0.5, -0.5, 0.5, 0.0, 0.0, 0.0, 0.0, 1.0,
		0.5, -0.5, 0.5, 1.0, 0.0, 0.0, 0.0, 1.0,
		0.5, 0.5, 0.5, 1.0, 1.0, 0.0, 0.0, 1.0,
		-0.5, 0.5, 0.5, 0.0, 1.0, 0.0, 0.0, 1.0,

		// Left quad.
		-0.5, -0.5, -0.5, 0.0, 0.0, -1.0, 0.0, 0.0,
		-0.5, -0.5, 0.5, 1.0, 0.0, -1.0, 0.0, 0.0,
		-0.5, 0.5, 0.5, 1.0, 1.0, -1.0, 0.0, 0.0,
		-0.5, 0.5, -0.5, 0.0, 1.0, -1.0, 0.0, 0.0,

		// Right quad.
		0.5, -0.5, 0.5, 0.0, 0.0, 1.0, 0.0, 0.0,
		0.5, -0.5, -0.5, 1.0, 0.0, 1.0, 0.0, 0.0,
		0.5, 0.5, -0.5, 1.0, 1.0, 1.0, 0.0, 0.0,
		0.5, 0.5, 0.5, 0.0, 1.0, 1.0, 0.0, 0.0,

		// Bottom quad.
		0.5, -0.5, 0.5, 1.0, 0.0, 0.0, -1.0, 0.0,
		-0.5, -0.5, 0.5, 0.0, 0.0, 0.0, -1.0, 0.0,
		-0.5, -0.5, -0.5, 0.0, 1.0, 0.0, -1.0, 0.0,
		0.5, -0.5, -0.5, 1.0, 1.0, 0.0, -1.0, 0.0,

		// Top quad.
		-0.5, 0.5, 0.5, 0.0, 0.0, 0.0, 1.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 0.0, 0.0, 1.0, 0.0,
		0.5, 0.5, -0.5, 1.0, 1.0, 0.0, 1.0, 0.0,
		-0.5, 0.5, -0.5, 0.0, 1.0, 0.0, 1.0, 0.0,
	}

	indices := []uint32{
		0, 2, 3,
		0, 1, 2,

		4, 6, 7,
		4, 5, 6,

		8, 10, 11,
		8, 9, 10,

		12, 14, 15,
		12, 13, 14,

		16, 18, 19,
		16, 17, 18,

		20, 22, 23,
		20, 21, 22,
	}
	c.size = int32(len(indices) * 4)

	gl.GenVertexArrays(1, &c.vao)
	gl.BindVertexArray(c.vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	// Pass data to the shader.
	// Positions.
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)
	// Texture coordinates.
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 8*4, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(1)
	// Normals.
	gl.VertexAttribPointer(2, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(5*4))
	gl.EnableVertexAttribArray(2)

	var ibo uint32
	// Store the indices in a buffer.
	gl.GenBuffers(1, &ibo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ibo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, int(c.size), gl.Ptr(indices), gl.STATIC_DRAW)

	return c
}

// SetRot takes in x, y, and z values in degrees and rotates the Entity.
func (e *Entity) SetRot(x, y, z float32) {
	e.rotX, e.rotY, e.rotZ = x, y, z
	e.Trans = mgl32.HomogRotate3DX(e.rotX).Mul4(mgl32.HomogRotate3DY(e.rotY)).Mul4(mgl32.HomogRotate3DZ(e.rotZ))
	e.Trans = e.Trans.Mul4(mgl32.Translate3D(e.PosX, e.PosY, e.PosZ))
}