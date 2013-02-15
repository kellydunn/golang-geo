package geo

import (
	"database/sql"
	"fmt"
	_ "github.com/bmizerany/pq"
	"math"
)

// Represents a Physical Point in geographic notation [lat, lng]
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

// The Mapper interface
type Mapper interface {
	Within(p *Point, radius int) bool
}

type SQLConf struct {
	driver  string
	openStr string
	table   string
	latCol  string
	lngCol  string
}

// A Mapper that uses Standard SQL Syntax 
// to perform intersting geo-related mapping functions and queries
type SQLMapper struct {
	conf    *SQLConf
	sqlConn *sql.DB
}

var DefaultSQLConf = &SQLConf{driver: "postgres", openStr: "user=postgres password=*** dbname=points sslmode=disable", table: "points", latCol: "lat", lngCol: "lng"}

// @return [*SQLMapper]. An instantiated SQLMapper struct with the DefaultSQLConf.
// @return [Error]. Any error that might have occured during instantiating the SQLMapper.  
func HandleWithSQL() (*SQLMapper, error) {
	s := &SQLMapper{conf: DefaultSQLConf}

	db, err := sql.Open(s.conf.driver, s.conf.openStr)
	if err != nil {
		panic(err)
	}

	s.sqlConn = db

	return s, err
}

// Original implemenation from : http://www.movable-type.co.uk/scripts/latlong-db.html
func (s *SQLMapper) Within(p *Point, radius float32) (*sql.Rows, error) {
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

/*
TODO Incoporate into README
func main() {
	s, err := HandleWithSQL()
	if err != nil {
		panic(err)
	}

	p := &Point{lat: 42.333, lng: 121.111}
	rows, err2 := s.Within(p, 15)
	if err2 != nil {
		panic(err)
	}

	for rows.Next() {
		var lat float32
		var lng float32
		err = rows.Scan(&lat, &lng)
		if err != nil {
			panic(err)
		}
		fmt.Printf("[%f, %f]", lat, lng)
	}
}
*/
