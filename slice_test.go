package stringsp

import (
	"testing"

	"github.com/golangplus/testing/assert"
)

func TestSliceRemove(t *testing.T) {
	assert.Equal(t, "SliceRemove([1 2 3], 1)", SliceRemove([]string{"1", "2", "3"}, 1), []string{"1", "3"})
}

func TestSliceInsert(t *testing.T) {
	assert.Equal(t, "SliceInsert([1 2 3], 3, 5)", SliceInsert([]string{"1", "2", "3"}, 3), []string{"1", "2", "3"})
	assert.Equal(t, "SliceInsert([1 2 3], 3, 5)", SliceInsert([]string{"1", "2", "3"}, 3, "5"), []string{"1", "2", "3", "5"})
	assert.Equal(t, "SliceInsert([1 2 3], 1, 5)", SliceInsert([]string{"1", "2", "3"}, 1, "5"), []string{"1", "5", "2", "3"})
	assert.Equal(t, "SliceInsert([1 2 3], 1, 5, 6)", SliceInsert([]string{"1", "2", "3"}, 1, "5", "6"), []string{"1", "5", "6", "2", "3"})
	assert.Equal(t, "SliceInsert([1 2], 1, 5, 6)", SliceInsert([]string{"1", "2"}, 1, "5", "6"), []string{"1", "5", "6", "2"})
}

func TestSliceInsert_NotExpaned(t *testing.T) {
	src := make([]string, 3, 4)
	copy(src, []string{"1", "2", "3"})
	assert.Equal(t, "cap(src)", cap(src), 4)
	assert.Equal(t, "SliceInsert(src, 1, 5)", SliceInsert(src, 1, "5"), []string{"1", "5", "2", "3"})
}
