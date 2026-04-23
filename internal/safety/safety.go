package safety

import (
	"fmt"

	"github.com/mlguels/amprig/internal/device"
)

// Validates devices by checking for status error/warning
// Instantly called EmergencyStop()
// Returns Error
func ValidateDevice(d device.Device) error {
	status := d.GetStatus()
	if status == device.Error || status == device.Warning {
		d.EmergencyStop()
		return fmt.Errorf("safety violation: device %s is in %s state", d.GetID(), status)
	}
	return nil
}
