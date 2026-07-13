package loader

import ("bufio"
		"os"
		"strings"
		"strconv"
		"process-engine/internal/engine")

// reads a program file & converts it into a list of instructions
func LoadProgram(path string) ([]engine.Instruction, error) {
	// open program file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// read line by line, parsing each line into an instruction struct
	var instructions []engine.Instruction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// skip blank lines & comments
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		lineNumber, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}

		// make actual instruction struct & add to list
		inst := engine.Instruction{
			Line: lineNumber,
			Opcode: engine.Opcode(strings.ToUpper(parts[1])),
			Args: parts[2:],
		}
		instructions = append(instructions, inst)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return instructions, nil
}