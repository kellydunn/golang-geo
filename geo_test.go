package geo

import (
	"fmt"
	"testing"
	"strconv"
	"database/sql"
)

// TODO This paticular test is just one big integration for using the entire library.
//      Should seperate this out into sperate tests once I determine an effective
//      And reasonable way to test formulae and configuration handling.
// @spec: golang-geo should 
//   - Should correctly return a set of [lat, lng] within a certain radius
func TestFullIntegration(t *testing.T) {
	s, _ := HandleWithSQL()

	// SFO
	origin := &Point{37.619002, -122.37484}

	for i := 0; i < 360; i++ {
		bearing := (float64)(i * 1.0)

		in_point := origin.PointAtDistanceAndBearing(7.9, bearing)
		s.sqlConn.Exec(fmt.Sprintf("INSERT INTO points(lat, lng) VALUES(%f, %f);", in_point.lat, in_point.lng))

		gcd := RoundFloat(in_point.GreatCircleDistance(origin), 2)
		if (gcd != 7.9) {
			t.Error("ERROR: Expected certain Great Circle Distance", gcd)
		}

		out_point := origin.PointAtDistanceAndBearing(8.1, bearing)
		s.sqlConn.Exec(fmt.Sprintf("INSERT INTO points(lat, lng) VALUES(%f, %f);", out_point.lat, out_point.lng))

		gcd = RoundFloat(out_point.GreatCircleDistance(origin), 2)
		if (gcd != 8.1) {
			t.Error("ERROR: Expected certain Great Circle Distance", gcd)
		}

	}

	//  Should be able to calculate 360 points within 8km
	res, err := s.PointsWithinRadius(origin, 8)
	if err != nil {
		panic(err)
	}

	count := ResultsCount(res)

	if count < 360 {
		t.Error("Expected 360 results to be within 8km of origin.  Got", count)
	}

	//  Should be able to calculate 720 points within 9km
	res2, err2 := s.PointsWithinRadius(origin, 9)
	if err2 != nil {
		panic(err2)
	}

	count = ResultsCount(res2)

	if count < 720 {
		t.Error("Expected 720 results to be within 9km of origin.  Got", count)
	}

	// Clear Test DB
	FlushTestDB(s)
}

// TODO Test sql configuration
// TODO Test Great Circle Distance
// TODO Test Point At Distance And Bearing

func FlushTestDB(s *SQLMapper) {
	s.sqlConn.Exec("DELETE FROM points;")
}

func ResultsCount(res *sql.Rows) int {
	count := 0

	// TODO Am I missing a res.Len()?
	for res.Next() {
		count += 1
	}

	return count
}

// Taken from: http://play.golang.org/p/cwJj8ZJUhl
func RoundFloat(x float64, prec int) float64 {
	frep := strconv.FormatFloat(x, 'g', prec, 64)
	f, _ := strconv.ParseFloat(frep, 64)
	return f
}