package power_map

import "github.com/yimiaoxiehou/go-tools/queue"

type FiFoMap struct {
	data  map[string]interface{}
	queue queue.Queue
}

func New(size int) *FiFoMap {
	return &FiFoMap{
		data:  map[string]interface{}{},
		queue: queue.New(size),
	}
}

func (m *FiFoMap) Get(key string) interface{} {
	if v, ok := m.data[key]; ok {
		return v
	}
	return nil
}

func (m *FiFoMap) Has(key string) bool {
	_, ok := m.data[key]
	return ok
}

func (m *FiFoMap) Set(key string, value interface{}) {
	if m.queue.IsFull() {
		removeKey := m.queue.Pop().(string)
		delete(m.data, removeKey)
	}
	m.data[key] = value
}
