package test

/*
112. 路径总和
给你二叉树的根节点 root 和一个表示目标和的整数 targetSum ，判断该树中是否存在 根节点到叶子节点 的路径，
这条路径上所有节点值相加等于目标和 targetSum 。

叶子节点 是指没有子节点的节点。
*/

/*
思路一
递归，步骤如下
1 判断root是否为叶节点
2 如果是，判断root节点的值是否为targetSum
	2.1 如果是，返回true
	2.2 如果false，返回false
3 如果不是，targetSum等于targetSum - root的值
4 判断root的左子树是否存在path的和为targetSum
5 判断root的右子树是否存在path的和为targetSum
*/
func hasPathSum(root *TreeNode, targetSum int) bool {

	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		if targetSum == root.Val {
			return true
		} else {
			return false
		}
	}

	var leftHas, rightHas bool
	leftHas = hasPathSum(root.Left, targetSum-root.Val)
	rightHas = hasPathSum(root.Right, targetSum-root.Val)

	if leftHas || rightHas {
		return true
	}

	return false
}
