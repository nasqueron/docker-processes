package main

import (
	"context"
	"devcentral.nasqueron.org/source/docker-processes/pkg/dockerutils"
	"devcentral.nasqueron.org/source/docker-processes/pkg/process"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"log"
	"os"
)

func main() {
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Println("Can't connect to Docker engine.")
		os.Exit(1)
	}

	containers, err := dockerClient.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", getProcessLineTitle())

	psArgs := []string{"auxw"}
	for _, c := range containers {
		name := dockerutils.GetContainerName(c)

		response, err := dockerClient.ContainerTop(context.Background(), c.ID, psArgs)
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
