package main

import (
	tree "Golang-Learning/GoBasic/02.structAndInterface/002.tree"
	"fmt"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = &tree.Node{Value: 2}
	root.Right.Left.SetValue(4)

	root.Traverse()

	fmt.Println("==============")
	myRoot := myTreeNode{&root}
	myRoot.postOrder()

}
