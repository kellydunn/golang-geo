package geo

import (
	"database/sql"
	"fmt"
)

// A Mapper that uses Standard SQL Syntax to perform mapping functions and queries
type SQLMapper struct {
	conf    *SQLConf
	sqlConn *sql.DB
}

// Uses SQL to retrieve all points within the radius (in meters) passed in from the origin point passed in.
// Original implemenation from : http://www.movable-type.co.uk/scripts/latlong-db.html
// Returns a pointer to a sql.Rows as a result, or an error if one occurs during the query.
func (s *SQLMapper) PointsWithinRadius(p *Point, radius float64) (*sql.Rows, error) {
	select_str := fmt.Sprintf("SELECT * FROM %v a", s.conf.table)
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
