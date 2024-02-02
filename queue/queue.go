package queue

type Queue struct {
	items []interface{}
	limit int
}

func New(size int) Queue {
	return Queue{
		items: make([]interface{}, 0),
		limit: size,
	}
}

func (q *Queue) Push(v interface{}) {
	// queue full, remove first.
	if len(q.items) == q.limit {
		q.items = append(q.items[1:], v)
	}
	q.items = append(q.items, v)
}

func (q *Queue) Pop() interface{} {
	head := q.items[0]
	q.items = q.items[1:]
	return head
}

func (q *Queue) IsImpty() bool {
	return len(q.items) == 0
}

func (q *Queue) IsFull() bool {
	return len(q.items) == q.limit
}

func (q *Queue) Has(v interface{}) bool {
	for _, e := range q.items {
		if e == v {
			return true
		}
	}
	return false
}
