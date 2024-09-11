package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Config struct {
	Filename string
	Contents map[string]string
}

func NewConfig() Config {
	newConfig := Config{
		Filename: "config.json",
		Contents: readConfig(),
	}
	return newConfig
}

func readConfig() map[string]string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	fp := fmt.Sprintf("%s/Documents/flowtool/config.json", homeDir)
	configFile, err := os.ReadFile(fp)
	if err != nil {
		fmt.Printf("Cannot find 'config.json' file in path %s, %s\n", fp, err.Error())
		os.Exit(1)
	}

	m := map[string]string{}

	// read json file into a map[string]string
	json.Unmarshal([]byte(configFile), &m)

	return m
}

func writeToConfig(m map[string]string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error finding Home Directory: %v", err)
	}

	fp := fmt.Sprintf("%s/Documents/flowtool/config.json", homeDir)

	jsonString, _ := json.MarshalIndent(m, "", "    ")
	err = os.WriteFile(fp, jsonString, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error writing to config file: %w", err)
	}

	return nil
}

func (c *Config) Add(key string, val string) error {

	// check if key already exists
	for existing_key := range c.Contents {
		if key == existing_key {
			fmt.Printf("\nKey %s already exists in the config file, exiting...\n", key)
			os.Exit(1)
		}
	}

	// add the new k/v pair
	c.Contents[key] = val

	err := writeToConfig(c.Contents)
	if err != nil {
		return fmt.Errorf("error writing to 'config.json': %v", err)
	}

	return nil
}

func (c Config) List() {
	// find the longest key
	var longestKey int = 0
	for key := range c.Contents {
		if len(key) > longestKey {
			longestKey = len(key)
		}
	}

	var keylen int = 0
	for key, val := range c.Contents {
		keylen = len(key)
		indent := longestKey - keylen
		whitespace := strings.Repeat(" ", indent)
		fmt.Printf("%s: %s%s\n", key, whitespace, val)
	}
}

func (c *Config) Remove(key string) error {
	delete(c.Contents, key)
	err := writeToConfig(c.Contents)
	if err != nil {
		return fmt.Errorf("error writing to config.json: %v", err)
	}

	return nil
}
