package main

func run() {
	//q := Queue[*TreeNode[int]]{}
	//q.enqueue()
	//
	//fmt.Println(q.size)
	//q.DisplayList()
	//fmt.Println("done")
	//a := q.dequeue()

}

type Queue[T any] struct {
	LinkedList[T]
}

func (q *Queue[T]) enqueue(data *TreeNode[T]) {
	q.AddNode(data)
}

func (q *Queue[T]) dequeue() *TreeNode[T] {
	front := q.front()
	if !q.isEmpty() {
		q.Remove(q.head.data)
	}
	return front

}

func (q *Queue[T]) front() *TreeNode[T] {
	if q.isEmpty() {
		return nil
	}
	return q.head.data
}

func (q *Queue[T]) isEmpty() bool {
	return q.head == nil
}
