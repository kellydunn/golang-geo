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
