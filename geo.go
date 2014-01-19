package geo

import (
	"database/sql"
	_ "github.com/erikstmartin/go-testdb"
	_ "github.com/lib/pq"
	_ "github.com/ziutek/mymysql/godrv"
)

// Retrieves the SQL configuration specified in config.yml
// that resides at the root level of the project.
// Returns a pointer to a SQLMapper if successful, or an error
// if there is an issue opening a database connection.
func HandleWithSQL() (*SQLMapper, error) {
	sqlConf, sqlConfErr := GetSQLConf()
	if sqlConfErr == nil {
		s := &SQLMapper{conf: sqlConf}

		db, err := sql.Open(s.conf.driver, s.conf.openStr)
		if err != nil {
			panic(err)
		}

		s.sqlConn = db
		return s, err
	}

	return nil, sqlConfErr
}
