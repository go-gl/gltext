## Text

**Note**: This package is experimental and subject to change.
Use at your own discretion.

The text package offers a set of text rendering utilities for OpenGL
programs. It deals with TrueType fonts using freetype2.


### TODO

* This package currently uses some minimal CGO bindings to freetype2.
  It binds only those components needed to make the Font type work.
  At some point, this should be ported to freetype-go.
* Allow loading of fontdata from a byte slice, instead of a file. This
  allows us to embed font files in a Go program. Currently, `FT_New_Memory_Face`
  seems to fail for an unknown reason, hence why we reverted to loading from
  a file using `FT_New_Face`.


### Usage

    go get github.com/go-gl/text


### License

Unless otherwise stated, all of the work in this project is subject to a
1-clause BSD license. Its contents can be found in the enclosed LICENSE file.

