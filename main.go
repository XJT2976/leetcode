package main

import (
	"container/heap"
	"fmt"
	"strconv"
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
	fmt.Println(mergeSort(
		&ListNode{
			4,
			&ListNode{
				2,
				&ListNode{
					1,
					&ListNode{
						3,
						nil,
					},
				},
			},
		}))
	getHint("11", "10")
}
func getHint(secret string, guess string) string {
	m := make(map[byte]int)
	A := 0
	B := 0

	for i := range secret {
		if secret[i] == guess[i] {
			A++
		} else {
			m[secret[i]] += 1
		}
	}

	for i := range guess {
		if m[guess[i]] > 0 {
			B++
			m[guess[i]] -= 1
		}
	}

	return strconv.Itoa(A) + "A" + strconv.Itoa(B) + "B"
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

func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow := head
	fast := head
	prev := head

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev.Next = nil

	l1 := mergeSort(head)
	l2 := mergeSort(slow)

	return merge(l1, l2)
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}

		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	}

	if l2 != nil {
		cur.Next = l2
	}

	return dummy.Next
}
