// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

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
// The lowr and upper bounds are inclusive.
type Charset struct {
	Low  rune // Lower rune bound.
	High rune // Upper rune bound.
}

// Len returns the range of the character set.
func (c *Charset) Len() int { return int(c.High-c.Low) + 1 }

// Builtin charsets
var (
	Ascii = &Charset{32, 127}
	UTF8  = &Charset{0, MaxRune}
)
