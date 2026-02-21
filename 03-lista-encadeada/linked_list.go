package listaencadeada

import "fmt"

type LinkedList struct {
	Head *Node
}

type Node struct {
	Data any
	Next *Node
}

func (ll *LinkedList) printList() {
	current := ll.Head
	for current != nil {
		fmt.Print(current.Data, " --> ")
		current = current.Next
	}
	fmt.Println("null")
}

func (ll *LinkedList) findLength() int {
	count := 0
	current := ll.Head
	for current != nil {
		count++
		current = current.Next
	}
	return count
}

func (ll *LinkedList) insertAtBeginning(data any) {
	newNode := &Node{Data: data, Next: ll.Head}
	ll.Head = newNode
}

func (ll *LinkedList) insertAtEnd(data any) {
	newNode := &Node{Data: data, Next: nil}
	if ll.Head == nil {
		ll.Head = newNode
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (ll *LinkedList) removeFirstNode() {
	if ll.Head == nil {
		return
	}

	ll.Head = ll.Head.Next
}

func (ll *LinkedList) removeLastNode() {
	if ll.Head == nil || ll.Head.Next == nil {
		return
	}

	previous := ll.Head
	for previous.Next.Next != nil {
		previous = previous.Next
	}

	previous.Next = nil
}

func (ll *LinkedList) insertAtGivenPosition(data any, position int) {
	newNode := &Node{Data: data}
	if position == 1 {
		newNode.Next = ll.Head
		ll.Head = newNode
		return
	}

	previous := ll.Head
	count := 1
	for count < position-1 {
		previous = previous.Next
		count++
	}

	current := previous.Next
	newNode.Next = current
	previous.Next = newNode
}

func (ll *LinkedList) removeAtGivenPosition(position int) {
	if position <= 0 {
		return
	}

	if position == 1 {
		ll.Head = ll.Head.Next
		return
	}

	previous := ll.Head
	count := 1
	for count < position-1 {
		previous = previous.Next
		count++
	}

	current := previous.Next
	previous.Next = current.Next
	current = nil
}

func (ll *LinkedList) findElement(element any) int {
	current := ll.Head
	count := 1

	if current != nil {
		if current.Data == element {
			return count
		}

		count++
		current = current.Next
	}
	return -1
}

func (ll *LinkedList) reverseList() {
	current := ll.Head
	var previuos *Node
	var next *Node

	for current != nil {
		next = current.Next
		current.Next = previuos
		previuos = current
		current = next
	}

	ll.Head = previuos
}

func (ll *LinkedList) checkLoop() bool {
	fastPointer := ll.Head
	slowPointer := ll.Head

	for fastPointer != nil && fastPointer.Next != nil {
		fastPointer = fastPointer.Next.Next
		slowPointer = slowPointer.Next
		if slowPointer == fastPointer {
			return true
		}
	}
	return false
}

func (ll *LinkedList) findWhereLoopStarts() *Node {
	fastPointer := ll.Head
	slowPointer := ll.Head

	for fastPointer != nil && fastPointer.Next != nil {
		fastPointer = fastPointer.Next.Next
		slowPointer = slowPointer.Next

		if slowPointer == fastPointer {
			temp := ll.Head
			for slowPointer != temp {
				temp = temp.Next
				slowPointer = slowPointer.Next
			}
			return temp
		}
	}
	return nil
}

func (ll *LinkedList) removeLoop() {
	fastPointer := ll.Head
	slowPointer := ll.Head

	for fastPointer != nil && fastPointer.Next != nil {
		fastPointer = fastPointer.Next.Next
		slowPointer = slowPointer.Next

		if slowPointer == fastPointer {
			temp := ll.Head
			for slowPointer.Next != temp.Next {
				temp = temp.Next
				slowPointer = slowPointer.Next
			}

			slowPointer.Next = nil
		}
	}
}
