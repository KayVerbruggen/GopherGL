package gfx

import (
	"fmt"
	"image"
	"image/draw"
	"os"

	"github.com/disintegration/imaging"
	"github.com/go-gl/gl/v3.3-core/gl"
)

// Material can be attached to an Entity.
type Material struct {
	texID, specID     uint32
	Shininess float32
}

// createTex reads and sets the texture to whatever is passed.
func createTex(texFile string) (uint32, error) {
	// Open the texture file.
	tex, err := os.Open(texFile)
	check(err)
	defer tex.Close()

	texImage, _, err := image.Decode(tex)
	check(err)
	texImage = imaging.FlipV(texImage) // We need to flip it because OpenGL has 0, 0 in the bottom left.

	rgba := image.NewRGBA(texImage.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return 0, fmt.Errorf("stride from %v is unsupported", texFile)
	}
	draw.Draw(rgba, rgba.Bounds(), texImage, image.Point{0, 0}, draw.Src)

	// Generate and bind the buffer.
	var texture uint32
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)

	// Parameters for the texture.
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	// Pass the data to OpenGL and generate mipmaps.
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(rgba.Rect.Size().X), int32(rgba.Rect.Size().Y),
		0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))
	gl.GenerateMipmap(gl.TEXTURE_2D)

	return texture, nil
}

// CreateMaterial takes in an albedo and specular texture. And you can also set the shininess of the specular part.
func CreateMaterial(fileTex, fileSpec string, shininess float32) *Material {
	texID, err := createTex(fileTex)
	check(err)
	specID, err := createTex(fileSpec)
	check(err)
	return &Material{texID, specID, shininess}
}
