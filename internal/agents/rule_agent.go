package agents

import ("fmt"
		"os"
		"process-engine/internal/memory"
		"regexp"
		"gopkg.in/yaml.v3")

var processTargetPattern = regexp.MustCompile(`^process\.[A-Za-z0-9_]+\.[A-Za-z0-9_]+$`)

type RuleAgent struct {
	*BaseAgent
	rule Rule
}

func NewRuleAgent(name string, rule_path string) *RuleAgent {
	rule := ProcessRule(rule_path)

	return &RuleAgent{
		BaseAgent: NewAgent(name, rule_path),
		rule:      rule,
	}
}

func ProcessRule(rule_path string) Rule {
	data, err := os.ReadFile(rule_path)
	if err != nil {
		panic(err)
	}

	var rule Rule
	if err := yaml.Unmarshal(data, &rule); err != nil {
		panic(err)
	}

	if rule.Severity == "" {
		rule.Severity = "advisory"
	}

	// validation
	if rule.Name == "" {
		panic("Rule validation failed: name is required")
	}
	if rule.Enabled == nil {
		panic("Rule validation failed: enabled is required")
	}
	if rule.Target == "" {
		panic("Rule validation failed: target is required")
	}
	if !processTargetPattern.MatchString(rule.Target) {
		panic("Rule validation failed: target must match process.<DeviceType><DeviceID>.<Variable>")
	}
	if rule.Recommendation.Message == "" {
		panic("Rule validation failed: recommendation.message is required")
	}
	conditionCount := 0
	if rule.Condition.Above != nil {
		conditionCount++
	}
	if rule.Condition.Below != nil {
		conditionCount++
	}
	if rule.Condition.Equals != nil {
		conditionCount++
	}
	if conditionCount != 1 {
		panic("Rule validation failed: condition must define exactly one of above, below, or equals")
	}
	valid_severity := rule.Severity == "advisory" || rule.Severity == "warning" || rule.Severity == "emergency"
	if !valid_severity {
		panic("Rule validation failed: severity must be one of advisory, warning, or emergency")
	}

	return rule
}

func (a *RuleAgent) Scan(memory *memory.Memory) []Alert {
	if !a.rule.IsEnabled() {
		return nil
	}

	// evaluate rule against current memory state
	target := a.rule.Target
	value := memory.Get(target)
	if valueMap, ok := value.(map[string]any); ok {
		if val, ok := valueMap["value"].(float64); ok {
			if a.rule.Condition.Above != nil && val > *a.rule.Condition.Above {
				return []Alert{
					{
						Severity: a.rule.Severity,
						RuleName: a.rule.Name,
						Target: a.rule.Target,
						Value: val,
						Condition: fmt.Sprintf("Above %.2f", *a.rule.Condition.Above),
						Message: a.rule.Recommendation.Message,
					},
				}
			} else if a.rule.Condition.Below != nil && val < *a.rule.Condition.Below {
				return []Alert{
					{
						Severity: a.rule.Severity,
						RuleName: a.rule.Name,
						Target: a.rule.Target,
						Value: val,
						Condition: fmt.Sprintf("Below %.2f", *a.rule.Condition.Below),
						Message: a.rule.Recommendation.Message,
					},
				}
			} else if a.rule.Condition.Equals != nil && val == *a.rule.Condition.Equals {
				return []Alert{
					{
						Severity: a.rule.Severity,
						RuleName: a.rule.Name,
						Target: a.rule.Target,
						Value: val,
						Condition: fmt.Sprintf("Equals %.2f", *a.rule.Condition.Equals),
						Message: a.rule.Recommendation.Message,
					},
				}
			} else {
				fmt.Println("Rule not triggered")
			}
		}
	}
	return nil
}

func PrintRule(a RuleAgent) {
	fmt.Println("Name:", a.rule.Name)
	fmt.Println("Enabled:", a.rule.IsEnabled())
	fmt.Println("Target:", a.rule.Target)
	if a.rule.Condition.Above != nil {
		fmt.Println("Condition Above:", *a.rule.Condition.Above)
	}
	if a.rule.Condition.Below != nil {
		fmt.Println("Condition Below:", *a.rule.Condition.Below)
	}
	if a.rule.Condition.Equals != nil {
		fmt.Println("Condition Equals:", *a.rule.Condition.Equals)
	}
	fmt.Println("Severity:", a.rule.Severity)
	fmt.Println("Recommendation Message:", a.rule.Recommendation.Message)
}
