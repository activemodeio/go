// errorcheck -0 -m -l

// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package p

type t [10000]*int

func (t) f() {
}

func x() {
	x := t{}.f // ERROR "t literal.f escapes to heap"
	x()
}

func y() {
	var i int       // ERROR "moved to heap: i"
	y := (&t{&i}).f // ERROR "\(&t literal\).f escapes to heap" "&t literal escapes to heap"
	y()
}

func z() {
	var i int    // ERROR "moved to heap: i"
	z := t{&i}.f // ERROR "t literal.f escapes to heap"
	z()
}
