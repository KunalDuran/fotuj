package main

import (
	"fmt"
	"time"
)

const (
	CMD_PROCESS_IMG   = "1"
	CMD_LIST_PROJECTS = "2"
)

var initialPrompt = `Select action: Type number and press enter: 
1. Create new project
2. View projects
`

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	command := prompt(initialPrompt)
	switch command {
	case CMD_PROCESS_IMG:
		ProcessImages()
	case CMD_LIST_PROJECTS:
		ShowProjects()
	}
}
