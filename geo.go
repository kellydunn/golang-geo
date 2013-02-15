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

var DefaultSQLConf = &SQLConf{driver: "postgres", openStr: "user=postgres password=*** dbname=points sslmode=disable", table: "points", latCol: "lat", lngCol: "lng"}

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

func (s *SQLMapper) Within(p *Point, radius float32) (*sql.Rows, error) {
	// Taken from : http://www.movable-type.co.uk/scripts/latlong-db.html
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
