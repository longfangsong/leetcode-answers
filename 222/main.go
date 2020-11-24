func countDepth(root *TreeNode) int {
	depth := 0
    for node := root; node != nil; node = node.Left {
        depth++
	}
	return depth
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := countDepth(root.Left)
	rightDepth := countDepth(root.Right)
	if leftDepth == rightDepth {
		return countNodes(root.Right) + (1<<leftDepth)
	} else {
		return countNodes(root.Left) + (1<<rightDepth)
	}
}