package problems

/*
104. 二叉树的最大深度
*/

/*
思路1: 深度优先-栈，递归
max(leftMaxDepth, rightMaxDepth) + 1
时间复杂度：O(n),对每一个节点都进行遍历
空间复杂度：O(height),递归height层，每层O(1)
*/
func maxDepthDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepthDFS(root.Left)
	rightDepth := maxDepthDFS(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

/*
思路2: 广度优先-队列
按层遍历，记录遍历的层数
时间复杂度：O(n)，遍历所有节点
空间复杂度：O(n)
*/
func maxDepthBFS(root *TreeNode) int {

	if root == nil {
		return 0
	}

	var nodeQ []*TreeNode
	var height = 0
	nodeQ = append(nodeQ, root)
	for len(nodeQ) > 0 {
		lNum := len(nodeQ)
		for i := 0; i < lNum; i++ {
			if nodeQ[i].Left != nil {
				nodeQ = append(nodeQ, nodeQ[i].Left)
			}

			if nodeQ[i].Right != nil {
				nodeQ = append(nodeQ, nodeQ[i].Right)
			}
		}
		height++
		nodeQ = nodeQ[lNum:]
	}
	return height
}
