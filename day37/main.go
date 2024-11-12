package main

import "fmt"
// tree day

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


// postoder : left right node 
// inorder : left node right
// Preorder : node left right


func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return max(maxDepth(root.Left), maxDepth(root.Right))+1
}
func printTree(root *TreeNode) {
    if root == nil {
        return
    }
    fmt.Printf("%d ", root.Val)
    printTree(root.Left)
    printTree(root.Right)
}
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    root.Left, root.Right = root.Right, root.Left
    invertTree(root.Left)
    invertTree(root.Right)
    return root
}

func main() {
    root := &TreeNode{
        Val: 1,
        Left: &TreeNode{
            Val: 2,
            Left: &TreeNode{
                Val: 4,
            },
            Right: &TreeNode{
                Val: 5,
            },
        },
        Right: &TreeNode{
            Val: 3,
            Left: &TreeNode{
                Val: 6,
            },
            Right: &TreeNode{
                Val: 7,
            },
        },
    }

    // invertedRoot := invertTree(root)

    
    fmt.Println(maxDepth(root))
    // fmt.Printf("Inverted Tree Root: %+v\n", invertedRoot)
    fmt.Println("Hello, World!")
}
