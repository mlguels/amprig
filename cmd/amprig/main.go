package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mlguels/amprig/internal/plan"
	"github.com/mlguels/amprig/internal/runtime"
)

func main() {
	planPath := flag.String("plan", "plans/smoke.yaml", "path to test plan")
	flag.Parse()

	p, err := plan.Load(*planPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "load error:", err)
		os.Exit(1)
	}

	if err := plan.Validate(p); err != nil {
		fmt.Fprintln(os.Stderr, "validation error:", err)
		os.Exit(1)
	}

	fmt.Printf("Plan: %s\n", p.Name)
	fmt.Printf("Version: %d\n", p.Version)
	fmt.Printf("Steps: %d\n", len(p.Steps))

	for i, step := range p.Steps {
		fmt.Printf("%d. %s\n", i+1, step.Type)
	}

	result, err := runtime.ExecutePlan(p)
	if err != nil {
		fmt.Fprintln(os.Stderr, "execution error:", err)
		fmt.Printf("Plan result so far: %+v\n", result)
		os.Exit(1)
	}
	fmt.Println("Execution completed successfully")

	resultStatus := "FAILED"
	if result.Success {
		resultStatus = "PASSED"
	}

	passed := 0
	for _, step := range result.Steps {
		if step.Status == "passed" {
			passed++
		}
	}

	totalDuration := result.FinishedAt.Sub(result.StartedAt)

	fmt.Printf("\nSummary\n")
	fmt.Printf("-------\n")
	fmt.Printf("Plan: %s\n", result.PlanName)
	fmt.Printf("Result: %s\n", resultStatus)
	fmt.Printf("Steps: %d/%d\n", passed, len(result.Steps))
	fmt.Printf("Total Duration: %v\n", totalDuration)
}
