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

	if err := runtime.ExecutePlan(p); err != nil {
		fmt.Fprintln(os.Stderr, "execution error:", err)
		os.Exit(1)
	}
	fmt.Println("Execution completed successfully")
}
