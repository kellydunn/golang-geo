package geo

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
	"os"
	"path"
)

// Provides the configuration to query the database as necessary
type SQLConf struct {
	driver  string
	openStr string
	table   string
	latCol  string
	lngCol  string
}

const (
	DEFAULT_PGSQL_OPEN_STR = "user=postgres dbname=points sslmode=disable"
	DEFAULT_MYSQL_OPEN_STR = "points/root"
)

// Returns a Sql Conf
func sqlConfFromEnv() (*SQLConf) {
	var dbEnv = os.Getenv("DB")	

	switch dbEnv {
	case "mysql":
		return &SQLConf{driver: dbEnv, openStr: DEFAULT_PGSQL_OPEN_STR, table: "points", latCol: "lat", lngCol: "lng"}
	default:
		return &SQLConf{driver: "postgres", openStr: DEFAULT_PGSQL_OPEN_STR, table: "points", latCol: "lat", lngCol: "lng"}		
	}
}

// Attempts to read config/geo.yml, and creates a {SQLConf} as described in the file
// Returns the DefaultSQLConf if no config/geo.yml is found.
// @return [*SQLConf].  The SQLConfiguration, as supplied with config/geo.yml
// @return [Error].  Any error that might occur while grabbing configuration
func GetSQLConf() (*SQLConf, error) {
	DefaultSQLConf := sqlConfFromEnv()

	// TODO This should be redesigned so that the user specifies where the config file is
	//      We can still handle the issue where it doesn't exist, but that way it's not hardcoded.
	configPath := path.Join("config/geo.yml")
	_, err := os.Stat(configPath)
	if err != nil && os.IsNotExist(err) {
		return DefaultSQLConf, nil
	} else {

		// Defaults to development environment, you can override by changing the $GO_ENV variable:
		// `$ export GO_ENV=environment` (where environment can be "production", "test", "staging", etc.
		// TODO Potentially find a better solution to handling environments
		// https://github.com/adeven/goenv ?
		goEnv := os.Getenv("GO_ENV")
		if goEnv == "" {
			goEnv = "development"
		}

		config, readYamlErr := yaml.ReadFile(configPath)
		if readYamlErr == nil {

			// TODO Refactor this into a more generic method of retrieving info

			// Get driver
			driver, driveError := config.Get(fmt.Sprintf("%s.driver", goEnv))
			if driveError != nil {
				return nil, driveError
			}

			// Get openStr
			openStr, openStrError := config.Get(fmt.Sprintf("%s.openStr", goEnv))
			if openStrError != nil {
				return nil, openStrError
			}

			// Get table
			table, tableError := config.Get(fmt.Sprintf("%s.table", goEnv))
			if tableError != nil {
				return nil, tableError
			}

			// Get latCol
			latCol, latColError := config.Get(fmt.Sprintf("%s.latCol", goEnv))
			if latColError != nil {
				return nil, latColError
			}

			// Get lngCol
			lngCol, lngColError := config.Get(fmt.Sprintf("%s.lngCol", goEnv))
			if lngColError != nil {
				return nil, lngColError
			}

			sqlConf := &SQLConf{driver: driver, openStr: openStr, table: table, latCol: latCol, lngCol: lngCol}
			return sqlConf, nil

		}

		return nil, readYamlErr
	}

	return nil, err
}
