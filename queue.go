package main

type TSRequest struct {
	UserId        string
	PriceDollars  float64
	PriceCents    float64
	Command       string
	CommandNumber int
	Stock         string
	RequestType   string
}

type RequestsQueue struct {
	size int32
	head *QueueElement
	tail *QueueElement
}

type QueueElement struct {
	value interface{}
	next  *QueueElement
}

func (q *RequestsQueue) Size() int32 {
	return q.size
}

func (q *RequestsQueue) Head() *QueueElement {
	return q.head
}

func (q *RequestsQueue) Enqueue(element interface{}) {
	newElement := QueueElement{
		value: element,
		next:  nil,
	}

	if q.size == 0 {
		q.head = &newElement
		q.tail = &newElement
	} else {
		q.tail.next = &newElement
		q.tail = &newElement
	}

	q.size++
}

func (q *RequestsQueue) Dequeue() {
	if q.size > 0 {
		if q.size == 1 {
			q.head = nil
		} else {
			q.head = q.head.next
		}
		q.size--
	}
}

func (q *RequestsQueue) printQueue() {
	var current = q.head
	for current != nil {
		current = current.next
	}
}
