package server

const htmlContent = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Process Engine Control Panel</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        
        body {
            font-family: Arial, sans-serif;
            background: #a1a3a573;
            min-height: 100vh;
			display: flex;
			justify-content: center;
			align-items: center;
            padding: 20px;
        }
        
		.container {
			width: 100%;
			max-width: 1200px;
			display: grid;
			grid-template-columns: 1fr 1fr;
			gap: 20px;
		}
        
        .panel {
            background: white;
            border-radius: 6px;
            padding: 25px;
        }
        
        .panel h2 {
            color: #004487;
            margin-bottom: 20px;
            font-size: 1.5em;
            padding-bottom: 10px;
        }
        
        .form-group {
            margin-bottom: 15px;
        }
        
        label {
            display: block;
            font-weight: 600;
            color: #555;
            margin-bottom: 5px;
            font-size: 0.9em;
        }
        
        input, select {
            width: 100%;
            padding: 10px 12px;
            border: 1px solid #ddd;
            border-radius: 5px;
            font-size: 1em;
            transition: border-color 0.3s;
        }
        
        input:focus, select:focus {
            outline: none;
            border-color: #004b8d;
        }
        
        .checkbox-group {
            display: flex;
            align-items: center;
            gap: 10px;
        }
        
        input[type="checkbox"] {
            width: auto;
            cursor: pointer;
        }
        
        button {
            background: #004b8d;
            color: white;
            padding: 12px 24px;
            border: none;
            border-radius: 5px;
            font-size: 1em;
            font-weight: 600;
            cursor: pointer;
            transition: transform 0.2s, box-shadow 0.2s;
            width: 100%;
        }
        
        button:hover {
            background: #1b2552;
        }
        
        .alerts-container {
            max-height: 80vh;
            overflow-y: auto;
        }
        
        .alert-item {
            background: #f8f9fa;
            border-left: 4px solid #ddd;
            padding: 12px;
            margin-bottom: 10px;
            border-radius: 3px;
            font-size: 0.9em;
        }
        
        .alert-item.warning {
            border-left-color: #d6ba01;
            background: #fff5b57e;
        }
        
        .alert-item.emergency {
            border-left-color: #e35b39;
            background: rgba(245, 195, 183, 0.51);
        }
        
        // .alert-item.advisory {
        //     border-left-color: #00573d;
        //     background: #7ccf8b7c;
        // }

        .alert-item.advisory {
            border-left-color: #004b8d;
            background: #e3f2ff;
        }

        .alert-item.waiting {
            border-left-color: #004b8d;
            background: #e3f2ff;
        }
        
        .alert-severity {
            font-weight: 700;
            text-transform: uppercase;
            font-size: 0.8em;
            margin-bottom: 5px;
        }
        
        .alert-severity.warning {
            color: #d6ba01;
        }
        
        .alert-severity.emergency {
            color: #e35b39;
        }
        
        // .alert-severity.advisory {
        //     color: #00573d;
        // }

        .alert-severity.advisory {
            color: #004b8d;
        }
        
        .alert-severity.waiting {
            color: #004b8d;
        }
        
        .alert-rule {
            font-weight: 600;
            color: #333;
            margin-bottom: 3px;
        }
        
        .alert-details {
            font-size: 0.85em;
            color: #666;
            line-height: 1.4;
        }
        
        .status-message {
            padding: 10px;
            border-radius: 5px;
            margin-bottom: 15px;
            text-align: center;
            font-weight: 500;
        }
        
        .status-message.success {
            background: #c8e6c9;
            color: #2e7d32;
        }
        
        .status-message.error {
            background: #ffcdd2;
            color: #c62828;
        }
        
        .status-message.info {
            background: #bbdefb;
            color: #1565c0;
        }
        
        @media (max-width: 768px) {
            .container {
                grid-template-columns: 1fr;
            }
        }

        .alert-item {
            background: #f8f9fa;
            border-left: 4px solid #ddd;
            padding: 12px;
            margin-bottom: 10px;
            border-radius: 3px;
            font-size: 0.9em;
        }

        .alert-item.new-alert {
            animation: alertAppear 0.5s ease-out;
        }

        @keyframes alertAppear {
            from {
                opacity: 0;
                transform: translateY(-20px);
            }
            to {
                opacity: 1;
                transform: translateY(0);
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <!-- Create Rule Panel -->
        <div class="panel">
            <h2>Create New Rule</h2>
            <form id="ruleForm">
                <div id="statusMessage"></div>
                
                <div class="form-group">
                    <label for="name">Rule Name *</label>
                    <input type="text" id="name" name="name" placeholder="e.g., MaintainTankTemperature" required>
                </div>
                
                <div class="form-group">
                    <label for="target">Target Process Variable *</label>
                    <input type="text" id="target" name="target" placeholder="e.g., process.Tank01.Temperature" required>
                </div>
                
                <div class="form-group">
                    <label for="conditionType">Condition Type *</label>
                    <select id="conditionType" name="conditionType" required>
                        <option value="">-- Select --</option>
                        <option value="above">Above</option>
                        <option value="below">Below</option>
                        <option value="equals">Equals</option>
                    </select>
                </div>
                
                <div class="form-group">
                    <label for="conditionValue">Threshold Value *</label>
                    <input type="number" id="conditionValue" name="conditionValue" step="0.01" placeholder="e.g., 50" required>
                </div>
                
                <div class="form-group">
                    <label for="severity">Severity Level</label>
                    <select id="severity" name="severity">
                        <option value="advisory">Advisory</option>
                        <option value="warning">Warning</option>
                        <option value="emergency">Emergency</option>
                    </select>
                </div>
                
                <div class="form-group">
                    <label for="message">Recommendation Message *</label>
                    <input type="text" id="message" name="message" placeholder="e.g., Increase heater output by 10%" required>
                </div>
                
                <div class="form-group checkbox-group">
                    <input type="checkbox" id="enabled" name="enabled" checked>
                    <label for="enabled" style="margin: 0;">Enabled</label>
                </div>
                
                <button type="submit">Create Rule</button>
            </form>
        </div>
        
        <!-- Alerts Panel -->
        <div class="panel">
            <h2>Real-Time Alerts</h2>
            <div class="alerts-container" id="alertsContainer">
                <div class="alert-item waiting" id="connectingMessage">
                    <div class="alert-severity waiting">Connecting...</div>
                    <div class="alert-rule">Waiting for alerts</div>
                    <div class="alert-details">Connecting to alert stream...</div>
                </div>
            </div>
        </div>
    </div>
    
    <script>
        const ruleForm = document.getElementById('ruleForm');
        const statusMessage = document.getElementById('statusMessage');
        const alertsContainer = document.getElementById('alertsContainer');
        let alertCount = 0;
        
        // Form submission
        ruleForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            
            const formData = new FormData(ruleForm);
            
            try {
                const response = await fetch('/api/rules', {
                    method: 'POST',
                    body: formData
                });
                
                if (response.ok) {
                    const data = await response.json();
                    showStatus('Rule created successfully!', 'success');
                    ruleForm.reset();
                } else {
                    showStatus('Failed to create rule', 'error');
                }
            } catch (error) {
                console.error('Error:', error);
                showStatus('Error creating rule: ' + error.message, 'error');
            }
        });
        
        function showStatus(message, type) {
            statusMessage.textContent = message;
            statusMessage.className = 'status-message ' + type;
            setTimeout(() => {
                statusMessage.textContent = '';
                statusMessage.className = '';
            }, 5000);
        }
        
        // SSE Connection for alerts
        function connectToAlerts() {
            const eventSource = new EventSource('/api/alerts/stream');
            
            eventSource.onmessage = (event) => {
                try {
                    const alert = JSON.parse(event.data);
                    addAlertToUI(alert);
                } catch (error) {
                    console.error('Failed to parse alert:', error);
                }
            };
            
            eventSource.onerror = (error) => {
                console.error('SSE Error:', error);
                eventSource.close();
                // Reconnect after 3 seconds
                setTimeout(connectToAlerts, 3000);
            };
        }
        
        function addAlertToUI(alert) {
            // Remove connecting message on first alert
            if (alertCount === 0) {
                alertsContainer.innerHTML = '';
            }
            alertCount++;

            const knownSeverities = ['advisory', 'warning', 'emergency'];
            const rawSeverity = (alert.Severity || '').toLowerCase();
            const severityClass = knownSeverities.includes(rawSeverity) ? rawSeverity : 'waiting';
            const severityLabel = rawSeverity ? rawSeverity.toUpperCase() : 'INFO';
            
            const alertDiv = document.createElement('div');
            alertDiv.className = 'alert-item ' + severityClass + ' new-alert';
            
            const timestamp = new Date().toLocaleTimeString();
            
            const valueFixed = typeof alert.Value === 'number' && Number.isFinite(alert.Value)
                ? alert.Value.toFixed(2)
                : 'N/A';
            const ruleName = alert.RuleName || 'System';
            const message = alert.Message || 'Status update';
            const target = alert.Target || 'N/A';
            const condition = alert.Condition || 'N/A';
            
            alertDiv.innerHTML = '<div class="alert-severity ' + severityClass + '">' + severityLabel + '</div>' +
                '<div class="alert-rule">' + ruleName + '</div>' +
                '<div class="alert-details">' +
                '<strong>' + message + '</strong><br>' +
                'Target: ' + target + '<br>' +
                'Current Value: ' + valueFixed + '<br>' +
                'Condition: ' + condition + '<br>' +
                '<em style="color: #999;">' + timestamp + '</em>' +
                '</div>';
            
            alertsContainer.insertBefore(alertDiv, alertsContainer.firstChild);

            alertDiv.addEventListener('animationend', () => {
                alertDiv.classList.remove('new-alert');
            });
            
            // keep only last 50 alerts
            while (alertsContainer.children.length > 50) {
                alertsContainer.removeChild(alertsContainer.lastChild);
            }
        }
        
        // Connect to alerts stream on page load
        window.addEventListener('load', connectToAlerts);
    </script>
</body>
</html>
`
