package dockerutils

import (
	"github.com/docker/docker/api/types/container"
	"testing"
)

func TestGetContainerName(t *testing.T) {
	tests := []struct {
		name      string
		container container.Summary
		expected  string
	}{
		{
			name: "single name with leading slash",
			container: container.Summary{
				ID:    "1234567890abcdef",
				Names: []string{"/my-container"},
			},
			expected: "my-container",
		},
		{
			name: "no names returns truncated ID",
			container: container.Summary{
				ID:    "1234567890abcdef",
				Names: []string{},
			},
			expected: "1234567890",
		},
		{
			name: "linked container with link name first",
			container: container.Summary{
				ID:    "1234567890abcdef",
				Names: []string{"/parent/child-link", "/child-container"},
			},
			expected: "child-container",
		},
		{
			name: "only linked name available",
			container: container.Summary{
				ID:    "1234567890abcdef",
				Names: []string{"/parent/child-link"},
			},
			expected: "parent/child-link",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetContainerName(tt.container)
			if got != tt.expected {
				t.Errorf("GetContainerName() = %v, want %v", got, tt.expected)
			}
		})
	}
}
