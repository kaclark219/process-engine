package loader

import ("os"
		"path/filepath"
		"process-engine/internal/agents")


func LoadRuleAgents(rulesDir string) ([]agents.Agent, error) {
    var agentsList []agents.Agent

    entries, err := os.ReadDir(rulesDir)
    if err != nil {
        return nil, err
    }
    for _, entry := range entries {
        if entry.IsDir() {
            continue
        }
        rulePath := filepath.Join(rulesDir, entry.Name())
        agentsList = append(
            agentsList,
            agents.NewRuleAgent(entry.Name(), rulePath),
        )
    }
    return agentsList, nil
}