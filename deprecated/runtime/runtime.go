package runtime

type Runtime struct {
	currentVoltageV float64
	currentLimitA   float64
	measurements    map[string]float64
}

func NewRuntime() *Runtime {
	return &Runtime{
		measurements: make(map[string]float64),
	}
}
