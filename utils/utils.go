package utils

type ListNode struct {
	Val  int
	Next *ListNode
}

func ConvertArrayToLink(nums []int) *ListNode {
	head := (*ListNode)(nil)
	prev := head
	for _, i := range nums {
		if head == nil {
			head = &ListNode{
				i,
				nil,
			}
			prev = head
		} else {
			node := &ListNode{
				i,
				nil,
			}
			prev.Next = node
			prev = node
		}
	}
	return head
}

func ConvertLinkToArray(head *ListNode) []int {
	ret := make([]int, 0)
	for head != nil {
		ret = append(ret, head.Val)
		head = head.Next
	}

	return ret
}

func EasySortInt(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreorderTraverse(root *TreeNode) []int {
	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top != nil {
			ret = append(ret, top.Val)
			stack = append(stack, top.Right)
			stack = append(stack, top.Left)
		}
	}

	return ret
}
func PostTraverse(root *TreeNode) []int {
	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top != nil {
			ret = append(ret, top.Val)
			stack = append(stack, top.Right)
			stack = append(stack, top.Left)
		}
	}

	return ret
}

func levelOrder(root *TreeNode) [][]int {
	queue := make([]*TreeNode, 0)
	ret := make([][]int, 0)
	queue = append(queue, root)
	if root == nil {
		return ret
	}

	for len(queue) != 0 {
		size := len(queue)
		tmp := make([]int, 0)

		for i := 0; i < size; i++ {
			top := queue[0]
			queue = queue[1:]
			tmp = append(tmp, top.Val)
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
		ret = append(ret, tmp)
	}

	return ret
}
