package safety

import (
	"testing"

	"github.com/mlguels/amprig/internal/device"
)

func TestValidateDeviceTableDriven(t *testing.T) {
	var tests = []struct {
		name       string
		status     device.Status
		shouldFail bool
	}{
		{"Idle_ShouldPass", device.Idle, false},
		{"Running_ShouldPass", device.Running, false},
		{"Warning_ShouldFail", device.Warning, true},
		{"Error,ShouldFail", device.Error, true},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {
			psu := &device.SimulatedPSU{
				DeviceID:      "TEST_DEVICE",
				CurrentStatus: tc.status,
			}
			err := ValidateDevice(psu)

			gotError := err != nil
			if gotError != tc.shouldFail {
				t.Errorf("Test %s failed: expected error=%v, but got error=%v (err: %v)", tc.name, tc.shouldFail, gotError, err)
			}
		})

	}
}
