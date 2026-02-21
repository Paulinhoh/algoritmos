package pilhas

type StackLinkedList struct {
	Top    *StackNode
	Length int
}

type StackNode struct {
	Data any
	Next *StackNode
}

func (sl *StackLinkedList) push(data any) {
	newNode := &StackNode{Data: data}

	newNode.Next = sl.Top
	sl.Top = newNode
	sl.Length++
}

func (sl *StackLinkedList) pop() any {
	result := sl.Top.Data
	sl.Top = sl.Top.Next
	sl.Length--

	return result
}

func (sl *StackLinkedList) peek() any {
	return sl.Top.Data
}

func (sl *StackLinkedList) isEmpity() bool {
	return sl.Length == 0
}
