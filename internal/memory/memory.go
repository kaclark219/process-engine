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
	m := &Memory{
		values: make(map[string]*entry),
	}

	// test values
	m.values["process.Tank01.Temperature"] = &entry{
		value: map[string]any{
			"value":     45.0,
			"timestamp": "2026-07-17T10:15:00Z",
		},
	}
	m.values["process.Pump03.Flow"] = &entry{
		value: map[string]any{
			"value":     120.0,
			"timestamp": "2026-07-17T10:15:00Z",
		},
	}
	m.values["process.Tank02.Temperature"] = &entry{
		value: map[string]any{
			"value":     75.0,
			"timestamp": "2026-07-17T10:15:00Z",
		},
	}

	return m
}
