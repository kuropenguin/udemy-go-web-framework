package framework

import (
	"net/http"
	"strings"
)

type TreeNode struct {
	children []*TreeNode
	handler  func(w http.ResponseWriter, r *http.Request)
	param    string
}

func Constructor() TreeNode {
	return TreeNode{
		children: make([]*TreeNode, 0),
		param:    "",
	}
}

func (this *TreeNode) Insert(pathname string, handler func(w http.ResponseWriter, r *http.Request)) {
	node := this
	params := strings.Split(pathname, "/")
	for _, param := range params {
		child := node.findChild(param)

		if child == nil {
			child = &TreeNode{
				param:    param,
				children: []*TreeNode{},
			}
			node.children = append(node.children, child)
		}
		node = child
	}
	node.handler = handler
}

func (this *TreeNode) findChild(param string) *TreeNode {
	for _, child := range this.children {
		if child.param == param {
			return child
		}
	}
	return nil
}

func (this *TreeNode) Search(pathname string) func(w http.ResponseWriter, r *http.Request) {
	node := this
	params := strings.Split(pathname, "/")
	for _, param := range params {
		child := node.findChild(param)
		if child == nil {
			return nil
		}
		node = child
	}
	return node.handler
}
