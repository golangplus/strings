# strings
Plus to the standard "strings" package.

## Featured
```go
// CallbackFieldsFunc is similar to strings.FieldsFunc but results are returned by two callback functions.
// prepare is firstly called with the number of total fields. Then each field is given by calling newField
// in order. If no fields are found, prepare is called with a zero value.
//
// This function is especially useful when we want to split a string into a slice of named types other than
// string or in a non-slice data structure at all.
func CallbackFieldsFunc(s string, isSpace func(rune) bool, prepare func(n int), newField func(f string)) {...}
```

## LICENSE
BSD license
