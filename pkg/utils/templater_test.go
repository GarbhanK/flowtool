package utils

import (
	"os"
	"testing"
	// "testing/fstest"
)

var sql string = `
select *
  from ` + "`{{ params.project }}.transactions.coffee`" + ` c
  left join` + "`{{ params.web_project }}.user_data.signup`" + ` t
    on c.userId = t.userId
 where date(insertionTimestamp) >= '{{ ds_nodash }}'
 group by insertionTimestamp desc
`

// func TestReadSQL(t *testing.T) {

// 	sql := "select *\n"
// 	sql += "  from `{{ params.project }}.transactions.coffee` c\n"
// 	sql += "  left join `{{ params.web_project }}.user_data.signup` t\n"
// 	sql += "    on c.userId = t.userId\n"
// 	sql += " where date(insertionTimestamp) >= '{{ ds_nodash }}'\n"
// 	sql += " group by insertionTimestamp desc\n"

// 	fs := fstest.MapFS{
// 		"test.sql": {Data: []byte(sql)},
// 	}

// 	result := ReadSQL(fs)
// 	expected := sql

// if len(result) > 0 {
// 	t.Logf("ReadSQL('test.sql') PASSED. Is not an empty string\n")
// } else {
// 	t.Errorf("ReadSQL('test.sql') FAILED. Got an empty string\n")
// }

// if result != expected {
// 	t.Errorf("ReadSQL('test.sql') FAILED. Expected...\n%s, got..\n%s\n", expected, result)
// } else {
// 	t.Logf("ReadSQL('test.sql') PASSED. Expected... \n%s, got... \n%s\n", expected, result)
// }
// }

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
		t.Logf("ReadSQL('test.sql') PASSED. Is not an empty string\n")
	} else {
		t.Errorf("ReadSQL('test.sql') FAILED. Got an empty string\n")
	}

	// Check if the content matches what was written
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
	if result != expected {
		t.Errorf("ReadSQL(%s) FAILED. Expected...\n%s, got..\n%s\n", tmpFile.Name(), expected, result)
	} else {
		t.Logf("ReadSQL(%s) PASSED. Expected... \n%s, got... \n%s\n", tmpFile.Name(), expected, result)
	}

}

// func TestReadSQLTerraform(t *testing.T) {
// 	result := ReadSQL("terraform_test.sql")

// 	sql := "select * from `${params.project}.transactions.coffee` c\n"
// 	sql += "where date(insertionTimestamp) >= '${ds_nodash}'\n"
// 	sql += "left join `${params.web_project}.user_data.signup` t\n"
// 	sql += "on c.userId = t.userId\n"
// 	sql += "group by insertionTimestamp desc"
// 	expected := sql

// 	if len(result) > 0 {
// 		t.Logf("ReadSQL('terraform.sql') PASSED. Is not an empty string\n")
// 	} else {
// 		t.Errorf("ReadSQL('terraform_test.sql') FAILED. Got an empty string\n")
// 	}

// 	if result != expected {
// 		t.Errorf("ReadSQL('test.sql') FAILED. Expected %s, got %s\n", expected, result)
// 	} else {
// 		t.Logf("ReadSQL('test.sql') PASSED. Expected %s, got %s\n", expected, result)
// 	}
// }

// func TestValidateSQL(t *testing.T) {

// 	err := validateSQL(sqlFile)

// 	if err != nil {
// 		t.Errorf("ValidateSQL('test.sql') FAILED")
// 	} else {
// 		t.Logf("ValidateSQL for 'test.sql' returns %b", result)
// 	}
// }
