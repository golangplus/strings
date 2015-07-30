package stringsp

import (
	"testing"

	"github.com/golangplus/testing/assert"
)

func TestSliceRemove(t *testing.T) {
	assert.Equal(t, "SliceRemove([1 2 3], 1)", SliceRemove([]string{"1", "2", "3"}, 1), []string{"1", "3"})
}

func TestSliceInsert(t *testing.T) {
	assert.Equal(t, "SliceInsert([1 2 3], 3, 5)", SliceInsert([]string{"1", "2", "3"}, 3, "5"), []string{"1", "2", "3", "5"})
	assert.Equal(t, "SliceInsert([1 2 3], 1, 5)", SliceInsert([]string{"1", "2", "3"}, 1, "5"), []string{"1", "5", "2", "3"})
	assert.Equal(t, "SliceInsert([1 2 3], 1, 5, 6)", SliceInsert([]string{"1", "2", "3"}, 1, "5", "6"), []string{"1", "5", "6", "2", "3"})
	assert.Equal(t, "SliceInsert([1 2], 1, 5, 6)", SliceInsert([]string{"1", "2"}, 1, "5", "6"), []string{"1", "5", "6", "2"})
}
