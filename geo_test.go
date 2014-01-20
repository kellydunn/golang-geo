package geo

import (
	_ "database/sql"
	"fmt"
	"github.com/erikstmartin/go-testdb"
	"os"
	"strconv"
	"testing"
)

// TODO This paticular test is just one big integration for using the entire library.
//      Should seperate this out into sperate tests once I determine an effective
//      And reasonable way to test formulae and configuration handling.
// @spec: golang-geo should
//   - Should correctly return a set of [lat, lng] within a certain radius
func TestPointsWithinRadiusIntegration(t *testing.T) {
	// TODO Determine if we actually need to test SQL logic across databases.
	dbEnv := os.Getenv("DB")
	if dbEnv == "mock" {
		stubPointsWithinRadiusQueries()
	}

	s, sqlErr := HandleWithSQL()

	if sqlErr != nil {
		t.Error("ERROR: %s", sqlErr)
	}

	// SFO
	origin := &Point{37.619002, -122.37484}

	// Straight North
	bearing := 0.0

	// Make a point that is 1 meter within the desired radius
	in_point := origin.PointAtDistanceAndBearing(7.999, bearing)
	s.sqlConn.Exec(fmt.Sprintf("INSERT INTO points(lat, lng) VALUES(%f, %f);", in_point.lat, in_point.lng))

	// Make a point that is 1 meter outsied of the desired radius
	out_point := origin.PointAtDistanceAndBearing(8.001, bearing)
	s.sqlConn.Exec(fmt.Sprintf("INSERT INTO points(lat, lng) VALUES(%f, %f);", out_point.lat, out_point.lng))

	// Should only get the first point
	_, err := s.PointsWithinRadius(origin, 8)
	if err != nil {
		panic(err)
	}

	// TODO Write a test to check for expected results of PointAtDistanceAndBearing

	// Should get both the first point and second point
	_, err2 := s.PointsWithinRadius(origin, 9)
	if err2 != nil {
		panic(err2)
	}

	// TODO Write a test to check for expected results of PointAtDistanceAndBearing

	// Clear Test DB
	FlushTestDB(s)
}

// TODO Test sql configuration
// TODO Test Great Circle Distance
// TODO Test Point At Distance And Bearing

func FlushTestDB(s *SQLMapper) {
	s.sqlConn.Exec("DELETE FROM points;")
}

// Taken from: http://play.golang.org/p/cwJj8ZJUhl
func RoundFloat(x float64, prec int) float64 {
	frep := strconv.FormatFloat(x, 'g', prec, 64)
	f, _ := strconv.ParseFloat(frep, 64)
	return f
}

func stubPointsWithinRadiusQueries() {
	insideRangeQuery := fmt.Sprintf("SELECT * FROM points a WHERE acos(sin(radians(37.619002)) * sin(radians(a.lat)) + cos(radians(37.619002)) * cos(radians(a.lat)) * cos(radians(a.lng) - radians(-122.374840))) * %f <= 8.000000", float64(EARTH_RADIUS))
	testdb.StubQuery(insideRangeQuery, nil)

	outsideRangeQuery := fmt.Sprintf("SELECT * FROM points a WHERE acos(sin(radians(37.619002)) * sin(radians(a.lat)) + cos(radians(37.619002)) * cos(radians(a.lat)) * cos(radians(a.lng) - radians(-122.374840))) * %f <= 9.000000", float64(EARTH_RADIUS))
	testdb.StubQuery(outsideRangeQuery, nil)
}
