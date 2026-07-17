package engine

import "process-engine/internal/memory"

// hold the state of the interpreter, including memory & program counter
type Interpreter struct {
	Memory *memory.Memory
	PC int
}
