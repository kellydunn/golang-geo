package geo

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
	"os"
	"path"
)

// Provides a set of configuration variables that describe how to interact with a SQL database.
type SQLConf struct {
	driver  string
	openStr string
	table   string
	latCol  string
	lngCol  string
}

const (
	DEFAULT_PGSQL_OPEN_STR = "user=postgres dbname=points sslmode=disable"
	DEFAULT_MYSQL_OPEN_STR = "points/root/"
	DEFAULT_TEST_OPEN_STR  = "\"\""
)

// Returns a SQLConf based on the $DB environment variable
// Returns a PostgreSQL configuration as a default
func sqlConfFromEnv() *SQLConf {
	var dbEnv = os.Getenv("DB")

	switch dbEnv {
	case "mysql":
		return &SQLConf{driver: "mymysql", openStr: DEFAULT_MYSQL_OPEN_STR, table: "points", latCol: "lat", lngCol: "lng"}
	case "mock":
		return &SQLConf{driver: "testdb", openStr: DEFAULT_TEST_OPEN_STR, table: "points", latCol: "lat", lngCol: "lng"}
	default:
		return &SQLConf{driver: "postgres", openStr: DEFAULT_PGSQL_OPEN_STR, table: "points", latCol: "lat", lngCol: "lng"}
	}
}

// Attempts to read config/geo.yml, and creates a SQLConf as described therein.
// Returns the DefaultSQLConf if no config/geo.yml is found, or an error
// if one arises during the process of parsing the configuration file.
func GetSQLConf() (*SQLConf, error) {
	return GetSQLConfFromFile("config/geo.yml")
}

// Attempts to read from the passed in filename and creates a SQLconf as described therin.
// Retruns the DefaultSQLConf if the file cannot be found, or an error
// if one arises during the process of parsing the configuration file.
func GetSQLConfFromFile(filename string) (*SQLConf, error) {
	DefaultSQLConf := sqlConfFromEnv()
	configPath := path.Join(filename)
	_, err := os.Stat(configPath)

	if err != nil && os.IsNotExist(err) {
		return DefaultSQLConf, nil
	} else {

		// Defaults to development environment, you can override
		// by changing the $GO_ENV variable:
		// `$ export GO_ENV=environment` (where environment
		// can be "production", "test", "staging", etc.)
		// TODO: Potentially find a better solution to handling environments
		// Perhaps: https://github.com/adeven/goenv ?
		goEnv := os.Getenv("GO_ENV")
		if goEnv == "" {
			goEnv = "development"
		}

		config, readYamlErr := yaml.ReadFile(configPath)
		if readYamlErr != nil {
			return nil, readYamlErr
		}

		return confFromYamlFile(config, goEnv)
	}

	return nil, err
}

// Creates a new SQLConf and returns a pointer to it.
// If it encounters an error during parsing the file,
// it will return an error instead.
func confFromYamlFile(config *yaml.File, goEnv string) (*SQLConf, error) {

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
