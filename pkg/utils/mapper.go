package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

func CreateMapping(env string) map[string]string {

	// var mappingFilePath string

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// if isTest {
	// 	mappingFilePath = "./test_mappings.json"
	// } else {
	// 	mappingFilePath = fmt.Sprintf("%s/Documents/flowtool/config.json", homeDir)
	// }

	mappingFilePath := fmt.Sprintf("%s/Documents/flowtool/config.json", homeDir)
	mappingFile, err := os.ReadFile(mappingFilePath)
	if err != nil {
		fmt.Printf("Cannot find 'mappings.json' file in path %s, %s\n", mappingFilePath, err.Error())
		os.Exit(0)
	}
	m := map[string]string{}

	// read json file into a map[string]string
	json.Unmarshal([]byte(mappingFile), &m)

	// env variable priority is... (--env flag > config file > default)
	// if I have a flag set (default is dev), override config file
	cfgFileEnv := viper.GetString("env")
	if env == "dev" && cfgFileEnv != "" {
		env = cfgFileEnv
	}

	// template chosen environment into our mapping file
	for k, v := range m {
		m[k] = strings.Replace(v, "${env}", env, -1)
	}

	return m
}

func AddAirflowTemplateVars(m map[string]string) map[string]string {
	// grab the current airflow date (today -1)
	dt := time.Now().AddDate(0, 0, -1)
	yesterday := dt.AddDate(0, 0, -1)
	tomorrow := dt.AddDate(0, 0, 1)

	// create airflow template variables ref: https://airflow.apache.org/docs/apache-airflow/stable/templates-ref.html
	ds := fmt.Sprintf("%d-%02d-%02d", dt.Year(), dt.Month(), dt.Day())
	ds_nodash := fmt.Sprintf("%02d%02d%02d", dt.Year(), dt.Month(), dt.Day())
	ts := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d+00:00",
		dt.Year(), dt.Month(), dt.Day(),
		dt.Hour(), dt.Minute(), dt.Second())

	yesterday_ds := fmt.Sprintf("%d-%02d-%02d", yesterday.Year(), yesterday.Month(), yesterday.Day())
	yesterday_ds_nodash := fmt.Sprintf("%d%02d%02d", yesterday.Year(), yesterday.Month(), yesterday.Day())

	tomorrow_ds := fmt.Sprintf("%d-%02d-%02d", tomorrow.Year(), tomorrow.Month(), tomorrow.Day())
	tomorrow_ds_nodash := fmt.Sprintf("%d%02d%02d", tomorrow.Year(), tomorrow.Month(), tomorrow.Day())

	m["ds"] = ds
	m["ds_nodash"] = ds_nodash
	m["ts"] = ts
	m["yesterday_ds"] = yesterday_ds
	m["yesterday_ds_nodash"] = yesterday_ds_nodash
	m["tomorrow_ds"] = tomorrow_ds
	m["tomorrow_ds_nodash"] = tomorrow_ds_nodash

	return m
}
