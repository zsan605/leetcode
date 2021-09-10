package problems

/*
	给定一个二叉树的根节点 root ，返回它的 中序 遍历。
*/
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var nodes []int
	leftNodes := inorderTraversal(root.Left)
	if leftNodes != nil {
		nodes = append(nodes, leftNodes...)
	}
	nodes = append(nodes, root.Val)
	rightNodes := inorderTraversal(root.Right)
	if rightNodes != nil {
		nodes = append(nodes, rightNodes...)
	}
	return nodes
}

func inorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	leftNodes := inorderTraversal(root.Left)
	rightNodes := inorderTraversal(root.Right)
	if leftNodes != nil && rightNodes != nil {
		return append(append(leftNodes, root.Val), rightNodes...)
	} else if leftNodes != nil && rightNodes == nil {
		return append(leftNodes, root.Val)
	} else if leftNodes == nil && rightNodes != nil {
		return append([]int{root.Val}, rightNodes...)
	} else {
		return append([]int{root.Val})
	}

}
