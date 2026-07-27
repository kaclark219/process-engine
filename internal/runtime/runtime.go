package runtime

import ("fmt"
		"process-engine/internal/agents"
		"process-engine/internal/engine"
		"process-engine/internal/loader"
		"process-engine/internal/memory"
		"strings"
		"sync"
		"time")

type AlertPublisher interface {
	PublishAlert(alert *agents.Alert)
}

type Runtime struct {
	mu sync.Mutex
	Interpreter engine.Interpreter
	CurrentTick int
	CurrentTime time.Time
	Publisher AlertPublisher
}

func Start(rulesDir string, publisher AlertPublisher) *Runtime {
	rule_agents, err := loader.LoadRuleAgents(rulesDir)
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
		Publisher: publisher,
	}

	return &runtime
}

func (r *Runtime) AddAgent(agent agents.Agent) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Interpreter.Agents = append(r.Interpreter.Agents, agent)
}

func (r *Runtime) Tick() {
	r.mu.Lock()
	defer r.mu.Unlock()
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

				if r.Publisher != nil {
					r.Publisher.PublishAlert(&alert)
				}
			}
		}
	}

	// r.Interpreter.Memory.Print()

	r.CurrentTime = r.CurrentTime.Add(time.Minute)
}
