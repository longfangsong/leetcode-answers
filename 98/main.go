func check(root *TreeNode, min *int, max *int) bool {
    if root == nil {
        return true
    }
    if min != nil && *min >= root.Val {
        return false
    }
    if max != nil && *max <= root.Val {
        return false
    }
    if root.Left != nil {
        if !check(root.Left, min, &root.Val) {
            return false
        }
    }
    if root.Right != nil {
        if !check(root.Right, &root.Val, max) {
            return false
        }
    }
    return true
}

func isValidBST(root *TreeNode) bool {
    return check(root, nil, nil)
}
