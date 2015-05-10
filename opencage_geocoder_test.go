package geo

import (
	"testing"
)

func TestSetOpencageAPIKey(t * testing.T) {
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

func TestUnmarshalOpenCageGeocoderResponse(t *testing.T) {
	// TODO implement
}
