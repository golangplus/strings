// Copyright 2015 The Golang Plus Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package stringsp is a plus to the standard "strings" package.
*/
package stringsp

import (
	"unicode"
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
