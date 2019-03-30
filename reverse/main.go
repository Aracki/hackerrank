package main

import "fmt"

func main() {

	head := DoublyLinkedListNode{data: 1}
	h1 := DoublyLinkedListNode{data: 45}
	h2 := DoublyLinkedListNode{data: 3}

	head.next = &h1
	h1.prev = &head
	h1.next = &h2
	h2.prev = &h1
	h2.next = nil

	srtH := reverse(&head)

	node := srtH
	for node != nil {
		fmt.Println(node.data)
		node = node.next
	}
}

type DoublyLinkedListNode struct {
	data int32
	next *DoublyLinkedListNode
	prev *DoublyLinkedListNode
}

func reverse(head *DoublyLinkedListNode) *DoublyLinkedListNode {

	p1 := head
	p2 := p1.next
	p1.next = nil
	p1.prev = p2

	for p2 != nil {
		p2.prev = p2.next
		p2.next = p1
		p1 = p2
		p2 = p2.prev
	}
	return p1
}
