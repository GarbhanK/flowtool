package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func ReadSQL(fileName string) string {
	// give warning if the filename doesn't have a '.sql' suffix
	if !strings.HasSuffix(fileName, ".sql") {
		color.Yellow("WARNING - specified filename does not have '.sql' suffix!\n")
	}

	f, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("error reading file contents from %s", fileName)
		os.Exit(1)
	}
	fileString := string(f)
	return fileString
}

func ValidateSQL(sqlFile string) {
	formattedSQL := strings.ReplaceAll(sqlFile, "\n", " ")
	queryWords := strings.Split(formattedSQL, " ")

	var statementType string = queryWords[0]
	if strings.ToLower(statementType) == "create" {
		color.Yellow("WARNING - copied query is a CREATE statement!\n")
	}
}

// TODO: rename something like populateThing or mapConfigValues
func TemplateSQLFile(fileName string, mapping map[string]string) string {
	sqlFile := ReadSQL(fileName)
	sqlFilePtr := &sqlFile

	var template string
	for k, v := range mapping {
		template = fmt.Sprintf("{{ %s }}", k)
		*sqlFilePtr = strings.ReplaceAll(*sqlFilePtr, template, v)
	}

	return sqlFile
}
