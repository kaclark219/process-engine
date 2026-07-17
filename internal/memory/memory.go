package memory

import (
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
	return &Memory{ values: make(map[string]*entry) }
}
