package geo

import (
	"database/sql"
	_ "github.com/bmizerany/pq"
)

// @return [*SQLMapper]. An instantiated SQLMapper struct with the DefaultSQLConf.
// @return [Error]. Any error that might have occured during instantiating the SQLMapper.
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
