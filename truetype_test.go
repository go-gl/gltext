// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package text

import (
	"github.com/andrebq/gas"
	"github.com/go-gl/gl"
	"github.com/go-gl/glfw"
	"os"
	"testing"
)

func TestTruetype(t *testing.T) {
	fontFile, err := gas.Abs("code.google.com/p/freetype-go/luxi-fonts/luxisr.ttf")
	if err != nil {
		t.Fatal(err)
	}

	initGL(t, "truetype")
	defer glfw.Terminate()

	// Load the same truetype font at different scale factors.
	for i := range fonts {
		err := loadTTF(i, fontFile)
		if err != nil {
			t.Fatal(err)
		}
		defer fonts[i].Release()
	}

	for glfw.WindowParam(glfw.Opened) > 0 {
		gl.Clear(gl.COLOR_BUFFER_BIT)

		mx, my := glfw.MousePos()
		printf(t, 10, 10, "0 1 2 3 4 5 6 7 8 9 A B C D E F")
		printf(t, float32(mx), float32(my), "%d x %d", mx, my)

		glfw.SwapBuffers()
	}
}

func loadTTF(i int, f string) error {
	fd, err := os.Open(f)
	if err != nil {
		return err
	}

	defer fd.Close()

	fonts[i], err = LoadTruetype(fd, int32(10+i), 32, 127, LeftToRight)
	return nil
}
