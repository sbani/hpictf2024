package main

import (
	"bufio"
	"os"
	"strings"
)

// Config stores the configuration values read from the .env file.
type Config struct {
	data map[string]string
}

// NewConfig creates a new instance of Config.
func NewConfig() *Config {
	return &Config{
		data: make(map[string]string),
	}
}

// LoadConfig reads the .env file and parses it into the Config struct.
func (c *Config) LoadConfig(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			c.data[key] = value
		}
	}

	return scanner.Err()
}

// Get retrieves the value for a given key from the config.
func (c *Config) Get(key string) string {
	return c.data[key]
}
