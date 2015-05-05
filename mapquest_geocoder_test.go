package geo

import (
	"encoding/json"
	"fmt"
	"testing"
)

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
