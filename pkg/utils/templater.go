package utils

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
)

func ReadSQL(fileName string) (string, error) {
	// give warning if the filename doesn't have a '.sql' suffix
	if !strings.HasSuffix(fileName, ".sql") {
		color.Yellow("WARNING - specified filename does not have '.sql' suffix!\n")
	}

	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// read file into memory
	bytes, err := io.ReadAll(f)
	if err != nil {
		return "", fmt.Errorf("error reading file contents from %s", fileName)
	}

	return string(bytes), nil
}

func ValidateSQL(sqlFile string) {
	formattedSQL := strings.ReplaceAll(sqlFile, "\n", " ")
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

// TODO: rename something like populateThing or mapConfigValues
func TemplateSQLFile(fileName string, mapping map[string]string) string {
	sqlFile, err := ReadSQL(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	sqlFilePtr := &sqlFile

	var template string
	for k, v := range mapping {
		template = fmt.Sprintf("{{ %s }}", k)
		*sqlFilePtr = strings.ReplaceAll(*sqlFilePtr, template, v)
	}

	return sqlFile
}
