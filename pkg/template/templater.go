package template

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/viper"
)

type Templater struct {
	Filename      string
	FileContents  string
	FileTemplated string
	Mapping       map[string]string
}

func CreateMapping(env string) map[string]string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	mappingFilePath := fmt.Sprintf("%s/Documents/flowtool/config.json", homeDir)
	mappingFile, err := os.ReadFile(mappingFilePath)
	if err != nil {
		fmt.Printf("Cannot find 'mappings.json' file in path %s, %s\n", mappingFilePath, err.Error())
		os.Exit(0)
	}

	// read json file into a map[string]string
	m := map[string]string{}
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

func (t *Templater) ReadSQL() {
	// give warning if the filename doesn't have a '.sql' suffix
	if !strings.HasSuffix(t.Filename, ".sql") {
		color.Yellow("WARNING - specified filename does not have '.sql' suffix!\n")
	}

	f, err := os.Open(t.Filename)
	if err != nil {
		fmt.Printf("error reading file into memory: %s", err)
		os.Exit(1)
	}
	defer f.Close()

	// read file into memory
	bytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("error reading file contents from %s", t.Filename)
		os.Exit(1)
	}

	t.FileContents = string(bytes)
}

func (t Templater) ValidateSQL() {
	formattedSQL := strings.ReplaceAll(t.FileTemplated, "\n", " ")
	queryWords := strings.Split(formattedSQL, " ")

	// Print warning to screen if it contains DDL
	statementType := strings.ToLower(queryWords[0])
	switch statementType {
	case "create":
		color.Yellow("WARNING - copied query is a CREATE statement!\n")
	case "insert":
		color.Yellow("WARNING - copied query is an INSERT statement!\n")
	case "update":
		color.Yellow("WARNING - copied query is an UPDATE statement!\n")
	case "delete":
		color.Red("WARNING - copied query is a DELETE statement!\n")
	case "drop":
		color.Red("DANGER - copied query is a DROP statement! Proceed with caution.\n")
	}
}

func (t *Templater) TemplateSQLFile() {
	t.ReadSQL()

	sqlFilePtr := &t.FileContents

	var template string
	for k, v := range t.Mapping {
		template = fmt.Sprintf("{{ %s }}", k)
		*sqlFilePtr = strings.ReplaceAll(*sqlFilePtr, template, v)
	}

	t.FileTemplated = t.FileContents
}

func (t *Templater) AddAirflowVars() {
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

	ts_nodash := fmt.Sprintf("%d%02d%dT%02d%02d%02d+0000",
		dt.Year(), dt.Month(), dt.Day(),
		dt.Hour(), dt.Minute(), dt.Second())

	ts_nodash_with_tz := fmt.Sprintf("%d%02d%dT%02d%02d%02d",
		dt.Year(), dt.Month(), dt.Day(),
		dt.Hour(), dt.Minute(), dt.Second())

	yesterday_ds := fmt.Sprintf("%d-%02d-%02d", yesterday.Year(), yesterday.Month(), yesterday.Day())
	yesterday_ds_nodash := fmt.Sprintf("%d%02d%02d", yesterday.Year(), yesterday.Month(), yesterday.Day())

	tomorrow_ds := fmt.Sprintf("%d-%02d-%02d", tomorrow.Year(), tomorrow.Month(), tomorrow.Day())
	tomorrow_ds_nodash := fmt.Sprintf("%d%02d%02d", tomorrow.Year(), tomorrow.Month(), tomorrow.Day())

	// set airflow variables in the template mapping
	t.Mapping["ds"] = ds
	t.Mapping["ds_nodash"] = ds_nodash

	t.Mapping["ts"] = ts
	t.Mapping["ts_nodash"] = ts_nodash
	t.Mapping["ts_nodash_with_tz"] = ts_nodash_with_tz

	t.Mapping["yesterday_ds"] = yesterday_ds
	t.Mapping["yesterday_ds_nodash"] = yesterday_ds_nodash
	t.Mapping["tomorrow_ds"] = tomorrow_ds
	t.Mapping["tomorrow_ds_nodash"] = tomorrow_ds_nodash
}
