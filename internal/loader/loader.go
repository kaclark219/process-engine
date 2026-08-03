package loader

import ("fmt"
        "os"
		"path/filepath"
		"process-engine/internal/agents")


func LoadRuleAgents(rulesDir string) ([]agents.Agent, error) {
    var agentsList []agents.Agent
    seenRuleNames := map[string]string{}

    entries, err := os.ReadDir(rulesDir)
    if err != nil {
        return nil, err
    }
    for _, entry := range entries {
        if entry.IsDir() {
            continue
        }
        rulePath := filepath.Join(rulesDir, entry.Name())

		rule := agents.ProcessRule(rulePath)
		if existingPath, exists := seenRuleNames[rule.Name]; exists {
			return nil, fmt.Errorf("rule name %q must be unique (found in %s and %s)", rule.Name, existingPath, rulePath)
		}
		seenRuleNames[rule.Name] = rulePath

        agentsList = append(
            agentsList,
            agents.NewRuleAgent(entry.Name(), rulePath),
        )
    }
    return agentsList, nil
}