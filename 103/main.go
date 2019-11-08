package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	LeftFirst  = iota
	RightFirst = iota
)

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	direction := LeftFirst
	currentQueue := []*TreeNode{root}
	var nextQueue []*TreeNode
	var result [][]int
	for len(currentQueue) != 0 {
		result = append(result, []int{})
		for i := len(currentQueue) - 1; i >= 0; i-- {
			result[len(result)-1] = append(result[len(result)-1], currentQueue[i].Val)
			if direction == LeftFirst {
				if currentQueue[i].Left != nil {
					nextQueue = append(nextQueue, currentQueue[i].Left)
				}
				if currentQueue[i].Right != nil {
					nextQueue = append(nextQueue, currentQueue[i].Right)
				}
			} else {
				if currentQueue[i].Right != nil {
					nextQueue = append(nextQueue, currentQueue[i].Right)
				}
				if currentQueue[i].Left != nil {
					nextQueue = append(nextQueue, currentQueue[i].Left)
				}
			}
		}
		if direction == LeftFirst {
			direction = RightFirst
		} else {
			direction = LeftFirst
		}
		currentQueue = nextQueue
		nextQueue = nil
	}
	return result
}