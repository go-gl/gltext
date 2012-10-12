// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package text

// Maximum value a valid UTF8 rune can have.
const MaxRune = '\U0010FFFF'

// A Charset represents a range of runes which should
// be used to prepare a font for rendering.
//
// If a font is created specifically to render only a small subset of runes,
// a charset containing only these runes can ensure the font does not allocate
// more texture data than strictly necessary.
//
// The lower and upper bounds are inclusive.
type Charset struct {
	Low  rune // Lower rune bound.
	High rune // Upper rune bound.
}

// Len returns the range of the character set.
func (c *Charset) Len() int { return int(c.High-c.Low) + 1 }

// Builtin character sets.
var (
	Ascii = &Charset{32, 127}
	UTF8  = &Charset{0, MaxRune}
)
