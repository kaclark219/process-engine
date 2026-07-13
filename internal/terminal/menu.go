package terminal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
import "process-engine/internal/runtime"

func readLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(line), nil
}


func ShowMenu(programs []runtime.Program) {
	// load available programs, display selection menu
	fmt.Println("SELECT A PROGRAM # BELOW:")
	fmt.Println("------------------------------")
	for i, program := range programs {
		fmt.Printf("%d. %s\n", i+1, program.Name)
	}
}

func PromptSelection(programCount int) int {
	for {
		fmt.Print("Enter selection: ")
		line, err := readLine()
		if err != nil {
			fmt.Println("Input error. Try again.")
			continue
		}

		selection, err := strconv.Atoi(line)
		if err != nil || selection < 1 || selection > programCount {
			fmt.Printf("Please enter a number from 1 to %d.\n", programCount)
			continue
		}

		return selection
	}
}

func ContinuePrompt() bool {
	for {
		fmt.Print("CONTINUE? (y/q): ")
		continueSelection, err := readLine()
		if err != nil {
			fmt.Println("Input error. Exiting.")
			return false
		}

		switch strings.ToLower(continueSelection) {
		case "y":
			return true
		case "q":
			return false
		default:
			fmt.Println("Please type y to continue or q to quit.")
		}
	}
}


/*
lines to be printed:
starting
	STARTING PROGRAM ...
selection
	SELECT A PROGRAM # BELOW:
	------------------------------
	1. {program.name}
	2. {program.name}
after selection
	RUNNING {program.name} ...
run has finished
	{program.name} COMPLETED.
select a new one or quit?
	CONTINUE? (y/q)
*/