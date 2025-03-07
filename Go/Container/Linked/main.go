package main

import "fmt"

////////////////////////////////////////////////////////////////////////////////
// 포인터로 연결한 노드들로 구성
// 불연성 메모리
// Graph의 일종
// Single Linked List와 Double Linked List
// Cache miss가 발생할 확률이 높음
// 요소가 사라질 때 GC가 일어남
////////////////////////////////////////////////////////////////////////////////

type SingleNode[T any] struct {
	next *SingleNode[T] // Signle Linked List
	// prev *Node[T] // Double Linked List
	val T
}

func Apped[T any](node *SingleNode[T], next *SingleNode[T]) *SingleNode[T] {
	node.next = next
	return next
}

func signleLinked() {
	root := &SingleNode[int]{nil, 10}
	// root.next = &Node[int]{nil, 20}
	// root.next.next = &Node[int]{nil, 30}

	// for n := root; n != nil; n = n.next {
	// 	fmt.Printf("signleLinked node val: %d\n", n.val)
	// }

	node := root

	for i := 0; i < 3; i++ {
		node = Apped(node, &SingleNode[int]{nil, (i + 2) * 10})
	}

	// for n := root; n != nil; n = n.next {
	// 	fmt.Println("signleLinked Val:", n.val)
	// }

	node = root.next
	originNext := node.next
	node = Apped(node, &SingleNode[int]{nil, 100})
	node.next = originNext

	for n := root; n != nil; n = n.next {
		fmt.Println("signleLinked Val:", n.val)
	}
}

type DoubleNode[T any] struct {
	next *DoubleNode[T] // Signle Linked List
	prev *DoubleNode[T] // Double Linked List
	val  T
}

func doubleLinked() {
	root := &DoubleNode[int]{nil, nil, 10}
	root.next = &DoubleNode[int]{nil, root, 20}
	root.next.next = &DoubleNode[int]{nil, root.next, 30}

	tail := root.next.next

	for n := tail; n != nil; n = n.prev {
		fmt.Printf("doubleLinked node val: %d\n", n.val)
	}
}

func main() {
	signleLinked()
	doubleLinked()
}
