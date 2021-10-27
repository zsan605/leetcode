package test

/*
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。

思路一
计算树的最大深度步骤
1.计算左子树的最大深度
2.计算右子数的最大深度
3.左子树深度+1和右子树深度+1做比较，返回较大的
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var leftDepth = 1 + maxDepth(root.Left)
	var rightDepth = 1 + maxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth
	} else {
		return rightDepth
	}
}
