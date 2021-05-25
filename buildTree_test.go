package leetcode

import "testing"

func TestBuildTree(t *testing.T){
	preorder := []int{3,9,20,15,7}
	inorder:= []int{9,3,15,20,7}
	buildTree(preorder, inorder)
}

func TestSplitSubTree(t *testing.T){
	a, b := splitSubTree(3, []int{1,3,4,5,6})
	t.Log(a)
	t.Log(b)
}