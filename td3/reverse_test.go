package td3_test

import (
	"testing"

	"github.com/scaule/td3"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	t.Run("It should be able to reverse a slice", func(t *testing.T) {
		assert.Equal(t, []int{4, 3, 2, 1}, td3.Reverse([]int{1, 2, 3, 4}))
		assert.Equal(t, []int{2, 9, 5, 3}, td3.Reverse([]int{3, 5, 9, 2}))
	})
}
