package plan

import "time"

type Plan struct {
	Name    string `yaml:"name"`
	Version int    `yaml:"version"`
	Steps   []Step `yaml:"steps"`
}

type Step struct {
	Type string `yaml:"type"`

	// set_voltage
	ValueV        float64 `yaml:"value_v"`
	CurrentLimitA float64 `yaml:"current_limit_a"`

	// wait
	Duration time.Duration `yaml:"duration"`

	// measure
	Metrics []string `yaml:"metrics"`

	// assert
	Metric string  `yaml:"metric"`
	Op     string  `yaml:"op"`
	Value  float64 `yaml:"value"`
}
