package geo

import ( "testing" 
	_ "github.com/StefanSchroeder/Golang-Ellipsoid"
)

// @spec:
//   - Should correctly return a set of [lat, lng] within a certain radius
func TestWithin(t *testing.T) {
	// TODO create a bunch of [lat, lng] points around the edge of radius
	//      Assert that they are all "within"
	// 
	//      Create a series of [lat, lng] that are outside of radius
	//      Assert that they are not "within"

	geo1 := ellipsoid.Init("WGS84", ellipsoid.Degrees, ellipsoid.Mile, ellipsoid.Longitude_is_symmetric, ellipsoid.Bearing_is_symmetric)

	// SFO
	origin := &Point{37.619002, -122.37484}

	// Create set of "within" points	
	withinPoints := make([]Point, 360)
	for i := 0; i < 360; i++ {
		next_lat, next_lng := geo1.calculateTargetLocation(origin.lat, origin.lng, 4.9, i)
		next_point = &Point{lat:next_lat, lng:next_lng}
		withinPoints[i] = next_point
		// TODO save to db
	}

	// Create set of "outside" points
	outsidePoints := make([]Point, 360)
	for i := 0; i < 360; i++ {
		next_lat, next_lng := geo1.calculateTargetLocation(origin.lat, origin.lng, 5.1, i)
		next_point = &Point{lat:next_lat, lng:next_lng}
		outsidePoints[i] = next_point
		// TODO save to db
	}
}