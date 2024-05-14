package main

import (
	"config"
	"fmt"
	"os"
	"time"
)

func main() {
	cfg := config.NewConfig()
	err := cfg.LoadConfig(".env")
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		os.Exit(1)
	}

	// Example usage of the configuration
	performRandomTask(cfg)
}

// performRandomTask simulates a random task using the loaded config.
func performRandomTask(cfg *config.Config) {
	// Assume we have a key "TASK_FREQUENCY" that defines how often a task should run.
	frequency := cfg.Get("TASK_FREQUENCY")
	fmt.Println("Task will run with frequency:", frequency)

	// Simulation of a task that logs a timestamp every few seconds based on TASK_FREQUENCY
	tickDuration, err := time.ParseDuration(frequency)
	if err != nil {
		fmt.Println("Invalid duration format:", err)
		return
	}
	ticker := time.NewTicker(tickDuration)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Println("Task executed at:", time.Now().Format(time.RFC1123))
	}
}
