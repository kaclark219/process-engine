package agents

type TankAgent struct {
	name string
	program string
}

func (t *TankAgent) Name() string {
	return t.name
}

func (t *TankAgent) Program() string {
	return t.program
}