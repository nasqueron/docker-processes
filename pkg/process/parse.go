package process

import "strconv"

func ParseIntOrZero (data string) int64 {
	value, err := strconv.ParseInt(data, 10, 64)

	if err != nil {
		return 0
	}

	return value
}

func ParseFloatOrZero (data string) float64 {
	value, err := strconv.ParseFloat(data, 64)

	if err != nil {
		return 0
	}

	return value
}

