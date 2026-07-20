package main

import (
	"process-engine/internal/agents"
	"process-engine/internal/engine"
	"process-engine/internal/memory"
)

func main() {
	run()
}

func run() {
	rule_agent := agents.NewRuleAgent("MaintainTankTemperature", "../rules/MaintainTankTemperature.yaml")
	agents.PrintRule(*rule_agent)

	interpreter := &engine.Interpreter{
		Memory: memory.New(),
		Agents: []agents.Agent{rule_agent},
	}

	engine.Scan(interpreter)
}
