package filas

type Queue struct {
	Length int
	Front  *QueueNode
	Rear   *QueueNode
}

type QueueNode struct {
	Data any
	Next *QueueNode
}

func (q *Queue) enqueue(data any) {
	temp := &QueueNode{Data: data}

	if q.isEmpity() {
		q.Front = temp
	} else {
		q.Rear.Next = temp
	}

	q.Rear = temp
	q.Length++
}

func (q *Queue) dequeue() any {
	if q.isEmpity() {
		return nil
	}

	temp := q.Front
	q.Front = temp.Next
	temp.Next = nil

	q.Length--
	return temp.Data

}

func (q *Queue) length() int {
	return q.Length
}

func (q *Queue) isEmpity() bool {
	return q.length() == 0
}
