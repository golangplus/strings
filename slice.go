package stringsp

// SliceRemove removes the index-th element from the slice and returns the result slice.
func SliceRemove(s []string, index int) []string {
	copy(s[index:], s[index+1:])
	// Remove the reference
	s[len(s)-1] = ""
	return s[:len(s)-1]
}
