package main

import (
	"container/heap"
	"fmt"
)

type Sum interface {
	Add() error
}

type TestSum struct {
}

func (t *TestSum) Add() error {
	return nil
}

func main() {
	fmt.Println(mergeKLists([]*ListNode{
		&ListNode{
			1,
			&ListNode{
				4,
				&ListNode{
					5,
					nil,
				},
			},
		},
		&ListNode{
			1,
			&ListNode{
				3,
				&ListNode{
					4,
					nil,
				},
			},
		},
		&ListNode{
			2,
			&ListNode{
				6,
				nil,
			},
		},
	}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type ListNodeHeap []*ListNode

func (h *ListNodeHeap) Len() int { return len(*h) }

func (h *ListNodeHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *ListNodeHeap) Less(i, j int) bool { return (*h)[i].Val < (*h)[j].Val }

func (h *ListNodeHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }

func (h *ListNodeHeap) Pop() interface{} {
	ret := (*ListNode)(nil)

	if len(*h) > 0 {
		ret = (*h)[len(*h)-1]
		*h = (*h)[:len(*h)-1]
	}

	return ret
}

func mergeKLists(lists []*ListNode) *ListNode {
	var prev, head *ListNode
	h := &ListNodeHeap{}
	heap.Init(h)

	for i := range lists {
		if lists[i] != nil {
			heap.Push(h, lists[i])
		}
	}

	for h.Len() > 0 {
		n := heap.Pop(h).(*ListNode)
		if head == nil {
			head = n
		}
		if n.Next != nil {
			heap.Push(h, n.Next)
		}

		if prev != nil {
			prev.Next = n
		}

		prev = n
	}

	return head
}
