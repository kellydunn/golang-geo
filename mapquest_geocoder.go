package geo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// A Geocoder that makes use of open street map's geocoding service
type MapQuestGeocoder struct{}

var mapquestZeroResultsError = errors.New("ZERO_RESULTS")

// Issues a request to the open mapquest api geocoding services using the passed in url query.
// Returns an array of bytes as the result of the api call or an error if one occurs during the process.
func (g *MapQuestGeocoder) Request(url string) ([]byte, error) {
	client := &http.Client{}
	fullUrl := fmt.Sprintf("http://open.mapquestapi.com/nominatim/v1/%s", url)

	// TODO Refactor into an api driver of some sort
	//      It seems odd that golang-geo should be responsible of versioning of APIs, etc.
	req, _ := http.NewRequest("GET", fullUrl, nil)
	resp, requestErr := client.Do(req)

	if requestErr != nil {
		panic(requestErr)
	}

	// TODO figure out a better typing for response
	data, dataReadErr := ioutil.ReadAll(resp.Body)

	if dataReadErr != nil {
		return nil, dataReadErr
	}

	return data, nil
}

// Returns the first point returned by MapQuest's geocoding service or an error
// if one occurs during the geocoding request.
func (g *MapQuestGeocoder) Geocode(query string) (*Point, error) {
	url_safe_query := url.QueryEscape(query)
	data, err := g.Request(fmt.Sprintf("search.php?q=%s&format=json", url_safe_query))
	if err != nil {
		return nil, err
	}

	lat, lng, extractErr := g.extractLatLngFromResponse(data)
	if extractErr != nil {
		return nil, extractErr
	}

	p := &Point{lat: lat, lng: lng}

	return p, nil
}

// private

// Extracts the first lat and lng values from a MapQuest response body.
func (g *MapQuestGeocoder) extractLatLngFromResponse(data []byte) (float64, float64, error) {
	res := make([]map[string]interface{}, 0)
	json.Unmarshal(data, &res)

	if len(res) == 0 {
		return 0, 0, mapquestZeroResultsError
	}

	lat, _ := strconv.ParseFloat(res[0]["lat"].(string), 64)
	lng, _ := strconv.ParseFloat(res[0]["lon"].(string), 64)

	return lat, lng, nil
}

// Returns the first most available address that corresponds to the passed in point.
// It may also return an error if one occurs during execution.
func (g *MapQuestGeocoder) ReverseGeocode(p *Point) (string, error) {
	data, err := g.Request(fmt.Sprintf("reverse.php?lat=%f&lon=%f&format=json", p.lat, p.lng))
	if err != nil {
		return "", err
	}

	resStr := g.extractAddressFromResponse(data)

	return resStr, nil
}

// Return sthe first address in the passed in byte array.
func (g *MapQuestGeocoder) extractAddressFromResponse(data []byte) string {
	res := make(map[string]map[string]string)
	json.Unmarshal(data, &res)

	// TODO determine if it's better to have channels to receive this data on
	//      Provides for concurrency during HTTP requests, etc ~
	road, _ := res["address"]["road"]
	city, _ := res["address"]["city"]
	state, _ := res["address"]["state"]
	postcode, _ := res["address"]["postcode"]
	country_code, _ := res["address"]["country_code"]

	resStr := fmt.Sprintf("%s %s %s %s %s", road, city, state, postcode, country_code)
	return resStr
}
