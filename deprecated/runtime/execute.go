package runtime

import (
	"fmt"
	"time"

	"github.com/mlguels/amprig/internal/plan"
)

func NewExecutors() map[string]Executor {
	return map[string]Executor{
		"set_voltage": SetVoltageExecutor{},
		"wait":        WaitExecutor{},
		"measure":     MeasureExecutor{},
		"assert":      AssertExecutor{},
	}
}

func ExecutePlan(p *plan.Plan) (*PlanResult, error) {
	rt := NewRuntime()
	executors := NewExecutors()

	result := &PlanResult{
		PlanName:  p.Name,
		StartedAt: time.Now(),
	}

	for i, step := range p.Steps {
		executor, ok := executors[step.Type]
		if !ok {
			result.Success = false
			result.FinishedAt = time.Now()
			return result, fmt.Errorf("step %d: no exectuor registered for type %q", i+1, step.Type)
		}

		fmt.Printf("Running step %d: %s\n", i+1, step.Type)

		startTimer := time.Now()
		err := executor.Execute(step, rt)
		duration := time.Since(startTimer)

		if err != nil {
			stepResult := StepResult{
				StepNumber: i + 1,
				StepType:   step.Type,
				Status:     "failed",
				Message:    err.Error(),
				Duration:   duration,
			}

			result.Steps = append(result.Steps, stepResult)
			result.Success = false
			result.FinishedAt = time.Now()

			return result, fmt.Errorf("step %d (%s): %w", i+1, step.Type, err)
		}

		stepResult := StepResult{
			StepNumber: i + 1,
			StepType:   step.Type,
			Status:     "passed",
			Message:    "Passed",
			Duration:   duration,
		}

		result.Steps = append(result.Steps, stepResult)
	}

	result.Success = true
	result.FinishedAt = time.Now()

	return result, nil
}
