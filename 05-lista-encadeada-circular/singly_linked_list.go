package listaencadeadacircular

import "fmt"

type CircularLinkedList struct {
	Length int
	Last   *ListNode
}

type ListNode struct {
	Data any
	Next *ListNode
}

func (cl *CircularLinkedList) printList() {
	if cl.Length == 0 {
		return
	}

	first := cl.Last.Next
	for first != cl.Last {
		fmt.Print(first.Data, " --> ")
		first = first.Next
	}
	fmt.Printf("%v --> null\n", first.Data)
}

func (cl *CircularLinkedList) insertNodeAtBeginnig(data any) {
	newNode := &ListNode{Data: data}

	if cl.Length == 0 || cl.Last == nil {
		cl.Last = newNode
	} else {
		newNode.Next = cl.Last.Next
	}

	cl.Last.Next = newNode
	cl.Length++
}

func (cl *CircularLinkedList) insertNodeAtEnd(data any) {
	newNode := &ListNode{Data: data}

	if cl.Length == 0 || cl.Last == nil {
		cl.Last = newNode
		cl.Last.Next = cl.Last
	} else {
		newNode.Next, cl.Last.Next = cl.Last.Next, newNode
		cl.Last = newNode
	}

	cl.Length++
}

func (cl *CircularLinkedList) removeNodeAtEnd() {
	if cl.Last == nil || cl.Length == 0 {
		return
	}

	if cl.Last == cl.Last.Next {
		cl.Last = nil
	} else {
		first := cl.Last.Next
		for first.Next != cl.Last {
			first = first.Next
		}

		first.Next, cl.Last = cl.Last.Next, first
	}

	cl.Length--
}
