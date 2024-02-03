package power_map

import "github.com/yimiaoxiehou/go-tools/queue"

type FiFoMap struct {
	data      map[string]interface{}
	queue     queue.Queue
	clearFunc func(v interface{})
}

// NewFiFoMap creates a new FiFoMap with the specified size and clear function.
// It initializes the data map, queue, and clear function.
func NewFiFoMap(size int, clearFunc func(v interface{})) *FiFoMap {
	return &FiFoMap{
		data:      map[string]interface{}{},
		queue:     queue.New(size),
		clearFunc: clearFunc,
	}
}

// Get retrieves the value associated with the given key from the FiFoMap.
func (m *FiFoMap) Get(key string) interface{} {
	// Check if the key exists in the map, if yes, return the value associated with it.
	if v, ok := m.data[key]; ok {
		return v
	}
	// Return nil if the key does not exist in the map.
	return nil
}

// Has checks if the key exists in the FiFoMap and returns true if it does, false otherwise.
func (m *FiFoMap) Has(key string) bool {
	_, ok := m.data[key]
	return ok
}

// Set adds or updates a key-value pair in the FiFoMap.
func (m *FiFoMap) Set(key string, value interface{}) {
	// If the queue is full, remove the oldest key-value pair
	if m.queue.IsFull() {
		removeKey := m.queue.Pop().(string)
		// Clear the value using the clearFunc
		if v, ok := m.data[removeKey]; ok {
			m.clearFunc(v)
		}
		delete(m.data, removeKey)
	}
	// Add or update the key-value pair
	m.data[key] = value
}
