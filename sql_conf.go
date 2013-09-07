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

// TODO potentially package into file included with the package
var defaultOpenStr = "user=golang_geo_test dbname=points sslmode=disable"
var dbEnv = os.Getenv("DB")
var DefaultSQLConf = &SQLConf{}

// Attempts to read config/geo.yml, and creates a {SQLConf} as described in the file
// Returns the DefaultSQLConf if no config/geo.yml is found.
// @return [*SQLConf].  The SQLConfiguration, as supplied with config/geo.yml
// @return [Error].  Any error that might occur while grabbing configuration
func GetSQLConf() (*SQLConf, error) {
	if dbEnv != "" {
		DefaultSQLConf = &SQLConf{driver: dbEnv, openStr: defaultOpenStr, table: "points", latCol: "lat", lngCol: "lng"}
	} else {
		DefaultSQLConf = &SQLConf{driver: "postgres", openStr: defaultOpenStr, table: "points", latCol: "lat", lngCol: "lng"}
	}

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
