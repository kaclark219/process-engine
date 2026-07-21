package server

import ("encoding/json"
		"fmt"
		"net/http"
		"os"
		"path/filepath"
		"strconv"
		"sync"
		"process-engine/internal/agents")

type Server struct {
	mu sync.Mutex
	alertClients map[chan *agents.Alert]bool
	alertChan chan *agents.Alert
	rulesDir string
}

func NewServer(rulesDir string) *Server {
	return &Server{
		alertClients: make(map[chan *agents.Alert]bool),
		alertChan: make(chan *agents.Alert, 100),
		rulesDir: rulesDir,
	}
}

func (s *Server) Start(addr string) error {
	http.HandleFunc("/", s.handleIndex)
	http.HandleFunc("/api/rules", s.handleCreateRule)
	http.HandleFunc("/api/alerts", s.handleAlerts)
	http.HandleFunc("/api/alerts/stream", s.handleAlertsStream)

	go s.broadcastAlerts()

	return http.ListenAndServe(addr, nil)
}

// sends an alert to all connected clients
func (s *Server) PublishAlert(alert *agents.Alert) {
	s.alertChan <- alert
}

func (s *Server) broadcastAlerts() {
	for alert := range s.alertChan {
		s.mu.Lock()
		for client := range s.alertClients {
			select {
			case client <- alert:
			default:
				// skip
			}
		}
		s.mu.Unlock()
	}
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, htmlContent)
}

func (s *Server) handleCreateRule(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	// parse/validate form fields
	name := r.FormValue("name")
	enabled := r.FormValue("enabled") == "on"
	target := r.FormValue("target")
	conditionType := r.FormValue("conditionType")
	conditionValueStr := r.FormValue("conditionValue")
	severity := r.FormValue("severity")
	message := r.FormValue("message")

	if name == "" || target == "" || conditionType == "" || conditionValueStr == "" || message == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// parse condition value
	conditionValue, err := strconv.ParseFloat(conditionValueStr, 64)
	if err != nil {
		http.Error(w, "Invalid condition value", http.StatusBadRequest)
		return
	}
	var condition agents.Condition
	switch conditionType {
	case "above":
		condition.Above = &conditionValue
	case "below":
		condition.Below = &conditionValue
	case "equals":
		condition.Equals = &conditionValue
	default:
		http.Error(w, "Invalid condition type", http.StatusBadRequest)
		return
	}

	if severity == "" {
		severity = "advisory"
	}

	// create rule
	rule := agents.Rule{
		Name: name,
		Enabled: enabled,
		Target: target,
		Condition: condition,
		Severity: severity,
		Recommendation: agents.Recommendation{ Message: message },
	}

	// write yaml file
	filePath := filepath.Join(s.rulesDir, name+".yaml")
	err = s.writeRuleYAML(filePath, rule)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to save rule: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
		"name": name,
	})
}

func (s *Server) handleAlertsStream(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	clientChan := make(chan *agents.Alert, 10)
	s.mu.Lock()
	s.alertClients[clientChan] = true
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		delete(s.alertClients, clientChan)
		s.mu.Unlock()
		close(clientChan)
	}()

	// stream alerts
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	for {
		select {
		case alert, ok := <-clientChan:
			if !ok {
				return
			}
			data, _ := json.Marshal(alert)
			fmt.Fprintf(w, "data: %s\n\n", string(data))
			flusher.Flush()
		case <-r.Context().Done():
			return
		}
	}
}

func (s *Server) handleAlerts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "Use /api/alerts/stream for real-time alerts",
	})
}

func (s *Server) writeRuleYAML(filePath string, rule agents.Rule) error {
	yamlContent := fmt.Sprintf(`name: %s
		enabled: %v
		target: %s
		condition:
			%s: %v
		severity: %s
		recommendation:
			message: "%s"
		`,
		rule.Name,
		rule.Enabled,
		rule.Target,
		getConditionKey(&rule.Condition),
		getConditionValue(&rule.Condition),
		rule.Severity,
		rule.Recommendation.Message,
	)

	return os.WriteFile(filePath, []byte(yamlContent), 0644)
}

func getConditionKey(c *agents.Condition) string {
	if c.Above != nil {
		return "above"
	} else if c.Below != nil {
		return "below"
	}
	return "equals"
}

func getConditionValue(c *agents.Condition) string {
	if c.Above != nil {
		return fmt.Sprintf("%.2f", *c.Above)
	} else if c.Below != nil {
		return fmt.Sprintf("%.2f", *c.Below)
	}
	return fmt.Sprintf("%.2f", *c.Equals)
}
