package geo

import (
	"fmt"
	"testing"
)

func FlushTestDB(s *SQLMapper) {
	s.sqlConn.Exec("DELETE FROM points;")
}

//func TestPointAtDistanceAndBearing(t *testing.T) {
	// TODO Implement
//}

// @spec:
//   - Should correctly return a set of [lat, lng] within a certain radius
func TestPointsWithinRadius(t *testing.T) {
	s, _ := HandleWithSQL()

	// SFO
	origin := &Point{37.619002, -122.37484}

	for i := 0; i < 360; i++ {
		bearing := (float64)(i * 1.0)

		in_point := origin.PointAtDistanceAndBearing(7.9, bearing)
		s.sqlConn.Exec(fmt.Sprintf("INSERT INTO points(lat, lng) VALUES(%f, %f);", in_point.lat, in_point.lng))

		out_point := origin.PointAtDistanceAndBearing(8.1, bearing)
		s.sqlConn.Exec(fmt.Sprintf("INSERT INTO points(lat, lng) VALUES(%f, %f);", out_point.lat, out_point.lng))

	}

	//  Should be able to calculate 360 points within 8km
	res, err := s.PointsWithinRadius(origin, 8)
	if err != nil {
		panic(err)
	}

	count := 0
	// TODO Am I missing a res.Len()?
	for res.Next() {
		count += 1
	}

	if count < 360 {
		t.Error("Expected 360 results to be within 8km of origin.  Got", count)
	}

	//  Should be able to calculate 720 points within 9km
	res2, err2 := s.PointsWithinRadius(origin, 9)
	if err2 != nil {
		panic(err2)
	}

	count2 := 0
	// TODO Am I missing a res.Len()?
	for res2.Next() {
		count2 +=1
	}

	if count2 < 720 {
		t.Error("Expected 720 results to be within 9km of origin.  Got", count)
	}

	FlushTestDB(s)
}
