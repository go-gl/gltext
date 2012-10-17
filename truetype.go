// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package text

import (
	"code.google.com/p/freetype-go/freetype"
	"code.google.com/p/freetype-go/freetype/truetype"
	"fmt"
	"image"
	"io"
	"io/ioutil"
)

// LoadTruetype loads a truetype font from the given stream and 
// applies the given font scale in points.
//
// The low and high values determine the lower and upper rune limits
// we should load for this font. For standard ASCII this would be: 32, 127.
//
// The dir value determines the orientation of the text we render
// with this font. This should be any of the predefined Direction constants.
func LoadTruetype(r io.Reader, scale int32, low, high rune, dir Direction) (*Font, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	// Read the truetype font.
	ttf, err := truetype.Parse(data)
	if err != nil {
		return nil, err
	}

	// Create our FontConfig type.
	var fc FontConfig
	fc.Dir = dir
	fc.Low = low
	fc.High = high
	fc.Glyphs = make(Charset, high-low+1)

	// Create an image, large enough to store all requested glyphs.
	//
	// We limit the image to 16 glyphs per row. Then add as many rows as
	// needed to encompass all glyphs, while making sure the resulting image
	// has power-of-two dimensions.
	gc := int32(len(fc.Glyphs))
	gpr := int32(16)
	gpc := (gc / gpr) + 1

	if gpc == 0 {
		gpc = 1
	}

	gb := ttf.Bounds(scale)
	gw := gb.XMax - gb.XMin
	gh := gb.YMax - gb.YMin
	iw := pow2(uint32(gw * gpr))
	ih := pow2(uint32(gh * gpc))

	rect := image.Rect(0, 0, int(iw), int(ih))
	img := image.NewRGBA(rect)

	// Use a freetype context to do the drawing.
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(ttf)
	c.SetFontSize(float64(scale))
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.White)

	// Iterate over all relevant glyphs in the truetype font and
	// draw them all to the image buffer.
	//
	// For each glyph, we also create a corresponding Glyph structure
	// for our Charset. It contains the appropriate glyph coordinate offsets.
	var gi int
	var px, py int32
	buf := truetype.NewGlyphBuf()

	for ch := low; ch <= high; ch++ {
		index := ttf.Index(ch)
		err := buf.Load(ttf, scale, index, nil)

		if err != nil {
			return nil, fmt.Errorf("Failed to load glyph data for rune %c: %v", ch, err)
		}

		metric := ttf.HMetric(scale, index)
		fc.Glyphs[gi].Advance = int(metric.AdvanceWidth)
		fc.Glyphs[gi].X = int(px)
		fc.Glyphs[gi].Y = int(py)
		fc.Glyphs[gi].Width = int(gw)
		fc.Glyphs[gi].Height = int(gh)

		fmt.Printf("%d: %dx%d %+v\n", gi, iw, ih, fc.Glyphs[gi])

		pt := freetype.Pt(int(px), int(py)+int(c.PointToFix32(float64(scale))>>8))
		c.DrawString(string(ch), pt)

		if gi%16 == 0 {
			px = 0
			py += gh
		} else {
			px += gw
		}

		gi++
	}

	return loadFont(img, &fc)
}

/*
// makeList makes a display list for the given glyph.
func (f *TruetypeFont) makeList(ttf *truetype.Font, gb *truetype.GlyphBuf, r rune) (err error) {
	glyph := ttf.Index(r)

	err = gb.Load(ttf, f.scale, glyph, nil)
	if err != nil {
		return
	}

	// Glyph dimensions.
	metric := ttf.HMetric(f.scale, glyph)
	glyphWidth := float32(metric.AdvanceWidth)
	glyphHeight := float32(f.scale)

	// Create power-of-two texture dimensions.
	texWidth := pow2(int(glyphWidth))
	texHeight := pow2(int(glyphHeight))

	// Create a temporary image to render to.
	rect := image.Rect(0, 0, texWidth, texHeight)
	img := image.NewGray16(rect)

	// Use a freetype context to do the drawing.
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(ttf)
	c.SetFontSize(float64(f.scale))
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.White)

	// Draw the glyph.
	pt := freetype.Pt(0, int(glyphHeight)+int(gb.B.YMin))
	c.DrawString(string(r), pt)

	// Index for our display list and texture. This is the same as the rune
	// value, minus the character set's lower bound.
	tex := r - f.charset.Low

	// Initialize glyph texture and render the image to it.
	f.textures[tex].Bind(gl.TEXTURE_2D)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, texWidth, texHeight,
		0, gl.LUMINANCE_ALPHA, gl.UNSIGNED_BYTE, img.Pix)

	// Build the display list which renders the texture to an
	// adequately positioned and scaled quad.
	gl.NewList(f.listbase+uint(tex), gl.COMPILE)
	f.textures[tex].Bind(gl.TEXTURE_2D)

	gl.Translatef(float32(gb.B.XMin), 0, 0)
	gl.PushMatrix()
	//gl.Translatef(0, float32(gb.B.YMin), 0)

	x := float64(glyphWidth) / float64(texWidth)
	y := float64(glyphHeight) / float64(texHeight)

	// Draw the quad.
	gl.Begin(gl.QUADS)
	gl.TexCoord2d(0, 0)
	gl.Vertex2f(0, glyphHeight)
	gl.TexCoord2d(0, y)
	gl.Vertex2f(0, 0)
	gl.TexCoord2d(x, y)
	gl.Vertex2f(glyphWidth, 0)
	gl.TexCoord2d(x, 0)
	gl.Vertex2f(glyphWidth, glyphHeight)
	gl.End()

	gl.PopMatrix()

	// Advance the current transformation to the next glyph location.
	gl.Translatef(float32(metric.AdvanceWidth), 0, 0)

	gl.EndList()
	return
}

*/
