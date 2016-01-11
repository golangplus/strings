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
	if cap(s) >= l+n {
		// s has enough room, simply make some copies
		s = s[:l+n]
		copy(s[index+n:], s[index:])
		copy(s[index:], e)
		return s
	}
	// S has not enough room, use append smartly to expand it when necessary.
	if n < l-index {
		// In this branch: ...**** + ooo ==> ...ooo$$$$
		s = append(s, s[l-n:]...)
		// Now: ...*...$$$
		copy(s[index+n:l], s[index:])
		// Now: ......$$$$
		copy(s[index:], e)
		// Now: ...ooo$$$$
		return s
	}
	// In this branch: ...*** + oooo ==> ...oooo$$$
	s = append(append(s, e[l-index:]...), s[index:]...)
	// Now: ......o$$$
	copy(s[index:], e[:l-index])
	// Now: ...oooo$$$
	return s
}
