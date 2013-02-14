package geo

import (
	"database/sql"
	"fmt"
	_ "github.com/bmizerany/pq"
)

// Represents a Physical Point in geographic notation [lat, lng]
type Point struct {
	lat float64
	lng float64
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

type SQLMapper struct {
	conf    *SQLConf
	sqlConn *sql.DB
}

var DefaultSQLConf = &SQLConf{driver: "postgres", openStr: "user=postgres password=*** dbname=points sslmode=require", table: "points", latCol: "lat", lngCol: "lng"}

// geo.HandleWithSQL
func HandleWithSQL() (*SQLMapper, error) {
	s := &SQLMapper{conf: DefaultSQLConf}

	db, err := sql.Open(s.conf.driver, s.conf.openStr)
	if err != nil {
		panic(err)
	}

	s.sqlConn = db

	return s, err
}

func (s *SQLMapper) Within(p *Point, radius float32) bool {
	// Taken from : http://www.movable-type.co.uk/scripts/latlong.html
	query := fmt.Sprintf("SELECT * FROM %s a WHERE (acos(sin(a.%s * 0.0175) * sin(%f * 0.0175) + cos(a.%s * 0.0175) * cos(%f * 0.0175) * cos((%f * 0.0175) - (a.%s * 0.0175))) * 3959 <= %f) ", s.conf.table, s.conf.latCol, p.lat, s.conf.latCol, p.lat, p.lng, s.conf.latCol, radius)

	_, err := s.sqlConn.Exec(query)
	if err != nil {
		panic(err)
	}

	return false
}

/*
func main() {
	s, err := HandleWithSQL()
	if err != nil {
		panic(err)
	}

	p := &Point{lat: 42.3333, lng: 121.1111}
	s.Within(p, 15)
}
*/
