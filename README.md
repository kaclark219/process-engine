# Rule-Driven Agent Platform Prototype
This repository contains a prototype rule-driven agent platform for monitoring process data, evaluating user-defined rules, and generating operator alerts.

The system loads rule definitions from local YAML files, retrieves process values from a Historian SQLite database, evaluates rule conditions through Rule Agents, and publishes alerts to connected clients through an HTTP/SSE interface.

Developers:

    - Katelyn Clark / Graduate Intern (2026)

## Repository Structure
```text
cmd/
├── Main application entry points used to start the runtime.

internal/
├── Core runtime implementation, including:
│   ├── Runtime orchestration
│   ├── Rule loading and validation
│   ├── Rule Agents
│   ├── Shared Memory
│   ├── Historian integration
│   ├── Alert publishing
│   └── HTTP/SSE server components

rules/
├── YAML rule definitions loaded during startup.
```

## High-Level Workflow
1. Load rule definitions from the local `rules/` directory.
2. Create and register Rule Agents.
3. Query process data from the Historian SQLite database.
4. Store process values in Shared Memory.
5. Evaluate loaded rules against the current process state.
6. Generate alerts when rule conditions are met.
7. Stream alerts to connected clients through the HTTP/SSE server.

## Developer Setup
### Prerequisites
- Go 1.26.x (module currently targets `go 1.26.5`)
- Write access to `C:/sqlite/`

### Dependencies
From the repository root:
```powershell
go mod download
```

### Historian Database Setup
The runtime reads Historian data from this fixed path:

    `C:/sqlite/historian.db`

Create the directory and database table:
```powershell
New-Item -ItemType Directory -Path C:/sqlite -Force | Out-Null
sqlite3 C:/sqlite/historian.db "CREATE TABLE IF NOT EXISTS ProcessData (ProcessName TEXT, VariableName TEXT, Value REAL, TimestampUTC TEXT);"
```

Insert seed data that matches the runtime start timestamp (`2026-07-17T10:15:00Z`):
```powershell
sqlite3 C:/sqlite/historian.db "DELETE FROM ProcessData WHERE TimestampUTC='2026-07-17T10:15:00Z';"
sqlite3 C:/sqlite/historian.db "INSERT INTO ProcessData (ProcessName, VariableName, Value, TimestampUTC) VALUES ('Tank01','Temperature',45,'2026-07-17T10:15:00Z');"
sqlite3 C:/sqlite/historian.db "INSERT INTO ProcessData (ProcessName, VariableName, Value, TimestampUTC) VALUES ('Tank01','Level',70,'2026-07-17T10:15:00Z');"
```

Notes:
- The runtime advances one minute per tick, so add more timestamps if you want sustained alert output.
- If `sqlite3` CLI is not installed, install it or create/populate the same table with any SQLite client.

## Build + Run
Build the binary into `bin/`:
```powershell
go build -o bin/process-engine.exe ./cmd
```

Run the built executable:
```powershell
./bin/process-engine.exe
```

Expected startup output includes `Starting server on http://localhost:8080`.

Open the UI at `http://localhost:8080`.

## Development Notes

### Default Scan Behavior
- Runtime ticks once per second.
- Historian data is loaded each tick for the current timestamp.
- Rule scans run every third tick.

## API Endpoints
- `GET /` - Web UI
- `POST /api/rules` - Create and save a new rule YAML file
- `GET /api/alerts` - Alerts endpoint info
- `GET /api/alerts/stream` - Server-Sent Events stream for live alerts
