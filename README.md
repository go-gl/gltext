## Text

**Note**: This package is experimental and subject to change.
Use at your own discretion.

The text package offers a set of text rendering utilities for OpenGL
programs. It deals with TrueType and Bitmap (raster) fonts. Text can be
rendered in predefined directions (Left-to-right, right-to-left and
top-to-bottom). This allows for correct display of text for various
languages.

The package supports the full set of unicode characters, provided the loaded
font does as well.


### Known bugs

* Determining the height of truetype glyphs is not entirely accurate.
  It is unclear at this point how to get to this information reliably.
  Specifically the parts in `LoadTruetype` at truetype.go#L54+.
  The vertical glyph bounds computed by freetype-go are not correct for
  certain fonts. Right now we manually offset the value by added `4` to
  the height. This is an unreliable hack and should be fixed.
* `freetype-go` does not expose `AdvanceHeight` for vertically rendered fonts.
  This may mean that the Advance size for top-to-bottom fonts is incorrect.


### Dependencies

	go get code.google.com/p/freetype-go
    

### Usage

    go get github.com/go-gl/text

Refer to [go-gl/examples/text][ex] for usage examples.

[ex]: https://github.com/go-gl/examples/text


### License

Copyright 2012 The go-gl Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

