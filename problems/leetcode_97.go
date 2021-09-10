package problems

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return genChildTree(1, n)
}
func genChildTree(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	var ret []*TreeNode
	for i := start; i <= end; i++ {
		leftTree := genChildTree(start, i-1)
		rightTree := genChildTree(i+1, end)

		for _, lt := range leftTree {
			for _, rt := range rightTree {
				curNode := new(TreeNode)
				curNode.Val = i
				curNode.Left = lt
				curNode.Right = rt
				ret = append(ret, curNode)
			}
		}
	}
	return ret
}
