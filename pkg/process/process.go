package process

type Process struct {
	User    string
	Pid     int64
	CPU     float64
	VSZ     int64
	RSS     int64
	Command string
}

// Parses from two string slices given by Docker API
func Parse(processLabels []string, processInfo []string) Process {
	process := Process{}

	for i, processInfoPart := range processInfo {
		switch processLabels[i] {
		case "USER":
			process.User = processInfoPart

		case "PID":
			process.Pid = ParseIntOrZero(processInfoPart)

		case "%CPU":
			process.CPU = ParseFloatOrZero(processInfoPart)

		case "VSZ":
			process.VSZ = ParseIntOrZero(processInfoPart)

		case "RSS":
			process.RSS = ParseIntOrZero(processInfoPart)

		case "COMMAND":
			process.Command = processInfoPart
		}
	}

	return process
}
