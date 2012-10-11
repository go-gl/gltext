// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package text

import (
	"github.com/andrebq/gas"
	"github.com/go-gl/gl"
	"github.com/go-gl/glfw"
	"testing"
	"time"
)

func TestFont(t *testing.T) {
	fontFile, err := gas.Abs("github.com/jteeuwen/text/testdata/goudy_bookletter_1911.ttf")
	if err != nil {
		return
	}

	initGL(t)

	font, err := LoadFontFile(fontFile, 22, Ascii)
	if err != nil {
		t.Fatal(err)
	}

	defer font.Release()

	for glfw.WindowParam(glfw.Opened) > 0 {
		mx, my := glfw.MousePos()

		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.MatrixMode(gl.MODELVIEW)
		gl.LoadIdentity()

		gl.Color4f(0.5, 0.6, 0.7, 1)
		rect(float32(mx)-16, float32(my)-16, 32, 32)

		gl.Color4f(1, 1, 1, 1)
		font.Printf(10, 0, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
		font.Printf(10, 40, "abcdefghijklmnopqrstuvwxyz 0123456789")
		font.Printf(10, 70, "Cake & Pie")
		font.Printf(10, 100, "%s", time.Now().Format(time.RFC1123))
		font.Printf(10, 130, "%d x %d", mx, my)

		glfw.SwapBuffers()
	}
}

func initGL(t *testing.T) {
	err := glfw.Init()
	if err != nil {
		t.Fatal(err)
	}

	err = glfw.OpenWindow(640, 480, 8, 8, 8, 8, 0, 0, glfw.Windowed)
	if err != nil {
		glfw.Terminate()
		t.Fatal(err)
	}

	glfw.SetSwapInterval(1)
	glfw.SetWindowSizeCallback(onResize)
	glfw.SetWindowCloseCallback(onClose)

	gl.Init()
	gl.Disable(gl.TEXTURE_2D)
	gl.Disable(gl.DEPTH_TEST)
	gl.Disable(gl.LIGHTING)
	gl.ClearColor(0.2, 0.2, 0.23, 0.0)
}

func onClose() int {
	glfw.Terminate()
	return 1
}

func onResize(w, h int) {
	if w < 1 {
		w = 1
	}

	if h < 1 {
		h = 1
	}

	gl.Viewport(0, 0, w, h)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, float64(w), float64(h), 0, 0, 1)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
}

// rect draws a quad of the given dimensions.
func rect(x, y, w, h float32) {
	x2 := x + w
	y2 := y + h

	gl.Begin(gl.QUADS)
	gl.Vertex2f(x, y)
	gl.Vertex2f(x2, y)
	gl.Vertex2f(x2, y2)
	gl.Vertex2f(x, y2)
	gl.End()
}
