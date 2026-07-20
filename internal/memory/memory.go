package memory

import ("fmt"
		"sync")

type entry struct {
	value any
	mu sync.RWMutex
}

type Memory struct {
	values map[string]*entry
	mu sync.RWMutex
}

func New() *Memory {
	m := &Memory{
		values: make(map[string]*entry),
	}

	return m
}

func (m *Memory) Get(key string) any {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if entry, exists := m.values[key]; exists {
		entry.mu.RLock()
		defer entry.mu.RUnlock()
		return entry.value
	}
	return nil
}

func (m *Memory) Set(key string, value any) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if entry, exists := m.values[key]; exists {
		entry.mu.Lock()
		entry.value = value
		entry.mu.Unlock()
		return
	}

	m.values[key] = &entry{
		value: value,
	}
}

func (m *Memory) Print() {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for key, entry := range m.values {
		entry.mu.RLock()
		fmt.Println("Key:", key, "Value:", entry.value)
		entry.mu.RUnlock()
	}
}
