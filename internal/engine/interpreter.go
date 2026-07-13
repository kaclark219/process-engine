package engine

import "fmt"
import "process-engine/internal/memory"

// hold the state of the interpreter, including memory & program counter
type Interpreter struct {
	Memory *memory.Memory
	PC int
}

// determines what function to call based on opcode, then executes the instruction
func (i *Interpreter) Execute(inst Instruction) {
	switch inst.Opcode {
	case OpSet:
		i.executeSet(inst)
		i.PC++
	case OpRead:
		i.executeRead(inst)
		i.PC++
	case OpWrite:
		i.executeWrite(inst)
		i.PC++
	case OpDelete:
		i.executeDelete(inst)
		i.PC++
	default:
		// handle unknown opcode
		message := fmt.Sprintf("Unknown opcode: %v", inst.Opcode)
		fmt.Println(message)
	}
}

// calls the appropriate function for each opcode, passing in the instruction arguments
func (i *Interpreter) executeSet(inst Instruction) {
	i.Memory.Set(inst.Args[0], inst.Args[1])
}
func (i *Interpreter) executeRead(inst Instruction) {
	i.Memory.Read(inst.Args[0])
}
func (i *Interpreter) executeWrite(inst Instruction) {
	i.Memory.Write(inst.Args[0], inst.Args[1])
}
func (i *Interpreter) executeDelete(inst Instruction) {
	i.Memory.Delete(inst.Args[0])
}