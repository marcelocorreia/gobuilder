package array_utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestStringInSlice(t *testing.T) {
	arr := []string{"one", "two", "three"}
	assert.True(t, StringInSlice("two", arr))
	assert.False(t, StringInSlice("blah", arr))
}