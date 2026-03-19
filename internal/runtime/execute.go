package runtime

import (
	"fmt"

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

func ExecutePlan(p *plan.Plan) error {
	rt := NewRuntime()
	executors := NewExecutors()

	for i, step := range p.Steps {
		executor, ok := executors[step.Type]
		if !ok {
			return fmt.Errorf("step %d: no executor registered for type %q", i+1, step.Type)
		}

		fmt.Printf("Running step %d: %s\n", i+1, step.Type)

		if err := executor.Execute(step, rt); err != nil {
			return fmt.Errorf("step %d (%s): %w", i+1, step.Type, err)
		}
	}

	return nil
}
