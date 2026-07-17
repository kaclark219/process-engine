package main

import ("fmt"
		"process-engine/internal/agents")

func main() {
	run()
}

func run() {
	ruleAgent := agents.NewRuleAgent("MaintainTankTemperature", "../rules/MaintainTankTemperature.yaml")
	fmt.Println(ruleAgent.Program())
	agents.PrintRule(*ruleAgent)
}
