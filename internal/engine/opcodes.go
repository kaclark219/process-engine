package engine

type Opcode string

const (
	OpSet Opcode = "SET"
	OpRead Opcode = "READ"
	OpWrite Opcode = "WRITE"
	OpDelete Opcode = "DELETE"
)

type Instruction struct {
	Line int
	Opcode Opcode
	Args []string
}