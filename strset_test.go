package stringsp_test

import (
	"testing"

	. "github.com/golangplus/strings"
	"github.com/golangplus/testing/assert"
)

func TestStrSet(t *testing.T) {
	var set StrSet

	set.Add("hello", "david")
	t.Logf("%v", set)
	assert.True(t, "set.Contain(hello)", set.Contain("hello"))
	assert.True(t, "set.Contain(david)", set.Contain("david"))
	assert.False(t, "set.Contain(villa)", set.Contain("villa"))

	assert.True(t, "set.Equal(david, hello)", set.Equal(NewStrSet("david", "hello")))
	assert.False(t, "set.Equal(david, hello)", set.Equal(NewStrSet("hello")))

	set.Delete("david")
	assert.False(t, "set.Contain(david)", set.Contain("david"))
	assert.True(t, "set.Equal(hello)", set.Equal(NewStrSet("hello")))

	assert.StringEqual(t, "set.Elements().Equal(hello)", set.Elements(), []string{"hello"})
}

func TestStrSet_nil(t *testing.T) {
	var s, ss StrSet
	assert.Equal(t, "nil.Contain(david)", s.Contain("david"), false)
	assert.Equal(t, "nil.Equal(nil)", s.Equal(ss), true)
	assert.StringEqual(t, "nil.Elements()", s.Elements(), []string{})
	s.Delete("david")
}
