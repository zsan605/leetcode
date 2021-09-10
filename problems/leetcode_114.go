package problems

func flatten(root *TreeNode) {
	if root != nil {
		flattenHelper(root)
	}

}

func flattenHelper(root *TreeNode) (*TreeNode, *TreeNode) {

	var leftBefore, leftAfter, rightBefore, rightAfter *TreeNode
	if root.Left != nil {
		leftBefore, leftAfter = flattenHelper(root.Left)
		root.Left = nil
	}

	if root.Right != nil {
		rightBefore, rightAfter = flattenHelper(root.Right)
	}

	if leftBefore != nil && rightBefore != nil {
		root.Right = leftBefore
		leftAfter.Right = rightBefore
		return root, rightAfter
	} else if leftBefore != nil && rightBefore == nil {
		root.Right = leftBefore
		return root, leftAfter
	} else if leftBefore == nil && rightBefore != nil {
		root.Right = rightBefore
		return root, rightAfter
	}

	return root, root
}
