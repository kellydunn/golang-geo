package geo

import (
	_ "github.com/StefanSchroeder/Golang-Ellipsoid/ellipsoid"
	"testing"
)

// @spec:
//   - Should correctly return a set of [lat, lng] within a certain radius
func TestWithin(t *testing.T) {
	// TODO create a bunch of [lat, lng] points around the edge of radius
	//      Assert that they are all "within"
	// 
	//      Create a series of [lat, lng] that are outside of radius
	//      Assert that they are not "within"

	s, err := geo.HandleWithSql()

	geo1 := ellipsoid.Init("WGS84", ellipsoid.Degrees, ellipsoid.Mile, ellipsoid.Longitude_is_symmetric, ellipsoid.Bearing_is_symmetric)

	// SFO
	origin := &Point{37.619002, -122.37484}

	in_points = make([]Points, 360)
	out_points = make([]Points, 360)
	for i := 0; i < 360; i++ {
		in_lat, in_lng := geo1.calculateTargetLocation(origin.lat, origin.lng, 4.9, i)
		s.sqlConn.Execute(fmt.Sprintf("INSERT INTO points(lat %f, lng %f);", in_lat, in_lng))

		out_lat, out_lng := geo1.calculateTargetLocation(origin.lat, origin.lng, 5.1, i)
		s.sqlConn.Execute(fmt.Sprintf("INSERT INTO points(lat %f, lng %f);", in_lat, in_lng))

		in_point := &Point{lat: in_lat, lng: in_lng}
		in_points[i] = in_point

		out_point := &Point{lat: out_lat, lng: out_lng}
		out_point[i] = out_point
	}

	res, err2 := s.Within(origin, 5)

	count := 0
	for res.Next() {
		count += 1
	}
	if count < 360 {
		t.Error("Expected 360 results to be within 5 miles of origin.")
	}

}
