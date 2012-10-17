// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package text

import (
	"fmt"
	"github.com/go-gl/gl"
	"github.com/go-gl/glfw"
	"image"
	"image/png"
	"os"
	"testing"
)

var fonts [8]*Font

// save saves the given bitmap to a file.
// Used to save images at various image conversion stages, so
// we can verify they behave correctly.
func save(img image.Image, file string, argv ...interface{}) {
	file = fmt.Sprintf(file, argv...)
	fd, err := os.Create(file)

	if err == nil {
		png.Encode(fd, img)
		fd.Close()
	}
}

// printf draws the same string twice with a colour and location offset,
// to simulate a drop-shadow. It does so for each loaded font.
func printf(t *testing.T, x, y float32, fs string, argv ...interface{}) {
	for i := range fonts {
		if fonts[i] == nil {
			continue
		}

		_, h := fonts[i].GlyphBounds()
		y := y + float32(i*h)

		gl.Color4f(0.1, 0.1, 0.1, 0.7)
		err := fonts[i].Printf(x+2, y+2, fs, argv...)
		if err != nil {
			t.Fatal(err)
		}

		gl.Color4f(1, 1, 1, 1)
		err = fonts[i].Printf(x, y, fs, argv...)
		if err != nil {
			t.Fatal(err)
		}
	}
}

// initGL initializes GLFW and OpenGL.
func initGL(t *testing.T, title string) {
	err := glfw.Init()
	if err != nil {
		t.Fatal(err)
	}

	err = glfw.OpenWindow(640, 480, 8, 8, 8, 8, 0, 0, glfw.Windowed)
	if err != nil {
		glfw.Terminate()
		t.Fatal(err)
	}

	glfw.SetWindowTitle("go-gl/text: " + title + " font test")
	glfw.SetSwapInterval(1)
	glfw.SetWindowSizeCallback(onResize)
	glfw.SetKeyCallback(onKey)

	gl.Init()
	gl.Disable(gl.DEPTH_TEST)
	gl.Disable(gl.LIGHTING)
	gl.ClearColor(0.2, 0.2, 0.23, 0.0)
}

// onKey handles key events.
func onKey(key, state int) {
	if key == glfw.KeyEsc {
		glfw.CloseWindow()
	}
}

// onResize handles window resize events.
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
