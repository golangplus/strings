// Copyright 2015 The Golang Plus Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package stringsp is a plus to the standard "strings" package.
*/
package stringsp

import (
	"unicode"

	"github.com/golangplus/bytes"
)

// CallbackFields calls CallbackFieldsFunc with unicode.IsSpace.
func CallbackFields(s string, prepare func(n int), newField func(f string)) {
	CallbackFieldsFunc(s, unicode.IsSpace, prepare, newField)
}

// CallbackFieldsFunc is similar to strings.FieldsFunc but results are returned by two callback functions.
// prepare is firstly called with the number of total fields. Then each field is given by calling newField
// in order. If no fields are found, prepare is called with a zero value.
//
// This function is especially useful when we want to split a string into a slice of named types other than
// string or in a non-slice data structure at all.
func CallbackFieldsFunc(s string, isSpace func(rune) bool, prepare func(n int), newField func(f string)) {
	// First count the fields.
	n := 0
	inField := false
	for _, r := range s {
		wasInField := inField
		inField = !isSpace(r)
		if inField && !wasInField {
			n++
		}
	}

	// Now create them.
	prepare(n)
	fieldStart := -1 // Set to -1 when looking for start of field.
	for i, r := range s {
		if isSpace(r) {
			if fieldStart >= 0 {
				newField(s[fieldStart:i])
				fieldStart = -1
			}
		} else if fieldStart == -1 {
			fieldStart = i
		}
	}
	if fieldStart >= 0 { // Last field might end at EOF.
		newField(s[fieldStart:])
	}
}

// FullJoin is similar to strings.Join with additional prefix/suffix if a is non-empty.
func FullJoin(a []string, prefix, sep, suffix string) string {
	if len(a) == 0 {
		return ""
	}

	if len(a) == 1 {
		return prefix + a[0] + suffix
	}

	n := len(prefix) + len(sep) + len(suffix)
	for _, s := range a {
		n += len(s)
	}

	buf := make(bytesp.ByteSlice, 0, n)
	buf.WriteString(prefix)
	buf.WriteString(a[0])
	for _, s := range a[1:] {
		buf.WriteString(sep)
		buf.WriteString(s)
	}
	buf.WriteString(suffix)

	return string(buf)
}

// Compares returns -1 if a < b, 1 if a > b, 0 otherwise.
func Compare(a, b string) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	}
	return 0
}
