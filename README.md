## Text

**Note**: This package is experimental and subject to change.
Use at your own discretion.

The text package offers a set of text rendering utilities for OpenGL
programs. It deals with TrueType fonts using [freetype-go][fg].

[fg]: https://code.google.com/p/freetype-go

![go-gl/text Screenshot 1](https://github.com/jteeuwen/text/blob/master/go-gl-text.png)


### TODO

* Fix ugly font borders (blend mode issue?).
* The positioning if strings isn't completely correct. The Y-axis
  seems to be off by a few pixels. This might be due to floating-point
  rounding errors, or a lack of understanding of the freetype-go glyph
  structure on my (jimt) part.
* Provide functions for string metrics (pixel width/height).
  These naturally have to take font scale and a bounding area into account.
  The bounding area cap may need to wrap the string over multiple lines
  in order to fit it. Possibly supply an extra parameter to determine how
  to handle these cases: Cut off the excess, or wrap lines.


### Dependencies

	go get code.google.com/p/freetype-go
    

### Usage

    go get github.com/go-gl/text


### License

Copyright 2012 The go-gl Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

