package main

import "fmt"
import "process-engine/internal/memory"
import "process-engine/internal/engine"
import "process-engine/internal/runtime"
import "process-engine/internal/terminal"

func main() {
	fmt.Println("STARTING PROGRAM ...")
	fmt.Println("")
	// init memory & interpreter
	mem := memory.New()
	interpreter := engine.Interpreter{Memory: mem, PC: 0}

	// load available programs, display selection menu
	programs := runtime.DiscoverPrograms()
	terminal.ShowMenu(programs)

	// get user input for selection
	selection := terminal.PromptSelection(len(programs))
	
	// start loop for running selected program
	for true {
		selectedProgram := programs[selection-1]
		fmt.Println("\nRUNNING", selectedProgram.Name, "...")
		runtime.RunProgram(selectedProgram, &interpreter)
		fmt.Println(selectedProgram.Name, "COMPLETED.")
		fmt.Println("")

		// continuation prompt
		continue_bool := terminal.ContinuePrompt()
		if !continue_bool {
			fmt.Println("SHUTTING DOWN PROGRAM ...")

			// print final memory for debugging purposes
			fmt.Println("\nFinal memory:")
			for key, value := range mem.Values() {
				fmt.Printf("%s: %v\n", key, value)
			}

			break
		}

		// show menu again for new selection
		terminal.ShowMenu(programs)
		selection = terminal.PromptSelection(len(programs))
	}
}