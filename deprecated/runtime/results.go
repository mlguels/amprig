package runtime

import "time"

type PlanResult struct {
	PlanName   string
	StartedAt  time.Time
	FinishedAt time.Time
	Success    bool
	Steps      []StepResult
}

type StepResult struct {
	StepNumber int
	StepType   string
	Status     string
	Message    string
	Duration   time.Duration
}
