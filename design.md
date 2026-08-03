### Rule Definition Schema
Rules are stored as YAML files in the local `rules/` directory and loaded by the Runtime during startup. The Runtime scans the rules directory, parses rule definitions, validates them, and creates Rule Agents for execution.

Each rule should be defined in this way:
```
name: MaintainTankTemperature

enabled: true

target: process.Tank01.Temperature

condition:
    below: 50

severity: warning

recommendation:
    message: "Increase heater output by 10%"
```

The required fields include `name`, `enabled`, `target`, `condition`, and `recommendation`. Every rule must define the recommendation message used when an alert is generated.

#### Rule Field Definitions
`name`:
 - Required, unique field
 - Value: `string`

`enabled`:
 - Required field
 - Values: {`true`, `false`}

`target`:
 - Required field
 - Value: valid process.* state

`condition`:
 - Required field
 - Values: {`above`, `below`, `equals`}
 - Numeric comparison definition. Supported operators are above, below, and equals, each implemented as numeric (float64) comparisons.

`severity`:
 - Optional field that can provide the operator with more context
 - Values: {`advisory`, `warning`, `emergency`}
 - Default value: advisory

`recommendation`:
 - Required field
 - Value: {`message`}
 - Message value should be a `string`


### Shared Memory
Shared Memory is an in-memory key/value store used by the Runtime and Interpreter to maintain the latest process state during rule evaluation. Process values loaded from the Historian are stored using the process.* namespace. These values are then used as inputs during rule evaluation.

#### Process State
The process.* namespace contains process values loaded from the Historian database.

Naming follows the pattern: `process.<DeviceType><DeviceID>.<Variable>`. For example, `process.Pump03.Flow`, `process.Tank01.Level`, `process.Valve02.Position`, etc..

These values represent the current state of the physical process and are considered the authoritative source for runtime process conditions.

Example payload:
```
process.Tank01.Temperature
{
    "value": 45,
    "timestamp": "2026-07-17T10:15:00Z"
}
```

#### Alert State
Rule Agents generate Alert objects in memory when rule conditions evaluate to true. These alerts are collected by the Runtime and forwarded to the Alert Publisher for delivery to connected clients.

Example payload:
```
{
      "rule": "MaintainTankTemperature",
      "target": "process.Tank01.Temperature",
      "severity": "warning",
      "message": "Increase heater output by 10%"
}
```

### Shared Memory Contract
To prevent conflicting writes and establish clear responsibility boundaries, each agent type owns specific memory domains.

#### Runtime
The Runtime is responsible for:

      Loading rule definitions from the local rules directory
      Creating and registering Rule Agents
      Querying process values from the Historian SQLite database
      Coordinating rule execution
      Maintaining Shared Memory state
      Publishing generated alerts to the server

#### Rule Agents
Rule Agents are responsible for:

      Reading required process values from the Shared Memory
      Evaluating rule conditions
      Generating Alert objects when conditions are met

### Rule Agent Lifecycle
The Runtime dynamically creates Rule Agents from rule definitions stored in the local rules directory.

Startup:
```
Local Rules (*.yaml)
      ↓
Load Rule Definitions
      ↓
Validate Rule Schema
      ↓
Create Rule Agent
      ↓
Register Agent
```

Scan Cycle:
```
Historian SQLite
      ↓
Load Process Values
      ↓
Shared Memory
      ↓
Rule Agents
      ↓
Evaluate Conditions
      ↓
Generate Alerts
      ↓
Alert Publisher
```

At startup, the Runtime loads rule definitions from the local rules directory, validates them, and creates Rule Agents. During each scan cycle, process values are loaded from the Historian and stored in Shared Memory. Rule Agents evaluate rule conditions against the current process state and generate Alert objects when conditions are satisfied. Generated alerts are then published to connected clients through the server.

### Runtime Responsibilities
The Runtime is responsible for loading rule definitions from the local rules directory, validating rule schemas, creating and registering Rule Agents, maintaining shared process state, coordinating rule execution, collecting generated alerts, and publishing alerts to the server layer.

### Historian Integration Design
The Historian SQLite database is the source of process data. During execution, process values are loaded from the Historian and written into Shared Memory for use during rule evaluation. No dedicated Device Agent layer currently exists in the implementation.

### Overall Architecture
```
           Local Rules (*.yaml)
                   │
                   ▼
             Rule Loader
                   │
                   ▼
                Runtime
                   │
       ┌───────────┴───────────┐
       │                       │
       ▼                       ▼
  Rule Agents            Historian SQLite
       │                (LoadTimestamp)
       │                       │
       └───────────┬───────────┘
                   ▼
              Interpreter
                   │
                   │
       ┌───────────┴───────────┐
       ▼                       ▼
 Shared Memory              Alerts
(process.* keys)      (in-memory objects)
       │                       │
       └───────────┬───────────┘
                   ▼
            Alert Publisher
                   │
                   ▼
             HTTP/SSE Server
                   │
                   ▼
               Operator UI
```