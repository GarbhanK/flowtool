package utils

import (
	"os"
	"testing"
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
