package geo

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestSetGoogleAPIKey(t *testing.T) {
	SetGoogleAPIKey("foo")
	if GoogleAPIKey != "foo" {
		t.Errorf("Mismatched value for GoogleAPIKey.  Expected: 'foo', Actual: %s", GoogleAPIKey)
	}
}

func TestSetGoogleGeocodeURL(t *testing.T) {
	SetGoogleGeocodeURL("foo")
	if googleGeocodeURL != "foo" {
		t.Errorf("Mismatched value for googleGeocoeURL.  Expected: 'foo', Actual: %s", googleGeocodeURL)
	}
}

func TestGoogleGeocoderQueryStr(t *testing.T) {
	// Empty API Key
	SetGoogleAPIKey("")
	address := "123 fake st"
	res, err := googleGeocodeQueryStr(address)
	if err != nil {
		t.Errorf("Error creating query string: %v", err)
	}

	expected := "address=123+fake+st"
	if res != expected {
		t.Errorf(fmt.Sprintf("Mismatched query string.  Expected: %s.  Actual: %s", expected, res))
	}

	// Set api key to some value
	SetGoogleAPIKey("foo")
	res, err = googleGeocodeQueryStr(address)
	if err != nil {
		t.Errorf("Error creating query string: %v", err)
	}

	expected = "address=123+fake+st&key=foo"
	if res != expected {
		t.Errorf(fmt.Sprintf("Mismatched query string.  Expected: %s.  Actual: %s", expected, res))
	}
}

func TestGoogleReverseGeocoderQueryStr(t *testing.T) {
	// Empty API Key
	SetGoogleAPIKey("")
	p := &Point{lat: 123.45, lng: 56.78}
	res, err := googleReverseGeocodeQueryStr(p)
	if err != nil {
		t.Errorf("Error creating query string: %v", err)
	}

	expected := "latlng=123.450000,56.780000"
	if res != expected {
		t.Errorf(fmt.Sprintf("Mismatched query string.  Expected: %s.  Actual: %s", expected, res))
	}

	// Set api key to some value
	SetGoogleAPIKey("foo")
	res, err = googleReverseGeocodeQueryStr(p)
	if err != nil {
		t.Errorf("Error creating query string: %v", err)
	}

	expected = "latlng=123.450000,56.780000&key=foo"
	if res != expected {
		t.Errorf(fmt.Sprintf("Mismatched query string.  Expected: %s.  Actual: %s", expected, res))
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
