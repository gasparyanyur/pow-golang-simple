package util

import (
	"log"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "[INFO] ", log.LstdFlags)
}

func Info(message string) {
	logger.Println(message)
}

func Error(message string) {
	logger.SetPrefix("[ERROR] ")
	logger.Println(message)
	logger.SetPrefix("[INFO] ")
}

func Debug(message string) {
	logger.SetPrefix("[DEBUG] ")
	logger.Println(message)
	logger.SetPrefix("[INFO] ")
}
