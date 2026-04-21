package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/mlguels/amprig/internal/device"
	"github.com/mlguels/amprig/internal/orchestrator"
	"github.com/mlguels/amprig/internal/safety"
)

func main() {
	logger := slog.Default()
	slog.SetDefault(logger)

	slog.Info("Starting Amprig Orchestrator", "version", "1.0.0")

	newPSU := &device.SimulatedPSU{
		DeviceID:       "TEST-RIG-01",
		CurrentVoltage: 400.0,
		CurrentStatus:  device.Running,
	}

	slog.Info("device initialized status", "device_id", newPSU.DeviceID, "device_status", newPSU.CurrentStatus, "device_voltage", newPSU.CurrentVoltage)

	err := safety.ValidateDevice(newPSU)
	if err != nil {
		slog.Error("safety validation failed", "error", err)
	}

	slog.Info("device_info", "device_id", newPSU.DeviceID, "current_status", newPSU.CurrentStatus.String(), "current_voltage", newPSU.CurrentVoltage)

	if newPSU.CurrentVoltage == 0.0 && newPSU.CurrentStatus == device.Error {
		slog.Info("shutdown check successful", "state", "SAFE")
	} else {
		slog.Error("shutdown check failed", "state", "UNSAFE")
	}

	o := orchestrator.Orchestrator{}

	o.AddDevice(newPSU)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	newPSU.Start()
	if err := o.Run(ctx); err != nil {
		slog.Warn("orchestrator stopped", "error", err)
	}
}
