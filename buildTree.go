package leetcode

/*
思路一
递归
1.取先序数组中的第一个元素e，构建根节点，值为e
2.在中序数组中查找a,并以e为中心，拆分为两个数组a1,a2，前者a1为左子树的中序遍历数组，后者a2为右子数的中序遍历数组
3.获取左子树的前序遍历数组b1，获取右子树的前序遍历数组b2
4.使用a1,b1构建左子树，使用a2,b2构建右子树
*/
func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 || len(inorder) == 0 || len(inorder) != len(preorder) {
		return nil
	}

	var root = new(TreeNode)
	root.Val = preorder[0]
	leftArr, rightArr := splitSubTree(root.Val, inorder)

	leftArrLen, rightArrLen := len(leftArr), len(rightArr)

	if leftArrLen == 0 {
		root.Left = nil
	} else {
		leftPreorder :=preorder[1:leftArrLen+1]
		root.Left = buildTree(leftPreorder, leftArr)
	}

	if rightArrLen == 0 {
		root.Right = nil
	} else {
		rightPreorder:= preorder[leftArrLen+1:]
		root.Right = buildTree(rightPreorder, rightArr)
	}

	return root
}

func splitSubTree(e int, inorder []int) ([]int, []int) {
	var index int
	for i, val := range inorder {
		if val == e {
			index = i
		}
	}
	return inorder[0:index], inorder[index+1:]
}
