package main

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

func (s *SQLMapper) Within(p *Point, radius float32) (* sql.Rows, error) {
	// Taken from : http://www.movable-type.co.uk/scripts/latlong.html
	query := fmt.Sprintf("SELECT * FROM %s a WHERE (acos(sin(a.%s * 0.0175) * sin(%f * 0.0175) + cos(a.%s * 0.0175) * cos(%f * 0.0175) * cos((%f * 0.0175) - (a.%s * 0.0175))) * 3959 <= %f) ", s.conf.table, s.conf.latCol, p.lat, s.conf.latCol, p.lat, p.lng, s.conf.lngCol, radius)

	res, err := s.sqlConn.Query(query)
	if err != nil {
		panic(err)
	}

	return res, err
}

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
