package engine

import (
	"process-engine/internal/agents"
	"process-engine/internal/memory"
)

type Interpreter struct {
	Memory *memory.Memory
	Agents []agents.Agent
}

func Scan(i *Interpreter) {
	for _, agent := range i.Agents {
		agent.Scan(i.Memory)
	}
}
