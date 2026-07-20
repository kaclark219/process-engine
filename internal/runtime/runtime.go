package runtime

import (
	"process-engine/internal/agents"
	"process-engine/internal/engine"
	"process-engine/internal/memory"
)

import "time"

type Runtime struct {
	Interpreter engine.Interpreter
	CurrentTick int
	CurrentTime time.Time
}

func Start() *Runtime {
	rule_agent := agents.NewRuleAgent("MaintainTankTemperature", "../rules/MaintainTankTemperature.yaml")

	runtime := Runtime{
		Interpreter: engine.Interpreter{
			Memory: memory.New(),
			Agents: []agents.Agent{rule_agent},
		},
		CurrentTick: 0,
		CurrentTime: time.Date(2026, 7, 17, 10, 15, 0, 0, time.UTC),
	}

	return &runtime
}

func (r *Runtime) Tick() {
	r.CurrentTick++

	err := r.Interpreter.LoadTick(r.CurrentTime.Format(time.RFC3339))
	if err != nil {
		panic(err)
	}

	if r.CurrentTick%3 == 0 {
		for _, agent := range r.Interpreter.Agents {
			agent.Scan(r.Interpreter.Memory)
		}
	}

	// r.Interpreter.Memory.Print()

	r.CurrentTime = r.CurrentTime.Add(time.Minute)
}
