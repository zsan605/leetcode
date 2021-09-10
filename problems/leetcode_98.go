package problems

import "math"

func isValidBST(root *TreeNode) bool {

	return isValidBSTHelper(root, math.MinInt32, math.MaxInt64)
}

func isValidBSTHelper(root *TreeNode, lower, upper int) bool {

	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}

	return isValidBSTHelper(root.Left, lower, root.Val) && isValidBSTHelper(root.Right, root.Val, upper)
}