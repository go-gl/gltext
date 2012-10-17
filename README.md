## Text

**Note**: This package is experimental and subject to change.
Use at your own discretion.

The text package offers a set of text rendering utilities for OpenGL
programs. It deals with TrueType and Bitmap (raster) fonts.

Text can be rendered in predefined directions (Left-to-right, right-to-left and
top-to-bottom). This allows for correct display of text for various languages.

This package supports the full set of unicode characters, provided the loaded
font does as well.


### TODO

* Provide functions for string metrics (pixel width/height). This hinges on
  bug fixes to freetype-go. Specifically: [issue 5](http://code.google.com/p/freetype-go/issues/detail?id=5)


### Known bugs

* Bitmap font rendering has issues with non-power-of-two scale factors.
  There seems to be a problem with texture coordinates for each glyph in these
  cases. I have verified that the values in the `Glyph` structs, as well as the
  scaled input image to `LoadFont` are correct in all situations. As well as
  the output of `ToPow2`. The fault lies somewhere in the calculations of the
  glyph texture coordinates at around line 96 of `font.go`.


### Dependencies

	go get code.google.com/p/freetype-go
    

### Usage

    go get github.com/go-gl/text

Refer to [go-gl/examples/text][ex] for a usage example.

[ex]: https://github.com/go-gl/examples/text


### License

Copyright 2012 The go-gl Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

