package agents

import "process-engine/internal/memory"


type Agent interface {
	Name() string
	Program() string
	Scan (*memory.Memory)
}