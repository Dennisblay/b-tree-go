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
				left: &TreeNode[int64]{data: 16}},
			right: &TreeNode[int64]{
				data:  36,
				right: &TreeNode[int64]{data: 56}},
		},
	}
	//bt.TraversalPreorder(bt.rootNode)
	//bt.BreadthFirstTraversal()
	//fmt.Println(bt.FindItem(36))
	//fmt.Println(bt.FindItem(12))
	//bt.FindItem(56)
	bt.AddItem(64)
	bt.AddItem(61)
	bt.AddItem(61)
	bt.TraversalPreorder(bt.rootNode)
	//fmt.Println(q.dequeue().data)
	//fmt.Println()
	//bt.BreadthFirstTraversal()

}

type Ordered interface {
	constraints.Ordered
}
type BinaryTree[T Ordered] struct {
	rootNode *TreeNode[T]
	TreeSize int64
}

type TreeNode[T Ordered] struct {
	data        T
	left, right *TreeNode[T]
}

func (bt *BinaryTree[T]) TraversalPreorder(node *TreeNode[T]) {
	if bt.isEmpty() {
		return
	}

	fmt.Println(node.data) // Preorder

	if node.left != nil {
		bt.TraversalPreorder(node.left)
	}

	//fmt.Println(node.data) // Inorder
	if node.right != nil {
		bt.TraversalPreorder(node.right)
	}
	//fmt.Println(node.data) // Postorder

}

func (bt *BinaryTree[T]) isEmpty() bool {
	return bt.rootNode == nil

}

func (bt *BinaryTree[T]) BreadthFirstTraversal() {
	if bt.isEmpty() {
		return
	}

	q := Queue[T]{}
	q.enqueue(bt.rootNode)
	for !q.isEmpty() {
		dequeued := q.dequeue()
		fmt.Println(dequeued.data)
		if dequeued.left != nil {
			q.enqueue(dequeued.left)
		}
		if dequeued.right != nil {
			q.enqueue(dequeued.right)
		}

	}
}

func (bt *BinaryTree[T]) FindItem(targetData T) bool {
	// Compare the rootNode to the target node
	currentNode := bt.rootNode
	for currentNode != nil {
		if currentNode.data == targetData {
			return true
		}
		if targetData < currentNode.data {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}
	return false
}

func (bt *BinaryTree[T]) AddItem(data T) {
	dataNode := &TreeNode[T]{data: data}
	node := bt.rootNode
	if bt.isEmpty() {
		bt.rootNode = dataNode
		return
	}
	for {
		if node.data == data {
			fmt.Println("Data already exists")
			return
		}
		if data < node.data {
			if node.left == nil {
				node.left = dataNode
				return
			}
			node = node.left
		} else {

			if data > node.data {
				if node.right == nil {
					node.right = dataNode
					return
				}
				node = node.right
			}
		}
	}

}
