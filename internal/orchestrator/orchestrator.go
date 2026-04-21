package orchestrator

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/mlguels/amprig/internal/device"
	"github.com/mlguels/amprig/internal/safety"
)

type Orchestrator struct {
	Devices []device.Device
}

func (o *Orchestrator) AddDevice(d device.Device) {
	o.Devices = append(o.Devices, d)
	// fmt.Printf("Devices: %v\n", d.GetID())
	slog.Info("device added to system", "device_id", d.GetID())
}

func (o *Orchestrator) Run(ctx context.Context) error {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			slog.Debug("starting system safety validation")
			for _, device := range o.Devices {
				if err := safety.ValidateDevice(device); err != nil {
					ticker.Stop()
					return fmt.Errorf("safety check failed for: %s: %w\n", device.GetID(), err)
				}
				tel := device.GetTelemetry()
				// fmt.Printf("📊 [%s] Telemetry: %.2fV | %.2fA\n", device.GetID(), tel.Voltage, tel.Current)
				slog.Info(
					"telemetry updated",
					"device_id", device.GetID(),
					"voltage", tel.Voltage,
					"amps", tel.Current,
				)
			}

		case <-ctx.Done():
			// fmt.Println("Orchestrator: shutdown signal recieved. Stopping...")
			slog.Warn("shutdown signal recieved, stopping all devices")
			for _, device := range o.Devices {
				device.EmergencyStop()
			}

			return ctx.Err()
		}
	}
}
