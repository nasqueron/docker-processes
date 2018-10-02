package main

import (
	"bufio"
	"context"
	"devcentral.nasqueron.org/source/docker-processes/pkg/dockerutils"
	"devcentral.nasqueron.org/source/docker-processes/pkg/process"
	"flag"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"log"
	"os"
	"strconv"
	"strings"
)

const DockerApiVersion = "1.37"

type Config struct {
	WithPosition bool
	Position int
}

func main() {
	config := parseArguments()
	scanner := bufio.NewScanner(os.Stdin)
	processesMap := getProcessesMap()

	for scanner.Scan() {
		line := addContainerName(scanner.Text(), processesMap, config)
		fmt.Println(line)
	}
}

func parseArguments() Config {
	config := Config{}

	positionPtr := flag.Int("p", -1, "the position of the field with the PID")

	flag.Parse()

	if *positionPtr > 0 {
		config.Position = *positionPtr
		config.WithPosition = true
	}

	return config
}

func addContainerName (line string, processesMap map[int64]string, config Config) string {
	fields := strings.Fields(line)

	for i, field := range fields {
		if !isValidFieldPosition(i + 1, config) {
			continue
		}

		pidCandidate, err := strconv.ParseInt(field, 10, 64)

		if err != nil {
			continue
		}

		if containerName, ok := processesMap[pidCandidate]; ok {
			return fmt.Sprintf("%s %s", line, containerName)
		}
	}

	return line
}

func isValidFieldPosition(position int, config Config) bool {
	if config.WithPosition {
		return position == config.Position
	} else {
		// 1 for top and ps -ef
		// 2 for ps auxw
		return position < 3
	}
}

func getProcessesMap () map[int64]string {
	dockerClient, err := client.NewClientWithOpts(client.WithVersion(DockerApiVersion))
	if err != nil {
		log.Println("Can't connect to Docker engine.")
		os.Exit(1)
	}

	containers, err := dockerClient.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	processesMap := make(map[int64]string)

	psArgs := []string{"auxw"}
	for _, container := range containers {
		containerName := dockerutils.GetContainerName(container)

		response, err := dockerClient.ContainerTop(context.Background(), container.ID, psArgs)
		if err != nil {
			continue
		}
		for _, containerProcess := range response.Processes {
			processInfo := process.Parse(response.Titles, containerProcess)
			processesMap[processInfo.Pid] = containerName
		}
	}

	return processesMap
}
