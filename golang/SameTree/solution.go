package SameTree

import (
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getPreorder(node *TreeNode, path []*int) []*int {
	if node == nil {
		path = append(path, nil)
		return path
	}
	path = append(path, &node.Val)
	path = getPreorder(node.Left, path)
	path = getPreorder(node.Right, path)
	return path
}

func getInorder(node *TreeNode, path []*int) []*int {
	if node == nil {
		path = append(path, nil)
		return path
	}
	path = getPreorder(node.Left, path)
	path = append(path, &node.Val)
	path = getPreorder(node.Right, path)
	return path
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	a := getPreorder(p, []*int{})
	b := getPreorder(q, []*int{})

	if !reflect.DeepEqual(a, b) {
		return false
	}
	c := getInorder(p, []*int{})
	d := getInorder(q, []*int{})

	if !reflect.DeepEqual(c, d) {
		return false
	}
	return true
}
