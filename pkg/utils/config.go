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

	// add the new k/v pair
	m[key] = val

	jsonString, _ := json.Marshal(m)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	fp := fmt.Sprintf("%s/Documents/flowtool/config.json", homeDir)

	err = ioutil.WriteFile(fp, jsonString, os.ModePerm)
	if err != nil {
		log.Fatal(err)
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
	fmt.Print("\n")

	fmt.Println("	==== config.json contents ====")
	for key, val := range m {
		// fmt.Println(key, ":", indent, val)
		fmt.Println(key, ":", val)
	}
}
