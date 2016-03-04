// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gltext

import (
	"testing"
)

func TestPow2(t *testing.T) {
	tests := [...][2]uint32{
		{0, 0}, {1, 1}, {2, 2}, {3, 4},
		{4, 4}, {5, 8}, {6, 8}, {7, 8},
		{8, 8}, {9, 16}, {10, 16},
	}

	for i := range tests {
		ret := Pow2(tests[i][0])
		if ret != tests[i][1] {
			t.Fatalf("Pow2(%d): Want %d, Have %d",
				tests[i][0], tests[i][1], ret)
		}
	}
}

func TestIsPow2(t *testing.T) {
	tests := [...]struct {
		In  uint32
		Out bool
	}{
		{1, true}, {2, true}, {3, false}, {4, true}, {5, false},
		{6, false}, {7, false}, {8, true}, {9, false}, {10, false},
	}

	for i := range tests {
		ret := IsPow2(tests[i].In)
		if ret != tests[i].Out {
			t.Fatalf("isPow2(%d): Want %d, Have %d",
				tests[i].In, tests[i].Out, ret)
		}
	}
}
