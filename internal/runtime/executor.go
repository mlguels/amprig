package runtime

import (
	"fmt"
	"time"

	"github.com/mlguels/amprig/internal/plan"
)

type Executor interface {
	Execute(step plan.Step, rt *Runtime) error
}

type SetVoltageExecutor struct{}

func (e SetVoltageExecutor) Execute(step plan.Step, rt *Runtime) error {
	rt.currentVoltageV = step.ValueV
	rt.currentLimitA = step.CurrentLimitA

	fmt.Printf("Set voltage to %.2f V (limit %.2f A)\n", rt.currentVoltageV, rt.currentLimitA)
	return nil
}

type WaitExecutor struct{}

func (e WaitExecutor) Execute(step plan.Step, rt *Runtime) error {
	fmt.Printf("Waiting for %s\n", step.Duration)
	time.Sleep(step.Duration)
	return nil
}

type MeasureExecutor struct{}

func (e MeasureExecutor) Execute(step plan.Step, rt *Runtime) error {
	for _, metric := range step.Metrics {
		switch metric {
		case "voltage_v":
			rt.measurements[metric] = rt.currentVoltageV
		case "current_a":
			rt.measurements[metric] = rt.currentLimitA
		default:
			return fmt.Errorf("Metric not found %q", metric)
		}
	}
	return nil
}
