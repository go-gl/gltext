// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package text

import (
	"fmt"
	"github.com/go-gl/gl"
	"github.com/go-gl/glu"
	"image"
)

// Some miscellaneous bits and bobs.

// pow2 returns the first power-of-two value >= to n.
// This is used to create suitable texture dimensions.
func pow2(x uint32) uint32 {
	x--
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	return x + 1
}

// isPow2 returns true if the given value is a power-of-two.
func isPow2(x uint32) bool { return (x & (x - 1)) == 0 }

// checkGLError returns an opengl error if one exists.
func checkGLError() error {
	errno := gl.GetError()

	if errno == gl.NO_ERROR {
		return nil
	}

	str, err := glu.ErrorString(errno)
	if err != nil {
		return fmt.Errorf("Unknown GL error: %d", errno)
	}

	return fmt.Errorf(str)
}

// toRGBA translates the given image to RGBA format if necessary.
// Optionally scales it by the given amount.
func toRGBA(src image.Image, scale int) *image.RGBA {
	if scale < 1 {
		scale = 1
	}

	dst, ok := src.(*image.RGBA)
	if ok && scale == 1 {
		return dst
	}

	// Scale image to match new size.
	ib := src.Bounds()
	rect := image.Rect(
		0, 0,
		ib.Dx()*scale,
		ib.Dy()*scale,
	)

	if !ok {
		// Image is not RGBA, so we create it.
		dst = image.NewRGBA(rect)
	}

	for sy := 0; sy < ib.Dy(); sy++ {
		for sx := 0; sx < ib.Dx(); sx++ {
			dx := sx * scale
			dy := sy * scale
			pixel := src.At(sx, sy)

			for scy := 0; scy < scale; scy++ {
				for scx := 0; scx < scale; scx++ {
					dst.Set(dx+scx, dy+scy, pixel)
				}
			}
		}
	}

	return dst
}

// toPow2 returns the given image, scaled to the smallest power-of-two
// dimensions >= the input image dimensions.
func toPow2(src *image.RGBA) *image.RGBA {
	ib := src.Bounds()

	if isPow2(uint32(ib.Dx())) && isPow2(uint32(ib.Dy())) {
		return src // Nothing to do.
	}

	rect := image.Rect(
		0, 0,
		int(pow2(uint32(ib.Dx()))),
		int(pow2(uint32(ib.Dy()))),
	)

	dst := image.NewRGBA(rect)

	for y := 0; y < ib.Dy(); y++ {
		for x := 0; x < ib.Dx(); x++ {
			dst.Set(x, y, src.At(x, y))
		}
	}

	return dst
}
