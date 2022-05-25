package leetcode707

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	head   *Node
	length int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		nil,
		0,
	}
}

func (l *MyLinkedList) Get(index int) int {
	if index >= l.length {
		return -1
	}

	node := l.head
	for i := 0; i < index; i++ {
		node = node.next
	}

	return node.val
}

func (l *MyLinkedList) AddAtHead(val int) {
	node := &Node{
		val,
		l.head,
	}
	l.head = node
	l.length++
}

func (l *MyLinkedList) AddAtTail(val int) {
	newNode := &Node{
		val,
		nil,
	}
	if l.head == nil {
		l.head = newNode
		l.length++
		return
	}
	node := l.head
	for i := 0; i < l.length-1; i++ {
		node = node.next
	}

	node.next = newNode
	l.length++
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index > l.length {
		return
	} else if index == l.length {
		l.AddAtTail(val)
	} else if index == 0 {
		l.AddAtHead(val)
	} else {
		newNode := &Node{val: val}
		node := l.head
		for i := 0; i < index-1; i++ {
			node = node.next
		}
		newNode.next = node.next
		node.next = newNode
		l.length++
	}
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index >= l.length {
		return
	}
	if index == 0 {
		l.head = l.head.next
		l.length--
		return
	}
	node := l.head
	for i := 0; i < index-1; i++ {
		node = node.next
	}
	node.next = node.next.next
	l.length--
}
