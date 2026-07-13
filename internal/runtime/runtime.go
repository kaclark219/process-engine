package runtime

import ("os"
		"fmt")
import "process-engine/internal/engine"
import "process-engine/internal/loader"


// add a program model with name & path
type Program struct {
	Name string
	Path string
}

// add discover programs function that scans files under programs/ & returns []Program
func DiscoverPrograms() []Program {
	var programs []Program
	// scan programs/ directory for files, create Program struct for each file with name & path, add to programs list
	files, err := os.ReadDir("../programs")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		programs = append(programs, Program{
			Name: file.Name(),
			Path: "../programs/" + file.Name(),
		})
	}
	return programs
}

// add run program function that uses loader & executes instructions with interpreter
func RunProgram(program Program, interpreter *engine.Interpreter) {
	instructions, err := loader.LoadProgram(program.Path)
	if err != nil {
		fmt.Println("Error loading program:", err)
		return
	}
	for _, inst := range instructions {
		interpreter.Execute(inst)
	}
}
