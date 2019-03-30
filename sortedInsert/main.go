package main

import "fmt"

func main() {
	f := DoublyLinkedListNode{data: 2}
	s := DoublyLinkedListNode{data: 3}
	h := DoublyLinkedListNode{data: 10}

	f.next = &s
	s.prev = &f
	s.next = &h
	h.prev = &s
	h.next = nil

	srtH := sortedInsert(&h, 7)

	node := srtH
	for node != nil {
		fmt.Println(node.data)
		node = node.prev
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
		if data > curHead.data {

			if curHead.next == nil {
				newNode.prev = curHead
				newNode.next = nil
				curHead.next = newNode
				return newNode
			}

			newNode.prev = curHead
			newNode.next = curHead.next
			curHead.next.prev = newNode
			curHead.next = newNode

			return head
		}
		if curHead.prev == nil {
			break
		}
		curHead = curHead.prev
	}

	// if the smallest data num
	newNode.next = curHead
	curHead.prev = newNode
	return head
}
