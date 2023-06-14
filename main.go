package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func main() {

	bt := BinaryTree[int64]{
		rootNode: &TreeNode[int64]{
			data: 12,
			left: &TreeNode[int64]{
				data: 24,
				left: &TreeNode[int64]{data: 24}},
			right: &TreeNode[int64]{
				data:  36,
				right: &TreeNode[int64]{data: 56}},
		},
	}
	bt.TraversalPreorder(bt.rootNode)

	//fmt.Println(q.dequeue().data)
	//fmt.Println()
	//bt.BreadthFirst()

}

type Ordered interface {
	constraints.Ordered
}
type BinaryTree[T any] struct {
	rootNode *TreeNode[T]
	TreeSize int64
}

type TreeNode[T any] struct {
	data        T
	left, right *TreeNode[T]
}

func (bt *BinaryTree[T]) TraversalPreorder(node *TreeNode[T]) {
	if bt.isEmpty() {
		return
	}

	//fmt.Println(node.data) // Preorder

	if node.left != nil {
		bt.TraversalPreorder(node.left)
	}

	fmt.Println(node.data) // Inorder
	if node.right != nil {
		bt.TraversalPreorder(node.right)
	}
	//fmt.Println(node.data) // Preorder

}

func (bt *BinaryTree[T]) isEmpty() bool {
	return bt.rootNode == nil

}

func (bt *BinaryTree[T]) BreadthFirst() {
	if bt.isEmpty() {
		return
	}

	q := Queue[T]{}
	q.enqueue(bt.rootNode)
	for !q.isEmpty() {
		dequeued := q.dequeue()
		fmt.Println(dequeued.data)
		if dequeued.right != nil {
			q.enqueue(dequeued.right)
		}
		if dequeued.left != nil {
			q.enqueue(dequeued.left)
		}

	}
}
