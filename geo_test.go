// package geo
package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"testing"
)

// TODO This paticular test is just one big integration for using the entire library.
//      Should seperate this out into sperate tests once I determine an effective
//      And reasonable way to test formulae and configuration handling.
// @spec: golang-geo should 
//   - Should correctly return a set of [lat, lng] within a certain radius
func TestPointsWithinRadiusIntegration(t *testing.T) {
	s, _ := HandleWithSQL()

	// SFO
	origin := &Point{37.619002, -122.37484}

	for i := 0; i < 360; i++ {
		bearing := (float64)(i * 1.0)

		in_point := origin.PointAtDistanceAndBearing(7.9, bearing)
		s.sqlConn.Exec(fmt.Sprintf("INSERT INTO points(lat, lng) VALUES(%f, %f);", in_point.lat, in_point.lng))

		gcd := RoundFloat(in_point.GreatCircleDistance(origin), 2)
		if gcd != 7.9 {
			t.Error("ERROR: Expected certain Great Circle Distance", gcd)
		}

		out_point := origin.PointAtDistanceAndBearing(8.1, bearing)
		s.sqlConn.Exec(fmt.Sprintf("INSERT INTO points(lat, lng) VALUES(%f, %f);", out_point.lat, out_point.lng))

		gcd = RoundFloat(out_point.GreatCircleDistance(origin), 2)
		if gcd != 8.1 {
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

func TestMapQuestGeocoderIntegration(t *testing.T) {
	m := &MapQuestGeocoder{}
	
	p1, geocodeErr := m.Geocode("Japantown San Francisco, CA")
	if geocodeErr != nil {
		t.Error("Error Geocoding!")
	}

	if p1 == nil {
		t.Error("Incorrect data response from Geocode")
	}

	res, reverseGeocodeErr := m.ReverseGeocode(p1)
	if reverseGeocodeErr != nil {
		t.Error("Error Reverse Geocoding")
	}
	
	p2, geocodeErr2 := m.Geocode(res)
	if geocodeErr2 != nil {
		t.Error("Error Geocoding Again!")
	}

	if p1.lat != p2.lat {
		fmt.Printf("%f, %f\n", p1.lat, p2.lat)
		t.Error("Longitudes do not match after Geocoding and Reverse Geocoding")
	}

	if p1.lng != p2.lng {
		fmt.Printf("%f, %f\n", p1.lng, p2.lng)
		t.Error("Latitudes do not match after Geocoding and Reverse Geocoding")
	}

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
