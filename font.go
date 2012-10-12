// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package text

import "io"

// A Font implementation allows rendering of text to an OpenGL context.
type Font interface {
	// LoadFile loads font data from the given file.
	LoadFile(string) error

	// LoadStream loads font data from the given stream.
	LoadStream(io.Reader) error

	// LoadBytes loads font data from the given byte data.
	LoadBytes([]byte) error

	// Release releases font resources.
	Release()

	// Charset returns the character set used to create the font.
	Charset() *Charset

	// Printf draws the given string at the specified coordinates.
	Printf(float32, float32, string, ...interface{})
}
