package geo

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSetMapquestAPIKey(t *testing.T) {
	SetMapquestAPIKey("foo")
	if MapquestAPIKey != "foo" {
		t.Errorf("Mismatched value for MapQuestAPIKey.  Expected: 'foo', Actual: %s", MapquestAPIKey)
	}
}

func TestSetMapquestGeocodeURL(t *testing.T) {
	SetMapquestGeocodeURL("foo")
	if mapquestGeocodeURL != "foo" {
		t.Errorf("Mismatched value for MapQuestGeocoeURL.  Expected: 'foo', Actual: %s", mapquestGeocodeURL)
	}
}

func TestMapquestGeocoderQueryStr(t *testing.T) {
	// Empty API Key
	SetMapquestAPIKey("")
	address := "123 fake st"
	res, err := mapquestGeocodeQueryStr(address)
	if err != nil {
		t.Errorf("Error creating query string: %v", err)
	}

	expected := "search.php?q=123+fake+st&format=json"
	if res != expected {
		t.Errorf(fmt.Sprintf("Mismatched query string.  Expected: %s.  Actual: %s", expected, res))
	}

	// Set api key to some value
	SetMapquestAPIKey("foo")
	res, err = mapquestGeocodeQueryStr(address)
	if err != nil {
		t.Errorf("Error creating query string: %v", err)
	}

	expected = "search.php?q=123+fake+st&key=foo&format=json"
	if res != expected {
		t.Errorf(fmt.Sprintf("Mismatched query string.  Expected: %s.  Actual: %s", expected, res))
	}
}

func TestMapquestReverseGeocoderQueryStr(t *testing.T) {
	// Empty API Key
	SetMapquestAPIKey("")
	p := &Point{lat: 123.45, lng: 56.78}
	res, err := mapquestReverseGeocodeQueryStr(p)
	if err != nil {
		t.Errorf("Error creating query string: %v", err)
	}

	expected := "reverse.php?lat=123.450000&lng=56.780000&format=json"
	if res != expected {
		t.Errorf(fmt.Sprintf("Mismatched query string.  Expected: %s.  Actual: %s", expected, res))
	}

	// Set api key to some value
	SetMapquestAPIKey("foo")
	res, err = mapquestReverseGeocodeQueryStr(p)
	if err != nil {
		t.Errorf("Error creating query string: %v", err)
	}

	expected = "reverse.php?lat=123.450000&lng=56.780000&key=foo&format=json"
	if res != expected {
		t.Errorf(fmt.Sprintf("Mismatched query string.  Expected: %s.  Actual: %s", expected, res))
	}
}

// Ensures that the Data Transfer Object used
// to get data from the Mapquest Geocoding API is well formed.
func TestMapQuestGeocodeFromRequest(t *testing.T) {
	data, err := GetMockResponse("test/data/mapquest_geocode_success.json")
	if err != nil {
		t.Error("%v\n", err)
	}

	res := []*mapQuestGeocodeResponse{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		t.Error("%v\n", err)
	}

	if len(res) <= 0 {
		t.Error("Unexecpected amount of results for mapquest mock response")
	}

	if res[0].Lat != "37.62181845" || res[0].Lng != "-122.383992092462" {
		t.Error(fmt.Sprintf("Expected: [37.62181845, -122.383992092462], Got: [%s, %s]", res[0].Lat, res[0].Lng))
	}
}
