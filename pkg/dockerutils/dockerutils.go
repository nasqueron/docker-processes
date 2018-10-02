package dockerutils

import (
	"github.com/docker/docker/api/types"
	"strings"
)

func GetContainerName(container types.Container) string {
	names := container.Names

	if len(names) == 0 {
		return container.ID[:10]
	}

	bestCandidate := names[0][1:]

	// Linked containers offer link names before the container name.
	if strings.Contains(bestCandidate, "/") {
		for _, name := range names {
			if !strings.Contains(name[1:], "/") {
				return name[1:]
			}
		}
	}

	return bestCandidate
}
