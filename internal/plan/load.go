package plan

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func Load(path string) (*Plan, error) {
	data, err := os.ReadFile(path) // reads the YAML file from disk, should return []byte which is a raw file data

	if err != nil {
		return nil, fmt.Errorf("read plan file %q: %w", path, err)
	}

	var p Plan
	// this converts YAML -> to Go struct
	if err := yaml.Unmarshal(data, &p); err != nil {
		return nil, fmt.Errorf("unmarshal yaml %q: %w", path, err)
	}
	// Prevents invalid test plans
	// check if name is filled
	if p.Name == "" {
		return nil, fmt.Errorf("plan %q: name is required", path)
	}
	// the steps length must be > 0
	if len(p.Steps) == 0 {
		return nil, fmt.Errorf("plan %q: steps must not be empty", path)
	}
	// We return a pointer because it avoids copying large structs, it also allows modification later
	return &p, nil
}

/*
smoke.yaml
    ↓
os.ReadFile
    ↓
yaml.Unmarshal
    ↓
Plan struct
    ↓
validation
    ↓
returned to CLI
*/
