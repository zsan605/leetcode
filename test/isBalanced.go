package test

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if (root.Left == nil && root.Right != nil && (root.Right.Left != nil || root.Right.Right != nil)) ||
		(root.Right == nil && root.Left != nil && (root.Left.Left != nil || root.Left.Right != nil)) {
		return false
	}

	if isBalanced(root.Left) && isBalanced(root.Right) {
		return true
	}
	return false
}


