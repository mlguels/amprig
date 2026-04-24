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
type WaitExecutor struct{}
type MeasureExecutor struct{}
type AssertExecutor struct{}

func (e SetVoltageExecutor) Execute(step plan.Step, rt *Runtime) error {
	rt.currentVoltageV = step.ValueV
	rt.currentLimitA = step.CurrentLimitA

	fmt.Printf("Set voltage to %.2f V (limit %.2f A)\n", rt.currentVoltageV, rt.currentLimitA)
	return nil
}

func (e WaitExecutor) Execute(step plan.Step, rt *Runtime) error {
	fmt.Printf("Waiting for %s\n", step.Duration)
	time.Sleep(step.Duration)
	return nil
}

func (e MeasureExecutor) Execute(step plan.Step, rt *Runtime) error {
	for _, metric := range step.Metrics {
		switch metric {
		case "voltage_v":
			rt.measurements[metric] = rt.currentVoltageV
			fmt.Printf("Measured %s = %.2f\n", metric, rt.currentVoltageV)
		case "current_a":
			rt.measurements[metric] = rt.currentLimitA
			fmt.Printf("Measured %s = %.2f\n", metric, rt.currentLimitA)
		default:
			return fmt.Errorf("metric not found %q", metric)
		}
	}
	return nil
}

func (e AssertExecutor) Execute(step plan.Step, rt *Runtime) error {
	value, ok := rt.measurements[step.Metric]
	if !ok {
		return fmt.Errorf("metric %q not found in measurements", step.Metric)
	}

	switch step.Op {
	case "<":
		if !(value < step.Value) {
			return fmt.Errorf("assert failed: %s %.2f < %.2f", step.Metric, value, step.Value)
		}
	case "<=":
		if !(value <= step.Value) {
			return fmt.Errorf("assert failed: %s %.2f <= %.2f", step.Metric, value, step.Value)
		}
	case ">":
		if !(value > step.Value) {
			return fmt.Errorf("assert failed: %s %.2f > %.2f", step.Metric, value, step.Value)
		}
	case ">=":
		if !(value >= step.Value) {
			return fmt.Errorf("assert failed: %s %.2f >= %.2f", step.Metric, value, step.Value)
		}
	case "==":
		if !(value == step.Value) {
			return fmt.Errorf("assert failed: %s %.2f == %.2f", step.Metric, value, step.Value)
		}
	case "!=":
		if !(value != step.Value) {
			return fmt.Errorf("assert failed: %s %.2f != %.2f", step.Metric, value, step.Value)
		}
	default:
		return fmt.Errorf("invalid operator: %q", step.Op)
	}
	fmt.Printf("Assert success: %s %s %.2f\n", step.Metric, step.Op, step.Value)
	return nil
}
