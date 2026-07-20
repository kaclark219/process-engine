package agents

import "process-engine/internal/memory"

type Agent interface {
	Name() string
	Program() string
	Scan(memory *memory.Memory) []Alert
}

type BaseAgent struct {
	name    string
	program string
}

func NewAgent(name string, program string) *BaseAgent {
	return &BaseAgent{
		name:    name,
		program: program,
	}
}

func (a *BaseAgent) Name() string {
	return a.name
}

func (a *BaseAgent) Program() string {
	return a.program
}
