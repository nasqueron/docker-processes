package main

import (
	"context"
	"devcentral.nasqueron.org/source/docker-processes/pkg/process"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"log"
	"os"
	"strings"
)

const DockerApiVersion = "1.37"

func main() {
	dockerClient, err := client.NewClientWithOpts(client.WithVersion(DockerApiVersion))
	if err != nil {
		log.Println("Can't connect to Docker engine.")
		os.Exit(1)
	}

	containers, err := dockerClient.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", getProcessLineTitle())

	psArgs := []string{"auxw"}
	for _, container := range containers {
		name := getContainerName(container)

		response, err := dockerClient.ContainerTop(context.Background(), container.ID, psArgs)
		if err != nil {
			continue
		}
		for _, containerProcess := range response.Processes {
			processInfo := process.Parse(response.Titles, containerProcess)
			processLine := getProcessLine(name, processInfo)
			fmt.Printf("%s\n", processLine)
		}
	}
}

func getProcessLineTitle() string {
	return "CONTAINER            USER         PID %CPU      VSZ      RSS COMMAND"
}

func getProcessLineFormat() string {
	return "%20s %9s %6d %.2f %8d %8d %s"
}

func getProcessLine(containerName string, processInfo process.Process) string {
	format := getProcessLineFormat()

	return fmt.Sprintf(
		format,
		containerName, processInfo.User, processInfo.Pid, processInfo.CPU,
		processInfo.VSZ, processInfo.RSS, processInfo.Command)
}

func getContainerName(container types.Container) string {
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
