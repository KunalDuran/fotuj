package main

import (
	"fmt"
	"time"
)

const (
	CMD_PROCESS_IMG  = "1"
	CMD_LIST_BUCKETS = "2"
)

type config struct {
	Client    string
	Vendor    string
	ServerURI string
}

var initialPrompt = `Select action: Type number and press enter: 
1. Process and upload Images
2. List Projects
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
	case CMD_LIST_BUCKETS:
		ListBuckets("vendor1")
	}
}
