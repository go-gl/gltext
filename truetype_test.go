// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package text

import (
	"bytes"
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
		if i%2 == 0 {
			err = loadTTF1(i, fontFile)
		} else {
			err = loadTTF2(i)
		}

		if err != nil {
			t.Fatal(err)
		}

		defer fonts[i].Release()
	}

	for glfw.WindowParam(glfw.Opened) > 0 {
		gl.Clear(gl.COLOR_BUFFER_BIT)

		w, h := glfw.WindowSize()
		x := float32(w / 2)
		y := float32(h / 4)

		printf(t, x, y, "0 1 2 3 4 5 6 7 8 9 A B C D E F")

		glfw.SwapBuffers()
	}
}

func loadTTF1(i int, f string) error {
	fd, err := os.Open(f)
	if err != nil {
		return err
	}

	defer fd.Close()

	fonts[i], err = LoadTruetype(fd, int32(12+i), 32, 127, RightToLeft)
	return err
}

func loadTTF2(i int) (err error) {
	buf := bytes.NewBuffer(goudy_bookletter_1911_ttf())
	fonts[i], err = LoadTruetype(buf, int32(12+i), 32, 127, LeftToRight)
	return
}
