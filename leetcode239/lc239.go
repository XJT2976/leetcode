package leetcode239

type MonotonicQueue struct {
	queue []int
}

func NewMonotonicQueue() *MonotonicQueue {
	return &MonotonicQueue{
		make([]int, 0),
	}
}

func (q *MonotonicQueue) Push(x int) {
	for !q.Empty() && q.queue[len(q.queue)-1] < x {
		q.queue = q.queue[:len(q.queue)-1]
	}
	q.queue = append(q.queue, x)
}

func (q *MonotonicQueue) Pop(x int) {
	if !q.Empty() && q.Front() == x {
		q.queue = q.queue[1:]
	}
}

func (q *MonotonicQueue) Empty() bool {
	return len(q.queue) == 0
}

func (q *MonotonicQueue) Front() int {
	return q.queue[0]
}

func MaxSlidingWindow(nums []int, k int) []int {
	q := NewMonotonicQueue()
	ret := make([]int, 0)
	for i := 0; i < k; i++ {
		q.Push(nums[i])
	}
	ret = append(ret, q.Front())
	for i := k; i < len(nums); i++ {
		q.Pop(nums[i-k])
		q.Push(nums[i])

		ret = append(ret, q.Front())
	}

	return ret
}
