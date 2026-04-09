package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
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

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutting down")
}
