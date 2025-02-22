/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func visitLayer(layer []*TreeNode) ([]int, []*TreeNode) {
    var intVal []int
    var nodeVal []*TreeNode
    for _, item := range layer {
        intVal = append(intVal, item.Val)
        if item.Left != nil {
            nodeVal = append(nodeVal, item.Left)
        }
        if item.Right != nil {
            nodeVal = append(nodeVal, item.Right)
        }
    }
    return intVal, nodeVal
}
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    currentLayer := []*TreeNode{root}
    var result [][]int
    for len(currentLayer) != 0 {
        subResult, nextLayer := visitLayer(currentLayer)
        result = append(result, subResult)
        currentLayer = nextLayer
    }
    return result
}
