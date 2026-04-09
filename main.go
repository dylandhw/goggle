package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// logic

	log.Println("started daemon")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutting down")
}
