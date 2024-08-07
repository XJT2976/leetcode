package leetcode23

import (
	"container/heap"
)

type ListNode struct {
	Val int
	Next *ListNode
}

type List []*ListNode

func (l *List) Len() int {
	return len(*l)
}

func (l *List) Less(i, j int) bool{
	if (*l)[i] == nil {
		return true
	} else if (*l)[j] == nil {
		return false
	} else {
		return (*l)[i].Val < (*l)[j].Val
	}
}

func (l *List) Swap(i, j int) {
	(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
}

func (l *List) Push(x any) {
	if i, ok := x.(*ListNode); ok {
		*l = append(*l, i)
	}
}

func (l *List) Pop() any{
	if len(*l) <= 0 {
		return nil
	}
	
	ret := (*l)[len(*l) - 1]
	*l = (*l)[:len(*l) - 1]
	return ret
}

func MergeKLists(lists []*ListNode) *ListNode {
	var head, tail *ListNode
    l := List(lists)
	heap.Init(&l)

	for ;len(l) > 0;{
		x := heap.Pop(&l)
		if x == nil {
			break
		} else if i, ok := x.(*ListNode); ok {
			if i == nil {
				continue
			}
			if head == nil {
				head = i
				tail = i
			} else {
				tail.Next = i
				tail = i
			}

			if i.Next != nil {
				heap.Push(&l, i.Next)
			}
		}
	}

	return head
}