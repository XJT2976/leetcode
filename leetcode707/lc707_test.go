package leetcode707

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	myLinkedList := Constructor()
	myLinkedList.AddAtHead(1)
	myLinkedList.AddAtTail(3)
	myLinkedList.AddAtIndex(1, 2)

	assert.Equal(t, 2, myLinkedList.Get(1))

	myLinkedList.DeleteAtIndex(1)
	assert.Equal(t, 3, myLinkedList.Get(1))
}
