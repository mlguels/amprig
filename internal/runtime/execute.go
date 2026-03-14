package runtime

func RegisterExecutors() map[string]Executor {
	return map[string]Executor{
		"set_voltage": SetVoltageExecutor{},
		"wait":        WaitExecutor{},
	}
}
