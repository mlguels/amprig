# Amprig Hardware Orchestrator

A simulation of a Power Supply Unit (PSU) orchestration system built in Go. This project demonstrates safety logic, concurrent telemetry monitoring, and structured observability.

## System Architecture

### 1. The Safety Engine ('internal/safety')
The core of the project is a Guard Clause validation system. It continuously monitors device states and enforces a "Fail-Safe" protocol:
- State Monitoring: Distinguishes between 'Idle', 'Running', 'Warning', and 'Error'.
- Automated Mitigation: Any device entering a 'Warning' or 'Error' state triggers an immediate 'EmergencyStop()', forcing voltage to zero to prevent simulated hardware damage.

### 2. Concurrency & Memory Safety ('internal/device')
To handle real-time data safely, the project implements:
- sync.Mutex: Ensures data consistency, preventing data races by locking the device state during updates.
- Goroutines: Decouples the hardware simulation from the control logic.
- Context API: Handles graceful shutdowns and timeout management for the control loop.

### 3. Observability & Testing
- Structured Logging: Utilizes 'log/slog' to show machine-readable telemetry data, great for log aggregation and real-time monitoring.
- Table-Driven Tests: Scalable testing that verifies all state transitions and safety violations across multiple scenarios.

## Getting Started

### Run the Orchestrator
```bash
go run cmd/amprig/main.go