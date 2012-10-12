## Text

**Note**: This package is experimental and subject to change.
Use at your own discretion.

The text package offers a set of text rendering utilities for OpenGL
programs. It deals with TrueType fonts using [freetype-go][fg].

[fg]: https://code.google.com/p/freetype-go


### TODO

* The positioning of strings isn't completely correct. The Y-axis
  seems to be off by a few pixels. This might be due to floating-point
  rounding errors, or a lack of understanding of the freetype-go glyph
  structure on my (jimt) part.
* Provide functions for string metrics (pixel width/height). This hinges on
  bug fixes to freetype-go. Specifically: [issue 5](http://code.google.com/p/freetype-go/issues/detail?id=5)
* Instead of supplying a low/high rune bound in Charset, perhaps we should
  allow selection of individual runes. When loading a font for the entire
  UTF8 character set, we are creating monstrous amounts of textures and
  display lists (one for each glyph). I can't think of a use-case where this
  would actually ever happen, but it increases loading time to unacceptable
  lenghts. Being able to specify individual runes, or disparate subsets, we can
  reduce this as much as possible by generating only those runes which are
  actually being used.
  
  The difficulty with this approach is to map each rune to its texture and
  display list in an efficient manner, so we can still render a string with
  a single call to `glCallLists`. One solution is to use an intermediate
  storage list which maps runes to texture/list ids. This will generate a fair
  amount of extra memory usage though.
* Implement a BitmapFont type which loads font data from a sprite sheet.


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

