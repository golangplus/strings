package stringsp

// SliceRemove removes the index-th element from the slice and returns the result slice.
func SliceRemove(s []string, index int) []string {
	copy(s[index:], s[index+1:])
	// Remove the reference
	s[len(s)-1] = ""
	return s[:len(s)-1]
}

// SliceInsert inserts some elements at index-th position and resturns the result slice.
func SliceInsert(s []string, index int, e ...string) []string {
	n := len(e)
	if n == 0 {
		// Nothing to insert
		return s
	}

	l := len(s)
	if index == l {
		return append(s, e...)
	}

	// Make room for inserted elements.
	if cap(s) >= l+n {
		s = s[:l+n]
		copy(s[index+n:], s[index:])
		copy(s[index:], e)
		return s
	}

	if n < len(s)-index {
		// ...**** + ooo ==> ...ooo$$$$
		s = append(s, s[l-n:]...)
		// ...*...$$$
		copy(s[index+n:l], s[index:])
		// ......$$$$
		copy(s[index:], e)
		// ...ooo$$$$
		return s
	}

	// ...*** + oooo ==> ...oooo$$$
	s = append(append(s, e[l-index:]...), s[index:]...)
	// ......o$$$
	copy(s[index:], e[:l-index])
	// ...oooo$$$
	return s
}
