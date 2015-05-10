package geo

import (
	"fmt"
	"testing"
)

func TestSetOpencageAPIKey(t *testing.T) {
	SetOpenCageAPIKey("foo")
	if OpenCageAPIKey != "foo" {
		t.Errorf("Mismatched value for OpencageAPIKey.  Expected: 'foo', Actual: %s", OpenCageAPIKey)
	}
}

func TestSetOpenCageGeocodeURL(t *testing.T) {
	SetOpenCageGeocodeURL("foo")
	if opencageGeocodeURL != "foo" {
		t.Errorf("Mismatched value for googleGeocoeURL.  Expected: 'foo', Actual: %s", opencageGeocodeURL)
	}
}

func TestOpencageGeocoderQueryStr(t *testing.T) {
	// Empty API Key
	SetOpenCageAPIKey("")
	address := "123 fake st"
	res, err := opencageGeocodeQueryStr(address)
	if err != nil {
		t.Errorf("Error creating query string: %v", err)
	}

	expected := "?q=123+fake+st&pretty=1"
	if res != expected {
		t.Errorf(fmt.Sprintf("Mismatched query string.  Expected: %s.  Actual: %s", expected, res))
	}

	// Set api key to some value
	SetOpenCageAPIKey("foo")
	res, err = opencageGeocodeQueryStr(address)
	if err != nil {
		t.Errorf("Error creating query string: %v", err)
	}

	expected = "?q=123+fake+st&key=foo&pretty=1"
	if res != expected {
		t.Errorf(fmt.Sprintf("Mismatched query string.  Expected: %s.  Actual: %s", expected, res))
	}
}

func TestOpencageReverseGeocoderQueryStr(t *testing.T) {
	// Empty API Key
	SetOpenCageAPIKey("")
	p := &Point{lat: 123.45, lng: 56.78}
	res, err := opencageReverseGeocodeQueryStr(p)
	if err != nil {
		t.Errorf("Error creating query string: %v", err)
	}

	expected := "?q=123.450000,56.780000&pretty=1"
	if res != expected {
		t.Errorf(fmt.Sprintf("Mismatched query string.  Expected: %s.  Actual: %s", expected, res))
	}

	// Set api key to some value
	SetOpenCageAPIKey("foo")
	res, err = opencageReverseGeocodeQueryStr(p)
	if err != nil {
		t.Errorf("Error creating query string: %v", err)
	}

	expected = "?q=123.450000,56.780000&key=foo&pretty=1"
	if res != expected {
		t.Errorf(fmt.Sprintf("Mismatched query string.  Expected: %s.  Actual: %s", expected, res))
	}
}
