package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func main() {

	bt := BinaryTree[int64]{rootNode: &Node[int64]{
		data:  12,
		left:  &Node[int64]{data: 24, left: &Node[int64]{data: 24}},
		right: &Node[int64]{data: 36},
	},
	}
	bt.TraversalPreorder(bt.rootNode)

}

type Ordered interface {
	constraints.Ordered
}
type BinaryTree[T Ordered] struct {
	rootNode *Node[T]
	TreeSize int64
}

type Node[T Ordered] struct {
	data        T
	left, right *Node[T]
}

func (BT *BinaryTree[T]) TraversalPreorder(node *Node[T]) {
	if BT.isEmpty() {
		return
	}

	//fmt.Println(node.data) // Preorder

	if node.left != nil {
		BT.TraversalPreorder(node.left)
	}

	fmt.Println(node.data) // Inorder
	if node.right != nil {
		BT.TraversalPreorder(node.right)
	}
	//fmt.Println(node.data) // Preorder

}

func (BT *BinaryTree[T]) isEmpty() bool {
	return BT.rootNode == nil

}

func (BT *BinaryTree[T]) BreadthFirst() {

}
