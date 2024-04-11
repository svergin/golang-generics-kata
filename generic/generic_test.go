package generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	s1 := []string{"A", "B", "C", "D", "E", "F"}

	s2 := Filter(s1, func(f string) bool {
		if f == "C" || f == "D" {
			return false
		}
		return true
	})

	assert.Equal(t, 4, len(s2))
	assert.Contains(t, s2, "A")
	assert.Contains(t, s2, "B")
	assert.Contains(t, s2, "E")
	assert.Contains(t, s2, "F")
}
