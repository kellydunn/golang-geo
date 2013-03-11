package geo

import (
	"fmt"
	"os"
	"path"
	"testing"
	"io/ioutil"
)

/// TODO Test extracting Address from Google Reverse Geocoding Response
func TestExtractAddressFromResponse(t *testing.T) {
	g := &GoogleGeocoder{}

	dataPath := path.Join("test/helpers/google_reverse_geocode_success.json")
	_, readErr := os.Stat(dataPath)
	if readErr != nil && os.IsNotExist(readErr) {
		t.Error("Could not open test json response for successful google reverse geocode.")
	}

	handler, handlerErr := os.Open(dataPath)
	if handlerErr != nil {
		t.Error(fmt.Sprintf("Error handling %s.", dataPath))
	}

	data, readErr := ioutil.ReadAll(handler)
	if readErr != nil {
		t.Error(fmt.Sprintf("Error reading data from %s.", dataPath))
	}

	address := g.extractAddressFromResponse(data)
	fmt.Printf(address)
}


// TODO Test extracting LatLng from Google Geocoding Response
func TestExtractPointFromRequest() {}