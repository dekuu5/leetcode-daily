package main

import "fmt"


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    if p == nil || q == nil {
        return false
    }
    if p.Val != q.Val {
        return false
    }
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
 }


func invertTree(root *TreeNode) *TreeNode {
	

	var dfs func(r *TreeNode)
	

	dfs = func(r *TreeNode) {
		if r == nil {
			return 
		}

		r.Left ,r.Right = r.Right ,r.Left

		dfs(r.Left)
		dfs(r.Right)
	}
	
	dfs(root)

	return root
}

func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}


	return max( maxDepth(root.Left)  , maxDepth(root.Right)) + 1 
}
 
func main(){
	// Test case 1: Same trees
	tree1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	tree2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Println("Test 1 - Same trees:", isSameTree(tree1, tree2))

	// Test case 2: Different structure
	tree3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	tree4 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
	fmt.Println("Test 2 - Different structure:", isSameTree(tree3, tree4))

	// Test case 3: Different values
	tree5 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}}
	tree6 := &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}
	fmt.Println("Test 3 - Different values:", isSameTree(tree5, tree6))

	// Test case 4: Both nil
	fmt.Println("Test 4 - Both nil:", isSameTree(nil, nil))

	// Test case 5: One nil, one not
	tree7 := &TreeNode{Val: 1}
	fmt.Println("Test 5 - One nil:", isSameTree(tree7, nil))

	// Test case 6: Different tree depths
	tree8 := &TreeNode{
		Val: 1,
		Left: &TreeNode{Val: 2},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
	}
	tree9 := &TreeNode{
		Val: 1,
		Left: &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println("Test 6 - Different depths:", isSameTree(tree8, tree9))

	// Test invertTree function
	fmt.Println("\nTesting invertTree function:")

	// Create a test tree
	originalTree := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{Val: 6},
			Right: &TreeNode{Val: 9},
		},
	}

	// Print tree before inversion (in-order traversal would be: 1,2,3,4,6,7,9)
	fmt.Println("Original tree structure created")

	// Invert the tree
	invertedTree := invertTree(originalTree)

	// The inverted tree should have structure where left and right are swapped
	fmt.Println("Tree inverted successfully")
	// Print the inverted tree using DFS
	fmt.Println("Inverted tree (DFS traversal):")
	var printDFS func(*TreeNode, string)
	printDFS = func(node *TreeNode, prefix string) {
		if node == nil {
			return
		}
		fmt.Printf("%s%d\n", prefix, node.Val)
		if node.Left != nil || node.Right != nil {
			printDFS(node.Left, prefix+"  L: ")
			printDFS(node.Right, prefix+"  R: ")
		}
	}
	printDFS(invertedTree, "")
	// Test maxDepth function
	fmt.Println("\nTesting maxDepth function:")

	// Test with the inverted tree
	depth := maxDepth(invertedTree)
	fmt.Printf("Max depth of inverted tree: %d\n", depth)

	// Test with nil tree
	fmt.Printf("Max depth of nil tree: %d\n", maxDepth(nil))

	// Test with single node
	singleNode := &TreeNode{Val: 1}
	fmt.Printf("Max depth of single node: %d\n", maxDepth(singleNode))

	// Test with unbalanced tree
	unbalancedTree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{Val: 4},
			},
		},
	}
	fmt.Printf("Max depth of unbalanced tree: %d\n", maxDepth(unbalancedTree))
}