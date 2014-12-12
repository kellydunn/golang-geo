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

// This struct contains all the funcitonality
// of interacting with the MapQuest Geocoding Service
type MapQuestGeocoder struct{}

type mapQuestGeocodeResponse struct {
	Results []struct {
		BoundingBox struct {
			Box []string
		} `json:"boundingbox"`
		Lat string
		Lng string `json:"lon"`
		DisplayName string `json:"display_name"`
	}
}

type mapQuestReverseGeocodeResponse struct {
	Results []struct {
		Address struct {
			Road string
			City string
			State string
			PostCode string `json:"postcode"`
			CountryCode string `json:"country_code"`
		}
	}
}

// This is the error that consumers receive when there
// are no results from the geocoding request.
var mapquestZeroResultsError = errors.New("ZERO_RESULTS")

// This contains the base URL for the Mapquest Geocoder API.
var mapquestGeocodeURL = "http://open.mapquestapi.com/nominatim/v1"

// Note:  In the next major revision (1.0.0), it is planned
//        That Geocoders should adhere to the `geo.Geocoder`
//        interface and provide versioning of APIs accordingly.
// Sets the base URL for the MapQuest Geocoding API.
func SetMapquestGeocodeURL(newGeocodeURL string) {
	mapquestGeocodeURL = newGeocodeURL
}

// Issues a request to the open mapquest api geocoding services using the passed in url query.
// Returns an array of bytes as the result of the api call or an error if one occurs during the process.
func (g *MapQuestGeocoder) Request(url string) ([]byte, error) {
	client := &http.Client{}
	fullUrl := fmt.Sprintf("%s/%s", mapquestGeocodeURL, url)

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

	point, extractErr := g.extractLatLngFromResponse(data)
	if extractErr != nil {
		return nil, extractErr
	}
	
	res := &mapQuestGeocodeResponse{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, mapquestZeroResultsError
	}

	p := &Point{
		lat: res[0].Lat, 
		lng: res[0].Lng,
	}

	return p, nil
}

// Returns the first most available address that corresponds to the passed in point.
// It may also return an error if one occurs during execution.
func (g *MapQuestGeocoder) ReverseGeocode(p *Point) (string, error) {
	data, err := g.Request(fmt.Sprintf("reverse.php?lat=%f&lon=%f&format=json", p.lat, p.lng))
	if err != nil {
		return "", err
	}

	res := &mapQuestReverseGeocodeResponse{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return "", err
	}

	if res == nil {
		return nil, mapquestZeroResultsError
	}

	resStr := fmt.Sprintf("%s %s %s %s %s", road, city, state, postcode, country_code)

	return resStr, nil
}