package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// This struct contains all the funcitonality
// of interacting with the OpenCage Geocoding Service
type OpenCageGeocoder struct{}

// This struct contains selected fields from OpenCage's Geocoding Service response
type opencageGeocodeResponse struct {
	Results []struct {
		Formatted string `json:"formatted"`
		Geometry  struct {
			Lat float64
			Lng float64
		}
	}
}

type opencageReverseGeocodeResponse struct {
	Results []struct {
		Formatted string `json:"formatted"`
	}
}

// This is the error that consumers receive when there
// are no results from the geocoding request.
var opencageZeroResultsError = errors.New("ZERO_RESULTS")

// This contains the base URL for the Mapquest Geocoder API.
var opencageGeocodeURL = "http://api.opencagedata.com/geocode/v1/json"

var OpenCageAPIKey = ""

// Note:  In the next major revision (1.0.0), it is planned
//        That Geocoders should adhere to the `geo.Geocoder`
//        interface and provide versioning of APIs accordingly.
// Sets the base URL for the OpenCage Geocoding API.
func SetOpenCageGeocodeURL(newGeocodeURL string) {
	opencageGeocodeURL = newGeocodeURL
}

func SetOpenCageAPIKey(newAPIKey string) {
	OpenCageAPIKey = newAPIKey
}

// Issues a request to the open OpenCage API geocoding services using the passed in url query.
// Returns an array of bytes as the result of the api call or an error if one occurs during the process.
// Note: Since this is an arbitrary request, you are responsible for passing in your API key if you want one.
func (g *OpenCageGeocoder) Request(url string) ([]byte, error) {
	client := &http.Client{}
	fullUrl := fmt.Sprintf("%s/%s", opencageGeocodeURL, url)

	// TODO Refactor into an api driver of some sort
	//      It seems odd that golang-geo should be responsible of versioning of APIs, etc.
	req, _ := http.NewRequest("GET", fullUrl, nil)
	resp, requestErr := client.Do(req)

	if requestErr != nil {
		return nil, requestErr

	}

	// TODO figure out a better typing for response
	data, dataReadErr := ioutil.ReadAll(resp.Body)

	if dataReadErr != nil {
		return nil, dataReadErr
	}

	return data, nil
}

// Returns the first point returned by OpenCage's geocoding service or an error
// if one occurs during the geocoding request.
func (g *OpenCageGeocoder) Geocode(address string) (*Point, error) {

	queryStr, err := opencageGeocodeQueryStr(address)
	if err != nil {
		return nil, err
	}

	data, err := g.Request(queryStr)
	if err != nil {
		return nil, err
	}

	res := &opencageGeocodeResponse{}
	json.Unmarshal(data, res)

	if len(res.Results) == 0 {
		return nil, opencageZeroResultsError
	}

	lat := res.Results[0].Geometry.Lat
	lng := res.Results[0].Geometry.Lng

	point := &Point{
		lat: lat,
		lng: lng,
	}

	return point, nil
}

func opencageGeocodeQueryStr(address string) (string, error) {
	url_safe_query := url.QueryEscape(address)

	var queryStr = bytes.NewBufferString("?")
	_, err := queryStr.WriteString(fmt.Sprintf("q=%s", url_safe_query))
	if err != nil {
		return "", err
	}

	if OpenCageAPIKey != "" {
		_, err := queryStr.WriteString(fmt.Sprintf("&key=%s", OpenCageAPIKey))
		if err != nil {
			return "", err
		}
	}

	_, err = queryStr.WriteString("&pretty=1")
	if err != nil {
		return "", err
	}

	return queryStr.String(), err
}

// Returns the first most available address that corresponds to the passed in point.
// It may also return an error if one occurs during execution.
func (g *OpenCageGeocoder) ReverseGeocode(p *Point) (string, error) {
	queryStr, err := opencageReverseGeocodeQueryStr(p)
	if err != nil {
		return "", err
	}

	data, err := g.Request(queryStr)
	if err != nil {
		return "", err
	}

	res := &opencageReverseGeocodeResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return "", err
	}

	if len(res.Results) == 0 {
		return "", opencageZeroResultsError
	}

	return res.Results[0].Formatted, nil
}

func opencageReverseGeocodeQueryStr(p *Point) (string, error) {
	var queryStr = bytes.NewBufferString("?")
	_, err := queryStr.WriteString(fmt.Sprintf("q=%f,%f", p.lat, p.lng))
	if err != nil {
		return "", err
	}

	if OpenCageAPIKey != "" {
		_, err := queryStr.WriteString(fmt.Sprintf("&key=%s", OpenCageAPIKey))
		if err != nil {
			return "", err
		}
	}

	_, err = queryStr.WriteString("&pretty=1")
	if err != nil {
		return "", err
	}

	return queryStr.String(), err
}
