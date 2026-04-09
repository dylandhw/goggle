package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func isWebcamActive() bool {
	cmd := exec.Command("fuser", "/dev/video0")
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("error: ", err)
	}

	outStr := strings.TrimSpace(string(output))
	if outStr != "" {
		return true
	}
	return false
}

func main() {
	// logic
	log.Println("started daemon")

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	var wasActive bool

	for {
		select {
		case <-ticker.C:
			active := isWebcamActive()
			if active != wasActive {
				if active {
					log.Println(">>>WEBCAM IS ACTIVE<<<")
				} else {
					log.Println(">>>WEBCAM IS INACTIVE<<<")
				}
				wasActive = active
			}
		case <-quit:
			log.Println("shutting down")
			return
		}
	}
}
