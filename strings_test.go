// Copyright 2015 The Golang Plus Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stringsp

import (
	"fmt"
	"testing"

	"github.com/golangplus/fmt"
	"github.com/golangplus/testing/assert"
)

func ExampleCallbackFields() {
	CallbackFields("Hello World ", func(n int) {
		fmtp.Printfln("Totally %d lines:", n)
	}, func(f string) {
		fmtp.Printfln("  %s", f)
	})
	// OUTPUT:
	// Totally 2 lines:
	//   Hello
	//   World
}

func TestCallbackFieldsFunc(t *testing.T) {
	s := " Hello  \tWorld,Go   "
	var cnt int
	var res string

	CallbackFieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == '\t' || r == ','
	}, func(n int) {
		cnt = n
	}, func(f string) {
		res += "|" + f
	})

	assert.Equal(t, "res", res, "|Hello|World|Go")
}

func ExampleFullJoin() {
	a := []string{
		"item1", "item two",
	}
	fmt.Println(FullJoin(a, "(", "), (", ")"))

	// OUTPUT:
	// (item1), (item two)
}

func TestFullJoin(t *testing.T) {
	a := []string{
		"item1", "item two",
	}
	assert.Equal(t, "FullJoin", FullJoin(a, "(", "), (", ")"), "(item1), (item two)")
}
