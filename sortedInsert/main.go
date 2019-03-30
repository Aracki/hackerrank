package main

import "fmt"

func main() {
	head := DoublyLinkedListNode{data: 2}
	h1 := DoublyLinkedListNode{data: 3}
	h2 := DoublyLinkedListNode{data: 10}

	head.next = &h1
	h1.prev = &head
	h1.next = &h2
	h2.prev = &h1
	h2.next = nil

	srtH := sortedInsert(&head, 1)

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

func sortedInsert(head *DoublyLinkedListNode, data int32) *DoublyLinkedListNode {

	newNode := &DoublyLinkedListNode{data: data}

	curHead := head
	for curHead != nil {
		if data < curHead.data {
			newNode.next = curHead
			newNode.prev = curHead.prev
			if curHead.prev != nil {
				curHead.prev.next = newNode
			}
			curHead.prev = newNode

			if curHead != head {
				return head
			}
			return newNode
		}

		if curHead.next == nil {
			break
		}
		curHead = curHead.next
	}

	curHead.next = newNode
	newNode.prev = curHead
	newNode.next = nil
	return head
}
