package main

import (
	"fmt"
	"math"
)

// 3423. Maximum Difference Between Adjacent Elements in a Circular Array
// daily
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func maxAdjacentDistance(nums []int) int {
	m := math.MinInt

	for i := 0 ; i < len(nums); i++ {
		q := (i+1)%len(nums)
		m = max(m, abs(nums[i] - nums[q]))
	}

	return m

}

// 108. Convert Sorted Array to Binary Search Tree
// top 150

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root

}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, " ")
	inorder(root.Left)
	
	inorder(root.Right)
}

func main() {
	nums := []int{0,1,2,3,4,5}
	root := sortedArrayToBST(nums)

	fmt.Print("Inorder Traversal: ")
	inorder(root)
}
