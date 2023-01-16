package td5_test

import (
	"testing"

	"github.com/scaule/td5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLinkedlist(t *testing.T) {
	t.Run("It should be able to create a linkedlist", func(t *testing.T) {
		SUT := td5.CreateLinkedList("1")
		assert.Equal(t, "1", SUT.Print())
	})

	t.Run("It should be able to insert an element to a linkedlist", func(t *testing.T) {
		SUT := td5.CreateLinkedList("1")
		assert.True(t, SUT.Add("2"))
		assert.Equal(t, "1 2", SUT.Print())
	})

	t.Run("It should be able to delete an element from linkedlist", func(t *testing.T) {
		SUT := td2.CreateLinkedList("1")
		require.True(t, SUT.Add("2"))
		assert.True(t, SUT.Remove(2))
		assert.Equal(t, "1", SUT.Print())
	})
}
