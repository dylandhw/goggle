package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var PHONE_NUMBER = os.Getenv("PHONE_NUMBER")

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

func sendSMS() string {

	values := url.Values{
		"phone":   {string(PHONE_NUMBER)},
		"message": {"In a meeting - don't come in!"},
		"key":     {"textbelt"},
	}

	resp, err := http.PostForm("https://textbelt.com/text", values)
	if err != nil {
		log.Println("request err: ", err)
		return "failed"
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	log.Println("response:", string(body))
	return "SENT MESSAGE"
}

func main() {
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
					fmt.Println(sendSMS())
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
