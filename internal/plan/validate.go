package plan

import (
	"fmt"
)

func Validate(plan *Plan) error {
	if plan == nil {
		return fmt.Errorf("plan is nil")
	}
	// Plan validation
	if plan.Name == "" {
		return fmt.Errorf("plan name is empty")
	}
	if plan.Version != 1 {
		return fmt.Errorf("plan version must be 1, got %d", plan.Version)
	}
	if len(plan.Steps) == 0 {
		return fmt.Errorf("plan must have at least 1 step")
	}

	// Steps validation
	for i, step := range plan.Steps {
		stepNum := i + 1

		switch step.Type {
		case "set_voltage":
			if step.ValueV <= 0 {
				return fmt.Errorf("step %d (set_voltage): value_v must be > 0", stepNum)
			}
			if step.CurrentLimitA <= 0 {
				return fmt.Errorf("step %d (set_voltage): current_limit_a must be > 0", stepNum)
			}
		case "wait":
			if step.Duration <= 0 {
				return fmt.Errorf("step %d (duration): duration must be > 0", stepNum)
			}
		case "measure":
			if len(step.Metrics) <= 0 {
				return fmt.Errorf("step %d (metrics): metrics must be > 0", stepNum)
			}
		case "assert":
			if step.Metric == "" {
				return fmt.Errorf("step %d (metric): metric must not be empty", stepNum)
			}
			if step.Op == "" {
				return fmt.Errorf("step %d (metric): op must not be empty", stepNum)
			}
			if step.Value == 0 {
				return fmt.Errorf("step %d (metric): value must be > 0", stepNum)
			}

			switch step.Op {
			case "<", "<=", ">", ">=", "==", "!=":
			default:
				return fmt.Errorf("step %d (assert): invalid operator %q", stepNum, step.Op)
			}

		default:
			return fmt.Errorf("step %d: unknown step unknown step type %q", stepNum, step.Type)
		}
	}
	return nil
}
