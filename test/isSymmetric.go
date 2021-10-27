package test
/*
给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3

但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3


思路一
递归思路，判断两颗树是否为镜像树的步骤
1.判断两棵树t1,t2的根节点值是否一致
2.判断t1的左子树和t2的右子数是否为镜像树
3.判断t1的右子数和t2的左子树是否为镜像树
*/


type TreeNode struct {
	Val int
    Left *TreeNode
    Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
        return true
    }
	return check(root.Left, root.Right)
}

func check(leftNode *TreeNode, rightNode *TreeNode) bool {

	if leftNode == nil && rightNode == nil {
		return true
	}

	if (leftNode== nil && rightNode!=nil) || (leftNode != nil && rightNode== nil) {
		return false
	}

	if (leftNode.Val != rightNode.Val) {
		return false
	}

	if check(leftNode.Left, rightNode.Right) && check(leftNode.Right, rightNode.Left) {
		return true
	}
	return false
}