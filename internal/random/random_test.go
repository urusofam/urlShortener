package random

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRandomString(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"len = 1", 1},
		{"len = 2", 2},
		{"len = 3", 3},
		{"len = 4", 4},
		{"len = 5", 5},
		{"len = 10", 10},
		{"len = 20", 20},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			str1 := NewRandomString(test.length)
			str2 := NewRandomString(test.length)

			assert.Len(t, str1, test.length)
			assert.Len(t, str2, test.length)

			assert.NotEqual(t, str1, str2)
		})
	}
}
