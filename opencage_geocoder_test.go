package geo

import (
	"fmt"
	"testing"
)

// Test extracting LatLng from OpenCage Geocoding Response
func TestOpenCageExtractLatLngFromRequest(t *testing.T) {
	g := &OpenCageGeocoder{}

	data, err := GetMockResponse("test/data/opencage_geocode_success.json")
	if err != nil {
		t.Error("%v\n", err)
	}

	point, err := g.extractLatLngFromResponse(data)
	if err != nil {
		t.Error("%v\n", err)
	}

	if point.lat != -23.5373732 || point.lng != -46.8374628 {
		t.Error(fmt.Sprintf("Expected: [-23.5373732, -46.8374628], Got: [%f, %f]", point.lat, point.lng))
	}
}

func TestOpenCageExtractAddressFromRequest(t *testing.T) {
	g := &OpenCageGeocoder{}

	data, err := GetMockResponse("test/data/opencage_geocode_success.json")
	if err != nil {
		t.Error("%v\n", err)
	}

	address := g.extractAddressFromResponse(data)

	if address != "Rua Cafelândia, Carapicuíba - SP, Brazil" {
		t.Error(fmt.Sprintf("Expected: Rua Cafelândia, Carapicuíba - SP, Brazil, Got: [%s]", address))
	}
}

// Test extracting LatLng from OpenCage Geocoding Response when no results are returned
func TestOpenCageExtractLatLngFromRequestZeroResults(t *testing.T) {
	g := &OpenCageGeocoder{}

	data, err := GetMockResponse("test/data/opencage_geocode_zero_results.json")
	if err != nil {
		t.Error("%v\n", err)
	}

	_, err = g.extractLatLngFromResponse(data)
	if err != opencageZeroResultsError {
		t.Error(fmt.Sprintf("Expected error: %v, Got: %v"), opencageZeroResultsError, err)
	}
}
