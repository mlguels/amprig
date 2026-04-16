package device

import (
	"fmt"
	"sync"
	"time"
)

// We have a struct that simulates a power supply unit (could be any type of machine)
// CurrentStatus is a custom type Status
type SimulatedPSU struct {
	mu             sync.Mutex
	DeviceID       string
	CurrentVoltage float64
	MaxVoltage     float64
	CurrentStatus  Status
	CurrentAmps    float64
}

// These methods take in a pointer to avoid copying and to
// stay consistent with other methods changing the state
func (s *SimulatedPSU) GetID() string {
	// We return the DeviceID
	return s.DeviceID
}

func (s *SimulatedPSU) GetStatus() Status {
	return s.CurrentStatus
}

// This method handles the EmergencyStop() func of the interface
func (s *SimulatedPSU) EmergencyStop() error {
	// We are setting the CurrentVoltage to 0 (simulating turning off the device)
	s.CurrentVoltage = 0
	// Then we set the status to Error which is in the type Status
	s.CurrentStatus = Error
	// We print it to the console with the deviceid
	fmt.Printf("[SAFEGUARD Emergency stop for device: %s]\n", s.DeviceID)

	return nil
}

func (s *SimulatedPSU) Start() {

	go func() {
		tick := time.NewTicker(100 * time.Millisecond)
		defer tick.Stop()
		for range tick.C {
			s.mu.Lock()

			if s.CurrentStatus == Running {
				s.CurrentVoltage += 0.02
				s.CurrentAmps += 0.01
			}

			s.mu.Unlock()
		}
	}()

}

func (s *SimulatedPSU) GetTelemetry() Telemetry {
	s.mu.Lock()
	defer s.mu.Unlock()

	return Telemetry{
		Voltage: s.CurrentVoltage,
		Current: s.CurrentAmps,
	}
}
