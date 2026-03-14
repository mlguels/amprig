package runtime

import "github.com/mlguels/amprig/internal/plan"

type Executor interface {
	Execute(step plan.Step, rt *Runtime) error
}
