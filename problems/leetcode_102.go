package problems

/*
102. 二叉树的层序遍历
*/

/*
思路1. 广度优先-队列
*/
func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	var nodeQ = []*TreeNode{root}
	var nodes [][]int

	for len(nodeQ)>0 {
		curLen := len(nodeQ)
		var levelNodes []int
		for j := 0; j < curLen; j++ {

			levelNodes = append(levelNodes, nodeQ[j].Val)

			if nodeQ[j].Left != nil {
				nodeQ = append(nodeQ, nodeQ[j].Left)
			}

			if nodeQ[j].Right != nil {
				nodeQ = append(nodeQ, nodeQ[j].Right)
			}
		}
		nodes = append(nodes, levelNodes)
		nodeQ = nodeQ[curLen:]
	}
	return nodes
}
