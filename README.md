## Text

**Note**: This package is experimental and subject to change.
Use at your own discretion.

The text package offers a set of text rendering utilities for OpenGL
programs. It deals with TrueType fonts using [freetype-go][fg].

[fg]: https://code.google.com/p/freetype-go


### TODO

* Fix ugly font borders (blend mode issue?).
* The positioning if strings isn't completely correct. The Y-axis
  seems to be off by a few pixels.
* Provide functions for string metrics (pixel width/height).
  These naturally have to take font scale and bounding area into account.
  The bounding area cap may need to wrap thestring over multiple lines
  in order to fit it. Possibly supply an extra parameter to determine how
  to handle these cases. Just cut off the excess, or wrap.


### Dependencies

	go get code.google.com/p/freetype-go
    

### Usage

    go get github.com/go-gl/text


### License

Copyright 2012 The go-gl Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

