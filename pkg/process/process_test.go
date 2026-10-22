package process

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	labels := []string{"USER", "PID", "%CPU", "VSZ", "RSS", "COMMAND"}
	info := []string{"root", "1234", "1.5", "10240", "2048", "/usr/bin/app"}

	expected := Process{
		User:    "root",
		Pid:     1234,
		CPU:     1.5,
		VSZ:     10240,
		RSS:     2048,
		Command: "/usr/bin/app",
	}

	got := Parse(labels, info)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Parse() = %+v, want %+v", got, expected)
	}
}

func TestParseIntOrZero(t *testing.T) {
	if got := ParseIntOrZero("123"); got != 123 {
		t.Errorf("ParseIntOrZero(\"123\") = %v, want 123", got)
	}
	if got := ParseIntOrZero("invalid"); got != 0 {
		t.Errorf("ParseIntOrZero(\"invalid\") = %v, want 0", got)
	}
}

func TestParseFloatOrZero(t *testing.T) {
	if got := ParseFloatOrZero("12.34"); got != 12.34 {
		t.Errorf("ParseFloatOrZero(\"12.34\") = %v, want 12.34", got)
	}
	if got := ParseFloatOrZero("invalid"); got != 0.0 {
		t.Errorf("ParseFloatOrZero(\"invalid\") = %v, want 0.0", got)
	}
}
