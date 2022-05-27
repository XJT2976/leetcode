package leetcode347

import "container/heap"

type minHeap struct {
	data      []int
	dataPrior map[int]int
}

func newMinHeap(prior map[int]int) *minHeap {
	return &minHeap{
		make([]int, 0),
		prior,
	}
}

func (h *minHeap) Len() int {
	return len(h.data)
}

func (h *minHeap) Less(i, j int) bool {
	return h.dataPrior[h.data[i]] < h.dataPrior[h.data[j]]
}

func (h *minHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *minHeap) Push(x any) {
	h.data = append(h.data, x.(int))
}

func (h *minHeap) Pop() any {
	if len(h.data) == 0 {
		return -1
	}
	ret := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	return ret
}

func TopKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] += 1
	}
	h := newMinHeap(m)
	heap.Init(h)
	i := 0
	for n := range m {
		heap.Push(h, n)
		if i >= k {
			heap.Pop(h)
		}
		i++
	}

	ret := make([]int, h.Len())
	for i, d := range h.data {
		ret[i] = d
	}
	return ret
}
