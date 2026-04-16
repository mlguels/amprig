package safety

import (
	"fmt"

	"github.com/mlguels/amprig/internal/device"
)

// this function will validate the device
func ValidateDevice(d device.Device) error {
	status := d.GetStatus()
	// if there is an error or warning we call the emergencyStop() function
	if status == device.Error || status == device.Warning {
		d.EmergencyStop()
		return fmt.Errorf("safety violation: device %s is in %s state", d.GetID(), status)
	}
	return nil
}
