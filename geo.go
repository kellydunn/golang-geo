package geo

import (
	"database/sql"
	"fmt"
	_ "github.com/bmizerany/pq"
	"math"
        "github.com/kylelemons/go-gypsy/yaml"
	"path"
	"os"
)

// TODO potentially package into file included with the package
var DefaultSQLConf = &SQLConf{driver: "postgres", openStr: "user=postgres password=postgres dbname=points sslmode=disable", table: "points", latCol: "lat", lngCol: "lng"}

// Attempts to read config/geo.yml, and creates a {SQLConf} as described in the file
// Returns the DefaultSQLConf if no config/geo.yml is found.
// @return [*SQLConf].  The SQLConfiguration, as supplied with config/geo.yml
// @return [Error].  Any error that might occur while grabbing configuration
func GetSQLConf() (*SQLConf, error) {
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
				return nil,openStrError
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

			sqlConf := &SQLConf{driver: driver, openStr:openStr, table:table, latCol:latCol, lngCol:lngCol}
			return sqlConf, nil
			
		}
		
		return nil, readYamlErr
	}
	
	return nil, err
}

// Represents a Physical Point in geographic notation [lat, lng]
// @field [float64] lat. The geographic latitude representation of this point.
// @field [float64] lng. The geographic longitude representation of this point.
type Point struct {
	lat float64
	lng float64
}

// Original Implementation from: http://www.movable-type.co.uk/scripts/latlong.html
// @param [float64] dist.  The arc distance in which to transpose the origin point (in meters).
// @param [float64] bearing.  The compass bearing in which to transpose the origin point (in degrees).
// @return [*Point].  Returns a Point struct populated with the lat and lng coordinates
//                    of transposing the origin point a certain arc distance at a certain bearing.
func (p *Point) PointAtDistanceAndBearing(dist float64, bearing float64) *Point {
	// Earth's radius ~= 6,356.7523km
	// TODO Constantize
	dr := dist / 6356.7523

	bearing = (bearing * (math.Pi / 180.0))

	lat1 := (p.lat * (math.Pi / 180.0))
	lng1 := (p.lng * (math.Pi / 180.0))

	lat2_part1 := math.Sin(lat1) * math.Cos(dr)
	lat2_part2 := math.Cos(lat1) * math.Sin(dr) * math.Cos(bearing)

	lat2 := math.Asin(lat2_part1 + lat2_part2)

	lng2_part1 := math.Sin(bearing) * math.Sin(dr) * math.Cos(lat1)
	lng2_part2 := math.Cos(dr) - (math.Sin(lat1) * math.Sin(lat2))

	lng2 := lng1 + math.Atan2(lng2_part1, lng2_part2)
	lng2 = math.Mod((lng2+3*math.Pi), (2*math.Pi)) - math.Pi

	lat2 = lat2 * (180.0 / math.Pi)
	lng2 = lng2 * (180.0 / math.Pi)

	return &Point{lat: lat2, lng: lng2}
}

// Original Implementation from: http://www.movable-type.co.uk/scripts/latlong.html
// Calculates the Haversine distance between two points.
// @param [*Point].  The destination point.
// @return [float64].  The distance between the origin point and the destination point.
func (p * Point) GreatCircleDistance(p2 * Point) (float64) {
	r := 6356.7523; // km
	dLat := (p2.lat-p.lat) * (math.Pi / 180.0)
	dLon := (p2.lng-p.lng) * (math.Pi / 180.0)
	
	lat1 := p.lat * (math.Pi / 180.0)
	lat2 := p2.lat * (math.Pi / 180.0)

	a1 := math.Sin(dLat/2) * math.Sin(dLat/2) 
	a2 := math.Sin(dLon/2) * math.Sin(dLon/2) * math.Cos(lat1) * math.Cos(lat2); 
	
	a := a1 + a2
	
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a)); 
	
	return r * c;
}

// Provides a Queryable interface for finding Points via some Data Storage mechanism
type Mapper interface {
	PointsWithinRadius(p *Point, radius int) bool
}

// Provides the configuration to query the database as necessary
type SQLConf struct {
	driver  string
	openStr string
	table   string
	latCol  string
	lngCol  string
}

// A Mapper that uses Standard SQL Syntax to perform mapping functions and queries
type SQLMapper struct {
	conf    *SQLConf
	sqlConn *sql.DB
}

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

// Original implemenation from : http://www.movable-type.co.uk/scripts/latlong-db.html
// Uses SQL to retrieve all points within the radius of the origin point passed in.
// @param [*Point]. The origin point.
// @param [float64]. The radius (in meters) in which to search for points from the Origin.
// TODO Potentially fallback to PostgreSQL's earthdistance module: http://www.postgresql.org/docs/8.3/static/earthdistance.html
// TODO Determine if valuable to just provide an abstract formula and then select accordingly, might be helpful for NOSQL wrapper
func (s *SQLMapper) PointsWithinRadius(p *Point, radius float64) (*sql.Rows, error) {
	select_str := fmt.Sprintf("SELECT * FROM %s a", s.conf.table)
	lat1 := fmt.Sprintf("sin(radians(%f)) * sin(radians(a.lat))", p.lat)
	lng1 := fmt.Sprintf("cos(radians(%f)) * cos(radians(a.lat)) * cos(radians(a.lng) - radians(%f))", p.lat, p.lng)
	where_str := fmt.Sprintf("WHERE acos(%s + %s) * %f <= %f", lat1, lng1, 6356.7523, radius)
	query := fmt.Sprintf("%s %s", select_str, where_str)

	res, err := s.sqlConn.Query(query)
	if err != nil {
		panic(err)
	}

	return res, err
}