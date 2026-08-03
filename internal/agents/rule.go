package agents

type Rule struct {
	Name string `yaml:"name"`
	Enabled *bool `yaml:"enabled"`
	Target string `yaml:"target"`
	Condition Condition `yaml:"condition"`
	Severity string `yaml:"severity"`
	Recommendation Recommendation `yaml:"recommendation"`
}

type Condition struct {
	Above  *float64 `yaml:"above,omitempty"`
	Below  *float64 `yaml:"below,omitempty"`
	Equals *float64 `yaml:"equals,omitempty"`
}

type Recommendation struct {
	Message string `yaml:"message"`
}

type Alert struct {
	Severity string
	RuleName string
	Target string
	Value float64
	Condition string
	Message string
}

func (r Rule) IsEnabled() bool {
	if r.Enabled == nil {
		return false
	}
	return *r.Enabled
}