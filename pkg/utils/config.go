package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadConfig() map[string]string {
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

func AddConfig(m map[string]string, key string, val string) error {

	// check if key already exists
	for existing_key := range m {
		if key == existing_key {
			fmt.Printf("\nKey %s already exists in the config file, exiting...\n", key)
			os.Exit(1)
		}
	}

	// add the new k/v pair
	m[key] = val

	err := writeToConfig(m)
	if err != nil {
		return fmt.Errorf("error writing to 'config.json': %v", err)
	}

	return nil
}

func ListConfig() {
	// grab json config as map
	m := ReadConfig()

	// find the longest key
	var longestKey int = 0
	for key := range m {
		if len(key) > longestKey {
			longestKey = len(key)
		}
	}

	var keylen int = 0
	for key, val := range m {
		keylen = len(key)
		indent := longestKey - keylen
		whitespace := strings.Repeat(" ", indent)
		fmt.Printf("%s: %s%s\n", key, whitespace, val)
	}
}

func RemoveConfig(m map[string]string, key string) error {
	delete(m, key)
	err := writeToConfig(m)
	if err != nil {
		return fmt.Errorf("error writing to config.json: %v", err)
	}

	return nil
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
