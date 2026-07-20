package engine

import ("fmt"
		"process-engine/internal/agents"
		"process-engine/internal/memory"
		"process-engine/internal/historian")

type Interpreter struct {
	Memory *memory.Memory
	Agents []agents.Agent
}

func Scan(i *Interpreter) {
	for _, agent := range i.Agents {
		agent.Scan(i.Memory)
	}
}

type ProcessData struct {
	ProcessName string
	VariableName string
	Value float64
	TimestampUTC string
}

func (i *Interpreter) LoadTick(timestamp string) error {
	data, err := historian.LoadTimestamp(timestamp)
	if err != nil {
		return err
	}

	for _, item := range data {
		key := fmt.Sprintf("process.%s.%s", item["process"], item["variable"])
		i.Memory.Set(
			key,
			map[string]any{
				"value":     item["value"],
				"timestamp": item["timestamp"],
			},
		)
	}

	return nil
}