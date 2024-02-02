package queue

type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsImpty() bool {
	return len((*q)) == 0
}

func (q *Queue) has(v interface{}) bool {
	for _, e := range *q {
		if e == v {
			return true
		}
	}
	return false
}
