package leetcode

/*
给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。

如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。



示例 1：
    1     1
   / \   / \
  2   3 2   3

输入：p = [1,2,3], q = [1,2,3]
输出：true
示例 2：
    1  1
   /    \
  2      2


输入：p = [1,2], q = [1,null,2]
输出：false
示例 3：
    1     1
   / \   / \
  2   1 1   2

输入：p = [1,2,1], q = [1,1,2]
输出：false


提示：
两棵树上的节点数目都在范围 [0, 100] 内
-104 <= Node.val <= 104

思路一
递归，判断两个树t1,t2是否相同
1.判断根节点的值是否相同
2.判断t1的左子树和t2的左子树是否相同
3.判断t1的右子树和t2的右子树是否相同
*/


func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) {
		return true
	}
	return false
}
