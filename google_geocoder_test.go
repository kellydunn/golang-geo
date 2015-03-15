package geo

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

/// TODO Test extracting Address from Google Reverse Geocoding Response
func TestExtractAddressFromResponse(t *testing.T) {
	g := &GoogleGeocoder{}

	data, err := GetMockResponse("test/data/google_reverse_geocode_success.json")
	if err != nil {
		t.Error("%v\n", err)
	}

	address := g.extractAddressFromResponse(data)
	if address != "285 Bedford Avenue, Brooklyn, NY 11211, USA" {
		t.Error(fmt.Sprintf("Expected: 285 Bedford Avenue, Brooklyn, NY 11211 USA.  Got: %s", address))
	}
}

func TestExtractAddressFromResponseZeroResults(t *testing.T) {
	g := &GoogleGeocoder{}

	data, err := GetMockResponse("test/data/google_geocode_zero_results.json")
	if err != nil {
		t.Error("%v\n", err)
	}

	address := g.extractAddressFromResponse(data)
	if address != "ERROR: ZERO_RESULTS" {
		t.Error(fmt.Sprintf("Expected: 'ERROR: ZERO_RESULTS'.  Got: '%s'", address))
	}
}

// TODO Test extracting LatLng from Google Geocoding Response
func TestExtractLatLngFromRequest(t *testing.T) {
	g := &GoogleGeocoder{}

	data, err := GetMockResponse("test/data/google_geocode_success.json")
	if err != nil {
		t.Error("%v\n", err)
	}

	point, err := g.extractLatLngFromResponse(data)
	if err != nil {
		t.Error("%v\n", err)
	}

	if point.lat != 37.615223 || point.lng != -122.389979 {
		t.Error(fmt.Sprintf("Expected: [37.615223, -122.389979], Got: [%f, %f]", point.lat, point.lng))
	}
}

// TODO Test extracting LatLng from Google Geocoding Response when no results are returned
func TestExtractLatLngFromRequestZeroResults(t *testing.T) {
	g := &GoogleGeocoder{}

	data, err := GetMockResponse("test/data/google_geocode_zero_results.json")
	if err != nil {
		t.Error("%v\n", err)
	}

	_, err = g.extractLatLngFromResponse(data)
	if err != googleZeroResultsError {
		t.Error(fmt.Sprintf("Expected error: %v, Got: %v"), googleZeroResultsError, err)
	}
}

func GetMockResponse(s string) ([]byte, error) {
	dataPath := path.Join(s)
	_, readErr := os.Stat(dataPath)
	if readErr != nil && os.IsNotExist(readErr) {
		return nil, readErr
	}

	handler, handlerErr := os.Open(dataPath)
	if handlerErr != nil {
		return nil, handlerErr
	}

	data, readErr := ioutil.ReadAll(handler)
	if readErr != nil {
		return nil, readErr
	}

	return data, nil
}
