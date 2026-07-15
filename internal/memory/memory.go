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
	// create new memory
	mem := &Memory{ values: make(map[string]*entry) }

	// initialize with conditions for valve (for now)
	fmt.Println("\nMEMORY INITIALIZED.")
	mem.Set("Valve.Alarm", false)
	mem.Set("Valve.Status", "OFF")
	mem.Set("Valve.Vlv", nil)
	fmt.Println("")

	return mem
}

func (m *Memory) Values() map[string]any {
    m.mu.RLock()
	entries := make(map[string]*entry, len(m.values))
	for k, e := range m.values {
		entries[k] = e
	}
	m.mu.RUnlock()

	result := make(map[string]any, len(entries))
	for k, e := range entries {
		e.mu.RLock()
		result[k] = e.value
		e.mu.RUnlock()
	}

    return result
}

// all opcode functionality below
func (m *Memory) Set(key string, value any) {
	m.mu.Lock()
	m.values[key] = &entry{ value: value }
	m.mu.Unlock()

	message := fmt.Sprintf("SET %s: %v", key, value)
	fmt.Println(message)
}

func (m *Memory) Read(key string) any {
	// lock entire memory to find if key exists
	m.mu.RLock()
	e, exists := m.values[key]
	m.mu.RUnlock()

	// check if key exists in memory, if not print error message
	if !exists {
		message := fmt.Sprintf("!ERROR: READ %s - key does not exist in memory", key)
		fmt.Println(message)
		return nil
	}

	// lock only entry for specific action
	e.mu.RLock()
	value := e.value
	e.mu.RUnlock()

	message := fmt.Sprintf("READ %s: %v", key, value)
	fmt.Println(message)
	return value
}

func (m *Memory) Write(key string, value any) {
	// lock entire memory to find if key exists
	m.mu.RLock()
	e, exists := m.values[key]
	m.mu.RUnlock()

	// check if key exists in memory, if not print error message
	if !exists {
		message := fmt.Sprintf("!ERROR: WRITE %s - key does not exist in memory", key)
		fmt.Println(message)
		return
	}

	e.mu.Lock()
	original_value := e.value
	e.value = value
	e.mu.Unlock()

	message := fmt.Sprintf("WRITE %s: %v -> %v", key, original_value, value)
	fmt.Println(message)
}

func (m *Memory) Delete(key string) {
	// have to keep locked the whole time to ensure key isn't modified by another thread during delete process
    m.mu.Lock()
    e, exists := m.values[key]

    if !exists {
        m.mu.Unlock()
        fmt.Printf("!ERROR: DELETE %s - key does not exist in memory\n", key)
        return
    }

    e.mu.Lock()
    delete(m.values, key)
    e.mu.Unlock()
    m.mu.Unlock()

    fmt.Printf("DELETE %s\n", key)
}