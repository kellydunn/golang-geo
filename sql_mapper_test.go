package geo

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
)

// Ensures that creating a new SQL Mapper does not encounter an error upon initialization
// And also matches the expected database connections and sql configurations.
func TestNewSQLMapper(t *testing.T) {
	conf := sqlConfFromEnv()
	db, _ := sql.Open(conf.driver, conf.openStr)

	env := os.Getenv("DB")
	filepath := fmt.Sprintf("db/%s/dbconf.yml", env)
	s, _ := NewSQLMapper(filepath, db)

	if s == nil {
		t.Error("Expected NewSqlMapper to return a non-nil pointer to a sql mapper")
	}

	if s.sqlConn != db {
		t.Error()
	}
}

// Ensures that creating a new SQLMapper and getting its SQL DB Connection
// is the same *sql.DB used in initialization.
func TestSqlDbConn(t *testing.T) {
	conf := sqlConfFromEnv()
	db, _ := sql.Open(conf.driver, conf.openStr)

	env := os.Getenv("DB")
	filepath := fmt.Sprintf("db/%s/dbconf.yml", env)
	s, _ := NewSQLMapper(filepath, db)

	if s.SqlDbConn() != db {
		t.Error("Expected db connections are mismatched.")
	}
}
