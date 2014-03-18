package geo

import (
	"fmt"
	"testing"
)

// TODO Test extracting LatLng from Google Geocoding Response
func TestMapQuestExtractLatLngFromRequest(t *testing.T) {
	g := &MapQuestGeocoder{}

	data, err := GetMockResponse("test/helpers/mapquest_geocode_success.json")
	if err != nil {
		t.Error("%v\n", err)
	}

	lat, lng, err := g.extractLatLngFromResponse(data)
	if err != nil {
		t.Error("%v\n", err)
	}

	if lat != 37.62181845 && lng != -122.383992092462 {
		t.Error(fmt.Sprintf("Expected: [37.62181845, -122.383992092462], Got: [%f, %f]", lat, lng))
	}
}

// TODO Test extracting LatLng from Google Geocoding Response when no results are returned
func TestMapQuestExtractLatLngFromRequestZeroResults(t *testing.T) {
	g := &MapQuestGeocoder{}

	data, err := GetMockResponse("test/helpers/mapquest_geocode_zero_results.json")
	if err != nil {
		t.Error("%v\n", err)
	}

	_, _, err = g.extractLatLngFromResponse(data)
	if err != mapquestZeroResultsError {
		t.Error(fmt.Sprintf("Expected error: %v, Got: %v"), mapquestZeroResultsError, err)
	}
}
