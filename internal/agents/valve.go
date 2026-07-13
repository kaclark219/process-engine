package agents

import (
	"fmt"
	"process-engine/internal/engine"
	"process-engine/internal/loader"
	"process-engine/internal/memory"
	"strconv"
)

type ValveAgent struct {
	name    string
	program string
}

// create ValveAgent
func NewValveAgent(name string, program string) *ValveAgent {
	return &ValveAgent{name: name, program: program}
}

func (v *ValveAgent) Name() string {
	return v.name
}

func (v *ValveAgent) Program() string {
	return v.program
}

func (v *ValveAgent) Scan(mem *memory.Memory, interpreter *engine.Interpreter) {
	// valve logic needs to read instructions from memory to determine ON/OFF, alarm, & value
	instructions, err := loader.LoadProgram(v.program)
	if err != nil {
		fmt.Println("Error loading program:", err)
		return
	}

	for _, inst := range instructions {
		interpreter.Execute(inst)

		vlv := mem.Read("Valve.Vlv")
		if vlv != nil {
			vlvInt, err := strconv.Atoi(fmt.Sprint(vlv))
			if err != nil {
				continue
			}

			if vlvInt > 100 {
				mem.Set("Valve.Alarm", true)
			} else {
				mem.Set("Valve.Alarm", false)
			}
			if vlvInt > 0 {
				mem.Set("Valve.Status", "ON")
			} else {
				mem.Set("Valve.Status", "OFF")
			}
		}
	}
}
