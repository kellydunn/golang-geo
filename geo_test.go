package geo

import (
	"fmt"
	"math"
	"testing"
)

func FlushTestDB(s * SQLMapper) {
	s.sqlConn.Exec("DELETE FROM points;")
}

func calculateTargetLocation(lat float64, lng float64, dist float64, bearing float64) (float64, float64) {
	// Earth's radius ~= 6,356.7523km
        dr := dist / 6356.7523

	bearing = (bearing * (math.Pi / 180.0))

	lat1 := (lat * (math.Pi / 180.0))
	lng1 := (lng * (math.Pi / 180.0))

	lat2_part1 := math.Sin(lat1) * math.Cos(dr)
	lat2_part2 := math.Cos(lat1) * math.Sin(dr) * math.Cos(bearing)

	lat2 := math.Asin(lat2_part1 + lat2_part2)

	lng2_part1 := math.Sin(bearing) * math.Sin(dr) * math.Cos(lat1)
	lng2_part2 := math.Cos(dr) - (math.Sin(lat1) * math.Sin(lat2))

	lng2 := lng1 + math.Atan2(lng2_part1, lng2_part2)

	// What why is this here?
	lng2 = math.Mod((lng2 + 3*math.Pi), (2 * math.Pi)) - math.Pi

	lat2 = lat2 * (180.0 / math.Pi)
	lng2 = lng2 * (180.0 / math.Pi)

	return lat2, lng2
}

// @spec:
//   - Should correctly return a set of [lat, lng] within a certain radius
func TestWithin(t *testing.T) {
	s, _ := HandleWithSQL()

	// SFO
	origin := &Point{37.619002, -122.37484}

	in_points := make([]*Point, 360)
	out_points := make([]*Point, 360)

	for i := 0; i < 360; i++ {
		bearing := (float64)(i * 1.0)

		in_lat, in_lng := calculateTargetLocation(origin.lat, origin.lng, 7.9, bearing)
		s.sqlConn.Exec(fmt.Sprintf("INSERT INTO points(lat, lng) VALUES(%f, %f);", in_lat, in_lng))

		out_lat, out_lng := calculateTargetLocation(origin.lat, origin.lng, 8.1, bearing)
		s.sqlConn.Exec(fmt.Sprintf("INSERT INTO points(lat, lng) VALUES(%f, %f);", in_lat, in_lng))

		in_point := &Point{lat: in_lat, lng: in_lng}
		in_points[i] = in_point

		out_point := &Point{lat: out_lat, lng: out_lng}
		out_points[i] = out_point
	}

	res, err := s.Within(origin, 8)
	if err != nil {
		panic(err)
	}

	count := 0
	for res.Next() {
		count += 1
	}

	if count < 360 {
		t.Error("Expected 360 results to be within 8km of origin.  Got", count)
	}

	FlushTestDB(s)
}
