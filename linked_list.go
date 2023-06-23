package main

import "fmt"

//node represents a node in the linked list

type node[T Ordered] struct {
	data *TreeNode[T]
	next *node[T]
}

// LinkedList represents a linked list
type LinkedList[T Ordered] struct {
	head *node[T]
	size int
}

// AddNode adds a new node to the linked list
func (ll *LinkedList[T]) AddNode(data *TreeNode[T]) {
	newNode := &node[T]{data: data, next: nil}

	if ll.head == nil {
		ll.head = newNode
		ll.size++
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
		ll.size++
	}
}

// DisplayList displays the elements of the linked list
func (ll *LinkedList[T]) DisplayList() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func (ll *LinkedList[T]) InsertAtBeginning(data *TreeNode[T]) {
	newNode := &node[T]{data: data, next: ll.head}
	ll.head = newNode
	ll.size++
}

func (ll *LinkedList[T]) InsertAtPosition(data *TreeNode[T], position int) {
	if position <= 0 {
		ll.InsertAtBeginning(data)
		ll.size++
		return
	}
	newNode := &node[T]{data: data}
	current := ll.head
	for count := 0; current.next != nil && count < position; count++ {
		if count == position-1 {
			newNode.next = current.next
			current.next = newNode
		}
		current = current.next
		ll.size++
	}
}

func (ll *LinkedList[T]) Remove(data *TreeNode[T]) {
	//newNodeToRemove := &node{data: data}
	if ll.head == nil {
		return
	}
	if ll.head.data == data {
		ll.head = ll.head.next
		ll.size--
		return
	}
	current := ll.head
	for current.next != nil {
		if current.next.data == data {
			current.next = current.next.next
			return
		}
		current = current.next
		ll.size--
	}

}
func (ll *LinkedList[T]) SearchSuccess(data *TreeNode[T]) bool {
	current := ll.head
	for current.next != nil {
		if current.next.data == data {
			return true
		}
		current = current.next
	}
	return false
}

func test() {
	ll := LinkedList[int64]{}
	//for i := 0; i <= 6; i++ {
	//	ll.AddNode(i)
	//	//ll.Remove(i)
	//}
	//ll.InsertAtBeginning(1000)
	//ll.Remove(1000)
	fmt.Println(ll.size)
	ll.DisplayList()
}
