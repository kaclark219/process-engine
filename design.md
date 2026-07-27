### Rule Definition Schema
Rules are stored in Blob Storage as JSON or YAML and loaded dynamically at runtime.

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

The required fields include `name`, `enabled`, `target`, `condition`, and `recommendation`. Every rule must define the recommendation that should be presented to an operator when a violation occurs.

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
 - Comparison value type depends on the target variable

`severity`:
 - Optional field that can provide the operator with more context
 - Values: {`advisory`, `warning`, `emergency`}
 - Default value: advisory

`recommendation`:
 - Required field
 - Value: {`message`}
 - Message value should be a `string`


### Memory Namespaces
Shared Memory serves as the primary communication layer between Device Agents and Rule Agents. Shared Memory ensures all Rule Agents evaluate the same snapshot of process data during a scan cycle and avoids repeated Historian SQL queries across multiple agents. Data is organized into logical namespaces to separate process state, rule evaluation results, and operator recommendations.

Device Agents refresh process values from the Historian SQL database during each scan cycle and publish the latest state into Shared Memory. Shared Memory should be considered a runtime cache of the most recent process state rather than a long-term data store.

#### Process State
The process.* namespace contains real-time equipment and process values published by Device Agents.

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

#### Violation State
The violation.* namespace contains active rule violations detected during evaluation.

Naming follows the pattern: `violation.<ViolationID>`.

Violations provide traceability into why a recommendation was generated.

Example payload:
```
violation.01
{
    "rule": "MaintainTankTemperature",
    "target": "process.Tank01.Temperature",
    "actualValue": 45,
    "condition": {
        "below": 50
    },
    "timestamp": "2026-07-17T10:15:00Z"
}
```

A violation represents a rule condition that has evaluated to true and may result in one or more recommendations being generated.

#### Recommendation State
The recommendation.* namespace contains actions recommended by Rule Agents.

Naming follows the pattern: `recommendation.<RecommendationID>`.

Recommendations are intended for operator review and do not directly affect equipment operation.

Example payload:
```
recommendation.001
{
    "rule": "MaintainTankTemperature",
    "target": "process.Tank01.Temperature",
    "message": "Increase heater output by 10%",
    "severity": "warning",
    "timestamp": "2026-07-17T10:15:00Z"
}
```

### Shared Memory Contract
To prevent conflicting writes and establish clear responsibility boundaries, each agent type owns specific memory domains.

#### Device Agents
Device Agents are allowed to write:

    process.*

Device Agents are responsible for reading process data, publishing current process state, and updating equipment variables. Device Agents do not evaluate rules or generate recommendations.

#### Rule Agents
Rule Agents are allowed to write:

    violation.*
    recommendation.*

Rule Agents are responsible for evaluating user-defined rules, identifying violations, and generating recommendations. Rule definitions are loaded by the Runtime from Blob Storage and supplied to Rule Agents during creation.

### Rule Agent Lifecycle
The Runtime dynamically creates Rule Agents from rule definitions stored in Blob Storage.

Startup:
```
Blob Storage
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
Historian SQL
      ↓
Device Agents
      ↓
Shared Memory
      ↓
Rule Agents
      ↓
Evaluate Rules
      ↓
Create Violations
      ↓
Create Recommendations
```

At startup, the runtime should read blob storage, load rule definitions, and create rule agents. During each scan cycle, Device Agents retrieve the latest process values from the Historian SQL database and publish them to Shared Memory. Rule Agents then evaluate their rules against the current process state stored in Shared Memory. If there's a violation, one should be created in shared memory followed by a new recommendation.

### Runtime Responsibilities
The Runtime is responsible for loading rule definitions from Blob Storage, validating rule schemas, creating and registering Rule Agents, managing agent lifecycle, and coordinating Device Agent and Rule Agent scan cycles.

### Historian Integration Design
The Historian SQL database is the source of process data. Device Agents are responsible for querying the Historian and loading the latest process values into Shared Memory during each execution cycle.

### Overall Architecture
```
              Blob Storage
                    │
                    ▼
            Rule Definitions
                    │
                    ▼
                 Runtime
                    │
        ┌───────────┴───────────┐
        │                       │
        ▼                       ▼
 Device Agents            Rule Agents
        │                       │
        │                       ▼
        │                 Violations
        │                       │
        │                       ▼
        │               Recommendations
        │                       │
        ▼                       ▼
     Shared Memory ────────> Operator
        ▲
        │
        ▼
    Historian SQL
```

The YAML/JSON rule parser
The RuleDefinition model
The Shared Memory interfaces
The Runtime logic that creates Rule Agents from Blob-loaded definitions