package agents

type ValveAgent struct {
	name string
	program string
}

func (v *ValveAgent) Name() string {
	return v.name
}

func (v *ValveAgent) Program() string {
	return v.program
}
