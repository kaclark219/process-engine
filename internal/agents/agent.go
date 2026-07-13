package agents

import "process-engine/internal/memory"
import "process-engine/internal/engine"


type Agent interface {
	Name() string
	Program() string
	Scan (*memory.Memory, *engine.Interpreter)
}