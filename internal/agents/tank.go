package agents

import (
	"fmt"
	"process-engine/internal/engine"
	"process-engine/internal/loader"
	"process-engine/internal/memory"
	"strconv"
)

type TankAgent struct {
	name    string
	program string
}

// create TankAgent
func NewTankAgent(name string, program string) *TankAgent {
	return &TankAgent{name: name, program: program}
}

func (t *TankAgent) Name() string {
	return t.name
}

func (t *TankAgent) Program() string {
	return t.program
}

func (t *TankAgent) Scan(mem *memory.Memory, interpreter *engine.Interpreter) {
	instructions, err := loader.LoadProgram(t.program)
	if err != nil {
		fmt.Println("Error loading program:", err)
		return
	}

	for _, inst := range instructions {
		interpreter.Execute(inst)

		level := mem.Read("Tank.Level")
		if level != nil {
			levelInt, err := strconv.Atoi(fmt.Sprint(level))
			if err != nil {
				continue
			}

			if levelInt > 100 {
				mem.Set("Tank.Alarm", true)
			} else {
				mem.Set("Tank.Alarm", false)
			}
			if levelInt > 0 {
				mem.Set("Tank.Status", "ON")
			} else {
				mem.Set("Tank.Status", "OFF")
			}
		}
	}
}
