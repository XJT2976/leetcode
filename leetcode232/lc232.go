package leetcode232

type MyQueue struct {
	stackA []int
	stackB []int
}

func Constructor() MyQueue {
	return MyQueue{
		stackA: []int{},
		stackB: []int{},
	}
}

func (q *MyQueue) Push(x int) {
	for len(q.stackA) > 0 {
		q.stackB = append(q.stackB, q.stackA[len(q.stackA)-1])
		q.stackA = q.stackA[:len(q.stackA)-1]
	}
	q.stackA = append(q.stackA, x)

	for len(q.stackB) > 0 {
		q.stackA = append(q.stackA, q.stackB[len(q.stackB)-1])
		q.stackB = q.stackB[:len(q.stackB)-1]
	}
}

func (q *MyQueue) Pop() int {
	if len(q.stackA) == 0 {
		return -1
	}

	ret := q.stackA[len(q.stackA)-1]
	q.stackA = q.stackA[:len(q.stackA)-1]
	return ret
}

func (q *MyQueue) Peek() int {
	if len(q.stackA) == 0 {
		return -1
	}

	return q.stackA[len(q.stackA)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.stackA) == 0
}
