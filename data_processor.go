package main

import (
	"config"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type DataProcessor struct {
	config *config.Config
}

func NewDataProcessor(cfg *config.Config) *DataProcessor {
	return &DataProcessor{
		config: cfg,
	}
}

func (dp *DataProcessor) ProcessData() {
	processingInterval := dp.config.Get("PROCESSING_INTERVAL")
	interval, err := time.ParseDuration(processingInterval)
	if err != nil {
		fmt.Printf("Invalid processing interval: %s\n", err)
		return
	}

	logFilePath := dp.config.Get("LOG_FILE_PATH")
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening log file: %s\n", err)
		return
	}
	defer logFile.Close()

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		dataItem := dp.generateData()
		logMessage := fmt.Sprintf("Processed data: %s at %s\n", dataItem, time.Now().Format(time.RFC1123))
		if _, err := logFile.WriteString(logMessage); err != nil {
			fmt.Printf("Error writing to log file: %s\n", err)
			continue
		}
		fmt.Println(logMessage)
	}
}

func (dp *DataProcessor) generateData() string {
	// Simulate data generation with random numbers
	randomData := rand.Intn(100)
	dataInfo := fmt.Sprintf("data_value_%d", randomData)
	return dataInfo
}

func main() {
	cfg := config.NewConfig()
	err := cfg.LoadConfig(".env")
	if err != nil {
		fmt.Printf("Error loading configuration: %s\n", err)
		os.Exit(1)
	}

	dataProcessor := NewDataProcessor(cfg)
	dataProcessor.ProcessData()
}
