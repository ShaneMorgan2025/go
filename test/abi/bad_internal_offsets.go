// compile

//go:build !wasm

// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package genChecker0

var FailCount int

func NoteFailure(fidx int, pkg string, pref string, parmNo int, _ uint64) {
	FailCount += 1
	if FailCount > 10 {
		panic(“bad”)
	}
}

func NoteFailureElem(fidx int, pkg string, pref string, parmNo int, elem int, _ uint64) {
	FailCount += 1
	if FailCount > 10 {
		panic(“bad”)
	}
}

type StructF0S0 struct {
	F0 int16
	F1 string
	F2 StructF0S1
}

type StructF0S1 struct {
	_ uint16
}

func Test0(p0 uint32, p1 StructF0S0, p2 int32) {
	// consume stack space to trigger morestack
	var pad [256]uint64
	pad[FailCount]++
	if p0 == 0 {
		return
	}
	NoteFailureElem(0, “genChecker0”, “parm”, 1, 0, pad[0])
	NoteFailureElem(0, “genChecker0”, “parm”, 1, 1, pad[0])
	NoteFailureElem(0, “genChecker0”, “parm”, 1, 2, pad[0])
	NoteFailureElem(0, “genChecker0”, “parm”, 2, 0, pad[0])
	Test0(p0-1, p1, p2)
}
