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

func TestCallbackFieldsFunc_SepEnds(t *testing.T) {
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

func TestCallbackFieldsFunc_FieldEnds(t *testing.T) {
	s := " Hello  \tWorld,Go"
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
	assert.Equal(t, "FullJoin", FullJoin(nil, "(", "), (", ")"), "")
	assert.Equal(t, "FullJoin", FullJoin([]string{"item"}, "(", "), (", ")"), "(item)")
	assert.Equal(t, "FullJoin", FullJoin([]string{"item1", "item two"}, "(", "), (", ")"), "(item1), (item two)")
}

func TestCompare(t *testing.T) {
	assert.Equal(t, "Compare(abc, def)", Compare("abc", "def"), -1)
	assert.Equal(t, "Compare(def, ab)", Compare("def", "ab"), 1)
	assert.Equal(t, "Compare(ab, ab)", Compare("ab", "ab"), 0)
}

func TestLessFunc(t *testing.T) {
	l := []string{"a", "b", "c"}
	less := LessFunc(l)
	assert.True(t, "less", less(0, 1))
	assert.False(t, "less", less(1, 1))
	assert.False(t, "less", less(2, 1))
}

func TestMatchPrefix(t *testing.T) {
	res, ok := MatchPrefix("abc", "a")
	assert.True(t, "ok", ok)
	assert.Equal(t, "res", res, "bc")

	res, ok = MatchPrefix("abc", "d")
	assert.False(t, "ok", ok)
	assert.Equal(t, "res", res, "abc")
}

func TestGet(t *testing.T) {
	assert.Equal(t, "Get(nil)", Get(nil), "")
	s := ""
	assert.Equal(t, "Get(&s)", Get(&s), "")
	s = "abc"
	assert.Equal(t, "Get(&s)", Get(&s), "abc")
}

func TestIndentLinesExceptFirst(t *testing.T) {
	assert.Equal(t, "IndentLinesExceptFirst", IndentLinesExceptFirst("", " "), "")
	assert.Equal(t, "IndentLinesExceptFirst", IndentLinesExceptFirst("line", " "), "line")
	assert.Equal(t, "IndentLinesExceptFirst", IndentLinesExceptFirst("one\ntwo", " "), "one\n two")
}
