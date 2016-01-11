package stringsp

import (
	"testing"

	"github.com/golangplus/testing/assert"
)

func TestSet(t *testing.T) {
	var set Set

	set.Add("hello", "david")
	t.Logf("%v", set)
	assert.True(t, "set.Contain(hello)", set.Contain("hello"))
	assert.True(t, "set.Contain(david)", set.Contain("david"))
	assert.False(t, "set.Contain(villa)", set.Contain("villa"))

	set.Delete("david")
	assert.False(t, "set.Contain(david)", set.Contain("david"))
	assert.True(t, "set.Equal(hello)", set.Equal(NewSet("hello")))

	assert.StringEqual(t, "set.Elements().Equal(hello)", set.Elements(), []string{"hello"})
}

func TestSet_nil(t *testing.T) {
	var s, ss Set
	assert.Equal(t, "nil.Contain(david)", s.Contain("david"), false)
	assert.Equal(t, "nil.Equal(nil)", s.Equal(ss), true)
	assert.StringEqual(t, "nil.Elements()", s.Elements(), []string{})
	s.Delete("david")
}

func TestSet_Equal(t *testing.T) {
	var set Set
	set.Add("hello", "david")

	assert.True(t, "set.Equal(david, hello)", set.Equal(NewSet("david", "hello")))
	assert.False(t, "set.Equal(david, hello)", set.Equal(NewSet("hello")))
	assert.False(t, "set.Equal(david, hello)", set.Equal(NewSet("sophie", "hello")))
}
