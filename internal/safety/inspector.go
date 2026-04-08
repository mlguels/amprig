package safety

import (
	"github.com/mlguels/amprig/internal/device"
)

// this function will validate the device
func ValidateDevice(d device.Device) error {
	currentStatus := d.GetStatus()
	// if there is an error or warning we call the emergencyStop() function
	if currentStatus == device.Error || currentStatus == device.Warning {
		return d.EmergencyStop()
	}

	return nil
}
