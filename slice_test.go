package stringsp

import (
	"testing"

	"github.com/golangplus/testing/assert"
)

func TestSliceRemove(t *testing.T) {
	assert.Equal(t, "SliceRemove([1 2 3], 1)", SliceRemove([]string{"1", "2", "3"}, 1), []string{"1", "3"})
}
