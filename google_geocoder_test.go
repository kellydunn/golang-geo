package geo

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestSetGoogleAPIKey(t * testing.T) {
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
