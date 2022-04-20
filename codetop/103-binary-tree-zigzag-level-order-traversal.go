package codetop

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	right, left := make([]*TreeNode, 0), make([]*TreeNode, 0)
	right = append(right, root)
	for len(right) > 0 {
		line := make([]int, 0)
		for len(right) > 0 {
			lastIndex := len(right) - 1
			node := right[lastIndex]
			right = right[:lastIndex]
			if node.Left != nil {
				left = append(left, node.Left)
			}
			if node.Right != nil {
				left = append(left, node.Right)
			}
			line = append(line, node.Val)
		}
		result = append(result, line)
		line = make([]int, 0)
		for len(left) > 0 {
			lastIndex := len(left) - 1
			node := left[lastIndex]
			left = left[:lastIndex]
			if node.Right != nil {
				right = append(right, node.Right)
			}
			if node.Left != nil {
				right = append(right, node.Left)
			}
			line = append(line, node.Val)
		}
		if len(line) != 0 {
			result = append(result, line)
		}
	}
	return result
}
