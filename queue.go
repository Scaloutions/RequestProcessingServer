package main

// Queue is a basic FIFO queue based on a circular list that resizes as needed.
// Borrowed from:  https://gist.github.com/moraes/2141121
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
	nodes []TSRequest
	size  int
	head  int
	tail  int
	count int
}

// func NewQueue(size int) *RequestsQueue {
// 	return &RequestsQueue{
// 		nodes: make([]*TSRequest, size),
// 		size:  size,
// 	}
// }

// Push adds a node to the queue.
func (q *RequestsQueue) Push(n TSRequest) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]TSRequest, len(q.nodes)+q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++

}

// Pop removes and returns a node from the queue in first to last order.
func (q *RequestsQueue) Pop() TSRequest {
	if q.count == 0 {
		return TSRequest{} //nil
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}
