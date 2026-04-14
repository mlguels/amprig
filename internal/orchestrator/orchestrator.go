package orchestrator

import (
	"context"
	"fmt"
	"time"

	"github.com/mlguels/amprig/internal/device"
	"github.com/mlguels/amprig/internal/safety"
)

type Orchestrator struct {
	Devices []device.Device
}

func (o *Orchestrator) AddDevice(d device.Device) {
	o.Devices = append(o.Devices, d)
	fmt.Printf("Devices: %v\n", d.GetID())
}

func (o *Orchestrator) Run(ctx context.Context) error {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("System Check: Validation all devices...")
			for _, device := range o.Devices {
				if err := safety.ValidateDevice(device); err != nil {
					ticker.Stop()
					return fmt.Errorf("safety check failed for: %s: %w\n", device.GetID(), err)
				}
			}

		case <-ctx.Done():
			fmt.Println("Orchestrator: shutdown signal recieved. Stopping...")
			for _, device := range o.Devices {
				device.EmergencyStop()
			}

			return ctx.Err()
		}
	}
}
