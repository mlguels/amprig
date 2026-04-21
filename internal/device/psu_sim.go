package device

import (
	"log/slog"
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

// If triggered sets voltage to 0
// Sets status to Error
func (s *SimulatedPSU) EmergencyStop() error {
	s.CurrentVoltage = 0
	s.CurrentStatus = Error

	slog.Error("emergency stop triggered",
		"device_id", s.DeviceID,
		"status", s.CurrentStatus.String(),
	)
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
