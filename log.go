package main

import (
	"log"
	"os"
)

var logFile *os.File

func initLogger() {
	file, err := os.OpenFile(LogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(file)

	logFile = file
}
