package device

import (
	"testing"
)

func TestEmergencyStop(t *testing.T) {
	psu := &SimulatedPSU{
		DeviceID:       "TEST-01",
		CurrentVoltage: 400.0,
		CurrentStatus:  Running,
	}

	psu.EmergencyStop()

	if psu.CurrentVoltage != 0 {
		t.Errorf("Expecting voltage 0, got %.2f", psu.CurrentVoltage)
	}

	if psu.CurrentStatus != Error {
		t.Errorf("Expecting Error %s, got %s", Error, psu.CurrentStatus)
	}

}
