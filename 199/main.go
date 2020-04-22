package main

type TreeNode struct {
	Val int
  Left *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, current_depth int, result *[]int) {
    if root == nil {
		return
	}
	if current_depth >= len(*result) {
		*result = append(*result, root.Val)
	}
	dfs(root.Right, current_depth + 1, result)
	dfs(root.Left, current_depth + 1, result)
}

func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)
	dfs(root, 0, &result)
	return result
}

func main()  {
	
}
