package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	// "strings"
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
		os.Exit(0)
	}

	m := map[string]string{}

	// read json file into a map[string]string
	json.Unmarshal([]byte(configFile), &m)

	return m
}

func AddConfig(m map[string]string, key string, val string) {
	// TODO: add check if key already exists

	for existing_key, _ := range m {
		if key == existing_key {
			fmt.Printf("\nKey %s already exists in the config file, exiting...\n", key)
			os.Exit(1)
		}
	}

	// add the new k/v pair
	m[key] = val

	err := writeToConfig(m)
	if err != nil {
		fmt.Printf("Error writing to config.json: %w\n", err)
	}
}

func ListConfig() {
	// grab json config as map
	m := ReadConfig()

	// find the longest key
	var longestKey int = 0
	for key, _ := range m {
		if len(key) > longestKey {
			longestKey = len(key)
		}
	}

	// indent := strings.Repeat(string(' '), longestKey)
	for key, val := range m {
		// fmt.Println(key, ":", indent, val)
		fmt.Println(key, ":", val)
	}
}

func RemoveConfig(m map[string]string, key string) {
	delete(m, key)
	err := writeToConfig(m)
	if err != nil {
		log.Fatal(err)
	}
}

func writeToConfig(m map[string]string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Error finding Home Directory: %w", err)
	}

	fp := fmt.Sprintf("%s/Documents/flowtool/config.json", homeDir)

	jsonString, _ := json.MarshalIndent(m, "", "    ")
	err = ioutil.WriteFile(fp, jsonString, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Error writing to config file: %w", err)
	}

	return nil
}
