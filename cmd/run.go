package main

import (
	"fmt"
	"process-engine/internal/engine"
	"process-engine/internal/memory"
	"process-engine/internal/runtime"
	"sync"
)

func main() {
	run()
}

func run() {
	fmt.Println("STARTING PROGRAM ...")
	fmt.Println("")
	
	// init shared memory
	mem := memory.New()

	// set agents
	agents := runtime.InitializeAgents(mem)

	// run scans concurrently and wait for all agents to finish
	var wg sync.WaitGroup
	wg.Add(len(agents))
	for _, agent := range agents {
		agent := agent
		go func() {
			defer wg.Done()
			interpreter := engine.Interpreter{Memory: mem, PC: 0}
			agent.Scan(mem, &interpreter)
		}()
	}
	wg.Wait()

	fmt.Println("\nFinal memory:")
	for key, value := range mem.Values() {
		fmt.Printf("%s: %v\n", key, value)
	}
}
