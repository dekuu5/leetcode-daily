package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
 }

func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]bool)
	tmp := head
	for tmp != nil {
		if m[tmp] {
			return true
		}
		m[tmp] = true
		tmp = tmp.Next
	}
	return false
}
// type Stack []*ListNode

// func (s *Stack) Push(node *ListNode) {
// 	*s = append(*s, node)
// }

// func (s *Stack) Pop() *ListNode {
// 	if len(*s) == 0 {
// 		return nil
// 	}
// 	idx := len(*s) - 1
// 	node := (*s)[idx]
// 	*s = (*s)[:idx]
// 	return node
// }
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0,Next: nil}
	curr := dummy

	c := 0

	for l1 != nil || l2 != nil {

		sum := c

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next

		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		c = sum / 10
		curr.Next = &ListNode{Val: sum %10, Next: nil}
		curr = curr.Next
		

	}
	if c > 0{
		curr.Next = &ListNode{Val: c, Next: nil}
	}
	return dummy.Next
}

func main(){
	fmt.Println("hi")
}