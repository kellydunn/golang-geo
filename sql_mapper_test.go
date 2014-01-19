package geo

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
)

// Ensures that
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
