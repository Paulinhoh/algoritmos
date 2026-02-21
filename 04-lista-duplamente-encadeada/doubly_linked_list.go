package listaduplamenteencadeada

import "fmt"

type DoublyLinkedList struct {
	Head   *ListNode
	Tail   *ListNode
	Length int
}

type ListNode struct {
	Data     any
	Previous *ListNode
	Next     *ListNode
}

func (dl *DoublyLinkedList) printList() {
	current := dl.Head

	for current != nil {
		fmt.Print(current.Data, " --> ")
		current = current.Next
	}
	fmt.Println("null")
}

func (dl *DoublyLinkedList) printListBackward() {
	current := dl.Tail

	for current != nil {
		fmt.Print(current.Data, " --> ")
		current = current.Previous
	}
	fmt.Println("null")
}

func (dl *DoublyLinkedList) insertNodeAtBeginnig(data any) {
	newNode := &ListNode{Data: data}

	if dl.Length == 0 {
		dl.Tail = newNode
	} else {
		dl.Head.Previous = newNode
	}

	newNode.Next = dl.Head
	dl.Head = newNode
	dl.Length++
}

func (dl *DoublyLinkedList) insertNodeAtEnd(data any) {
	newNode := &ListNode{Data: data}

	if dl.Length == 0 {
		dl.Head = newNode
	} else {
		dl.Tail.Next = newNode
		newNode.Previous = dl.Tail
	}

	dl.Tail = newNode
	dl.Length++
}

func (dl *DoublyLinkedList) removeNodeAtBeginning() {
	if dl.Length == 0 {
		return
	}

	temp := dl.Head
	if dl.Head == dl.Tail {
		dl.Tail = nil
	} else {
		dl.Head.Next.Previous = nil
	}

	dl.Head = dl.Head.Next
	temp.Next = nil

	dl.Length--
}

func (dl *DoublyLinkedList) removeNodeAtEnd() {
	if dl.Length == 0 {
		return
	}

	if dl.Head == dl.Tail {
		dl.Tail = nil
		dl.Head = nil
		return
	}

	dl.Tail.Previous.Next = nil
	dl.Tail.Previous, dl.Tail = nil, dl.Tail.Previous

	dl.Length--
}
