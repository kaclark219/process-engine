package runtime

import ("process-engine/internal/engine"
		"process-engine/internal/memory"
		"process-engine/internal/loader"
		"fmt"
		"strings"
		"time")

type Runtime struct {
	Interpreter engine.Interpreter
	CurrentTick int
	CurrentTime time.Time
}

func Start() *Runtime {
	rule_agents, err := loader.LoadRuleAgents("../rules")
	if err != nil {
		panic(err)
	}

	runtime := Runtime{
		Interpreter: engine.Interpreter{
			Memory: memory.New(),
			Agents: rule_agents,
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
			alerts := agent.Scan(r.Interpreter.Memory)
			for _, alert := range alerts {
				fmt.Printf("[%s] [%s] %s\n", strings.ToUpper(alert.Severity), alert.RuleName, alert.Message)
				fmt.Printf("  Target: %s\n", alert.Target)
				fmt.Printf("  Value: %.2f\n", alert.Value)
				fmt.Printf("  Condition: %s\n", alert.Condition)
			}
		}
	}

	// r.Interpreter.Memory.Print()

	r.CurrentTime = r.CurrentTime.Add(time.Minute)
}
