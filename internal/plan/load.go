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
returned to CLI
*/
