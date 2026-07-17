package agents


type Agent struct {
	name string
	program string
}

func NewAgent(name string, program string) *Agent {
	return &Agent{
		name: name,
		program: program,
	}
}

func (a *Agent) Name() string {
	return a.name
}

func (a *Agent) Program() string {
	return a.program
}