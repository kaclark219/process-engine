package agents

import (
	"fmt"
	"process-engine/internal/engine"
	"process-engine/internal/loader"
	"process-engine/internal/memory"
	"strconv"
)

type TankAgent struct {
	name string
	program string
}

// create TankAgent

func NewTankAgent(name string, program string, mem *memory.Memory) *TankAgent {
	mem.Set("Tank.Level", 0)
	mem.Set("Tank.Temperature", 0)
	mem.Set("Tank.Pressure", 0)
	mem.Set("Tank.Inflow", 0)
	mem.Set("Tank.Outflow", 0)
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
			if levelInt > 90 {
				mem.Set("Tank.Alarm.HighLevel", true)
			} else if levelInt < 10 {
				mem.Set("Tank.Alarm.LowLevel", true)
			} else {
				mem.Set("Tank.Alarm.HighLevel", false)
				mem.Set("Tank.Alarm.LowLevel", false)
			}
		}

		temp := mem.Read("Tank.Temperature")
		if temp != nil {
			tempInt, err := strconv.Atoi(fmt.Sprint(temp))
			if err == nil {
				if tempInt > 80 {
					mem.Set("Tank.Alarm.HighTemperature", true)
				} else if tempInt < 20 {
					mem.Set("Tank.Alarm.LowTemperature", true)
				} else {
					mem.Set("Tank.Alarm.HighTemperature", false)
					mem.Set("Tank.Alarm.LowTemperature", false)
				}
			}
		}

		pressure := mem.Read("Tank.Pressure")
		if pressure != nil {
			pressureInt, err := strconv.Atoi(fmt.Sprint(pressure))
			if err == nil {
				if pressureInt > 100 {
					mem.Set("Tank.Alarm.HighPressure", true)
				} else if pressureInt < 30 {
					mem.Set("Tank.Alarm.LowPressure", true)
				} else {
					mem.Set("Tank.Alarm.HighPressure", false)
					mem.Set("Tank.Alarm.LowPressure", false)
				}
			}
		}
	}
}
