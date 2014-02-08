package geo

import (
	"fmt"
	"os"
	"testing"
)

// Ensures that getting a SQL Configuration from a non existent YAML file will
// return a SQLConf struct that is holding the expeted data.
func TestGetSQLConfFromFile(t *testing.T) {
	env := os.Getenv("DB")
	path := fmt.Sprintf("db/%s/dbconf.yml", env)
	conf, err := GetSQLConfFromFile(path)

	if err != nil {
		fmt.Printf("%v\n", err)
		t.Error("Did not expect for an error when supplying an existing configuration file")
	}

	expected := sqlConfFromEnv()

	if conf.openStr != expected.openStr {
		t.Error("Expected the SQL configuration to match the expected SQL configuration.")
	}
}

// Ensures that getting a SQL Configuration from a non existent YAML file will
// return a Default Sql Configuration
func TestGetSQLConfFromFileNonExistent(t *testing.T) {
	conf, err := GetSQLConfFromFile("garbage")
	expected := sqlConfFromEnv()

	// TODO We should probably alert an error to the fact that we can't find the file
	//      But this would introduce a backwards compatible change since
	//      the GetSQLConf() code path
	//      Until then, we expect to just return the default sql conf and nothing else.
	if err != nil {
		t.Error("Did not expect and error when supplying a non-existent configuration file.")
	}

	if conf.openStr != expected.openStr {
		t.Error("Expected the SQL configuration to match the detault SQL configuration for the current environment.")
	}
}
