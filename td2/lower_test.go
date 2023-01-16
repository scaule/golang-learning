package td2_test

import (
	"testing"

	"github.com/scaule/td2"
	"github.com/stretchr/testify/assert"
)

func TestLower(t *testing.T) {
	t.Run("It should be able to find lower element", func(t *testing.T) {
		assert.Equal(t, 1, td2.Lower([]int{1, 2, 3, 4}))
		assert.Equal(t, 2, td2.Lower([]int{3, 5, 9, 2}))
	})
}
