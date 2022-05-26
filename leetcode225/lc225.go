package leetcode225

type MyStack struct {
	queueA []int
	queueB []int
}

func Constructor() MyStack {
	return MyStack{
		queueA: []int{},
		queueB: []int{},
	}
}

func (s *MyStack) Push(x int) {
	var queueE, queueF *[]int
	if len(s.queueA) == 0 {
		queueE = &s.queueA
		queueF = &s.queueB
	} else {
		queueE = &s.queueB
		queueF = &s.queueA
	}
	*queueE = append(*queueE, x)
	for i := 0; i < len(*queueF); i++ {
		*queueE = append(*queueE, (*queueF)[i])
	}
	*queueF = (*queueF)[:0]
}

func (s *MyStack) Pop() int {
	var queueF *[]int
	if len(s.queueA) == 0 {
		queueF = &s.queueB
	} else {
		queueF = &s.queueA
	}
	if len(*queueF) == 0 {
		return -1
	}

	ret := (*queueF)[0]
	*queueF = (*queueF)[1:]

	return ret
}

func (s *MyStack) Top() int {
	var queueF *[]int
	if len(s.queueA) == 0 {
		queueF = &s.queueB
	} else {
		queueF = &s.queueA
	}
	if len(*queueF) == 0 {
		return -1
	}

	return (*queueF)[0]
}

func (s *MyStack) Empty() bool {
	var queueF *[]int
	if len(s.queueA) == 0 {
		queueF = &s.queueB
	} else {
		queueF = &s.queueA
	}
	return len(*queueF) == 0
}
