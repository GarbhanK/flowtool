package utils

import (
	"bytes"
	"os"
	"testing"

	"github.com/fatih/color"
)

var sql string = `select *
  from ` + "`{{ params.project }}.transactions.coffee`" + ` c
  left join` + "`{{ params.web_project }}.user_data.signup`" + ` t
    on c.userId = t.userId
 where date(insertionTimestamp) >= '{{ ds_nodash }}'
 group by insertionTimestamp desc
`

func TestReadSQL(t *testing.T) {
	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "test.sql")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Write some content to the temporary file
	if _, err := tmpFile.WriteString(sql); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	expected := sql
	result, err := ReadSQL(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if len(result) > 0 {
		t.Logf("ReadSQL(%s) PASSED. Is not an empty string\n", tmpFile.Name())
	} else {
		t.Errorf("ReadSQL(%s) FAILED. Got an empty string\n", tmpFile.Name())
	}

	// Check if the content matches what was written
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
	if result != expected {
		t.Errorf("ReadSQL(%s) FAILED.\nExpected...\n%sGot..\n%s\n", tmpFile.Name(), expected, result)
	} else {
		t.Logf("ReadSQL(%s) PASSED.\nExpected... \n%sGot... \n%s\n", tmpFile.Name(), expected, result)
	}

}

func TestValidateSQL(t *testing.T) {
	// Capture the output of the function
	var buf bytes.Buffer
	color.Output = &buf

	tests := []struct {
		name           string
		sqlFile        string
		expectedOutput string
	}{
		{
			name:           "Create statement",
			sqlFile:        "CREATE TABLE users (id INT)",
			expectedOutput: "WARNING - copied query is a CREATE statement!",
		},
		{
			name:           "Insert statement",
			sqlFile:        "INSERT INTO users (id, name) VALUES (1, 'John')",
			expectedOutput: "WARNING - copied query is an INSERT statement!",
		},
		{
			name:           "Update statement",
			sqlFile:        "UPDATE users SET name = 'Jane' WHERE id = 1",
			expectedOutput: "WARNING - copied query is an UPDATE statement!",
		},
		{
			name:           "Delete statement",
			sqlFile:        "DELETE FROM users WHERE id = 1",
			expectedOutput: "WARNING - copied query is a DELETE statement!",
		},
		{
			name:           "Drop statement",
			sqlFile:        "DROP TABLE users",
			expectedOutput: "DANGER - copied query is a DROP statement! Proceed with caution.",
		},
		{
			name:           "Safe query",
			sqlFile:        "SELECT * FROM users",
			expectedOutput: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Reset the output buffer
			buf.Reset()

			// Call the function with the test case's SQL input
			ValidateSQL(test.sqlFile)

			// Get the output and remove any newlines/trailing spaces for easier comparison
			output := buf.String()
			if len(output) != 0 {
				output = output[:len(output)-1] // Remove the last newline character
			}

			// Check if the actual output matches the expected output
			if output != test.expectedOutput {
				t.Errorf("expected %q but got %q", test.expectedOutput, output)
			}
		})
	}
}
