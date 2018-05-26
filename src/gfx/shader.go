package gfx

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// check is used to check errors.
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Used for specifying the shader type
const (
	shaderVertex   uint16 = 0
	shaderFragment uint16 = 1
)

// Shader is an OpenGL shader.
type Shader struct {
	program uint32
}

// CompileShader compiles the shader and prints any errors to the console.
func compileShader(source string, shaderType uint32, length int32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csource, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csource, &length)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("Failed to compile: %v", log)
	}

	return shader, nil
}

// CreateShader sets the shader and takes in a file path.
func createShader(shaderFile string) *Shader {
	// Read the file.
	source, err := os.Open(shaderFile)
	defer source.Close()
	check(err)
	scanner := bufio.NewScanner(source)
	scanner.Split(bufio.ScanLines)

	var shaderType uint16
	var vertexSource string
	var fragmentSource string

	// Loops through each line.
	for scanner.Scan() {
		// Check which part of the shader we're currently in.
		if scanner.Text() == "#vertex" {
			shaderType = shaderVertex
		} else if scanner.Text() == "#fragment" {
			shaderType = shaderFragment
		} else {
			// Add the source to the string according to its type.
			if shaderType == shaderVertex {
				vertexSource += scanner.Text() + "\n"
			} else if shaderType == shaderFragment {
				fragmentSource += scanner.Text() + "\n"
			}
		}
	}

	program := gl.CreateProgram()
	vertexShader, err := compileShader(vertexSource, gl.VERTEX_SHADER, int32(len(vertexSource)))
	check(err)
	fragmentShader, err := compileShader(fragmentSource, gl.FRAGMENT_SHADER, int32(len(fragmentSource)))
	check(err)

	// Link the seperate shaders into one program.
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)
	gl.ValidateProgram(program)

	// The shader are useless and can be deleted.
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	gl.UseProgram(program)
	return &Shader{program}
}

// SetUniformVec3 sets a uniform variable of type vec3.
func (s *Shader) SetUniformVec3(name string, v mgl32.Vec3) {
	// Add null terminator.
	name += "\x00"

	// Get the location, and give the pointer to the first value.
	loc := gl.GetUniformLocation(s.program, gl.Str(name))
	gl.Uniform3fv(loc, 1, &v[0])
}

// SetUniformFloat sets a uniform variable of type float32.
func (s *Shader) SetUniformFloat(name string, f float32) {
	// Add null terminator.
	name += "\x00"

	// Get the location, and give the pointer to the first value.
	loc := gl.GetUniformLocation(s.program, gl.Str(name))
	gl.Uniform1f(loc, f)
}

// SetUniformMat4 sets a uniform variable of type mat4.
func (s *Shader) SetUniformMat4(name string, m mgl32.Mat4) {
	// Add null terminator.
	name += "\x00"

	// Get the location, and give the pointer to the first value.
	loc := gl.GetUniformLocation(s.program, gl.Str(name))
	gl.UniformMatrix4fv(loc, 1, false, &m[0])
}

// SetUniformInt32 sets a uniform variable of type int32.
func (s *Shader) SetUniformInt32(name string, i int32) {
	// Add null terminator.
	name += "\x00"

	// Get the location, and give the pointer to the first value.
	loc := gl.GetUniformLocation(s.program, gl.Str(name))
	gl.Uniform1i(loc, i)
}
/*
func (s* shader) setUniformDirectionalLight(name string, dl directionalLight) {
	
}
*/