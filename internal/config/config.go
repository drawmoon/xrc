package config

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
)

// Config represents the application configuration.
type Config struct {
	Verbose         bool
	SubscriptionUrl string
	SocksPort       int
	HttpPort        int
	Kernel          string
}

// LoadConfig loads the configuration from the config file.
// If the config file does not exist, it returns a default configuration.
func LoadConfig() (*Config, error) {
	ap, err := makeAppWorkingDir()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(ap)
	if err != nil {
		// If the config file does not exist, return default config
		if os.IsNotExist(err) {
			c := &Config{Verbose: false, SocksPort: 7897, HttpPort: 7897, Kernel: "xray"}
			// Save the default config
			err := SaveConfig(c)
			if err != nil {
				return nil, err
			}
			return c, nil
		}
		return nil, err
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.New("error reading app config file")
	}

	var c *Config
	err = json.Unmarshal(b, &c)
	if err != nil {
		return nil, errors.New("error parsing app config file")
	}

	return c, nil
}

// SaveConfig saves the configuration to the config file.
// It creates the working directory if it does not exist.
func SaveConfig(c *Config) error {
	ap, err := makeAppWorkingDir()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return errors.New("error serializing config")
	}
	err = os.WriteFile(ap, data, 0644)
	if err != nil {
		return errors.New("error writing config file")
	}
	return nil
}

// makeAppWorkingDir ensures the application working directory exists
// and returns the path to the config file.
func makeAppWorkingDir() (string, error) {
	workDirName := ".xrc"
	userHome, _ := os.UserHomeDir()
	workDir := filepath.Join(userHome, workDirName)

	// Ensure the working directory exists
	err := os.MkdirAll(workDir, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return "", err
	}

	return filepath.Join(workDir, "config.json"), nil
}
