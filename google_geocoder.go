package geo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type GoogleGeocoder struct{}

func (g *GoogleGeocoder) Request(params string) ([]byte, error) {
	client := &http.Client{}

	fullUrl := fmt.Sprintf("http://maps.googleapis.com/maps/api/geocode/json?sensor=false&%s", params)

	// TODO Potentially refactor out from MapQuestGeocoder as well
	req, _ := http.NewRequest("GET", fullUrl, nil)
	resp, requestErr := client.Do(req)

	if requestErr != nil {
		panic(requestErr)
	}

	data, dataReadErr := ioutil.ReadAll(resp.Body)

	if dataReadErr != nil {
		return nil, dataReadErr
	}

	return data, nil
}

func (g *GoogleGeocoder) Geocode(query string) (*Point, error) {
	url_safe_query := url.QueryEscape(query)
	data, err := g.Request(fmt.Sprintf("address=%s", url_safe_query))
	if err != nil {
		return nil, err
	}

	lat, lng := g.extractLatLngFromResponse(data)
	p := &Point{lat: lat, lng: lng}

	return p, nil
}

// private
// TODO Refactor out of MapQuestGeocoder
// @param [[]byte] data.  The response struct from the earlier mapquest request as an array of bytes.
// @return [float64] lat.  The first point's latitude in the response. 
// @return [float64] lng.  The first point's longitude in the response. 
func (g *GoogleGeocoder) extractLatLngFromResponse(data []byte) (float64, float64) {
	res := make(map[string][]map[string]map[string]map[string]interface{}, 0)
	json.Unmarshal(data, &res)

	lat, _ := res["results"][0]["geometry"]["location"]["lat"].(float64)
	lng, _ := res["results"][0]["geometry"]["location"]["lng"].(float64)

	return lat, lng
}

func (g *GoogleGeocoder) ReverseGeocode(p *Point) (string, error) {
	data, err := g.Request(fmt.Sprintf("latlng=%f,%f", p.lat, p.lng))
	if err != nil {
		return "", err
	}

	resStr := g.extractAddressFromResponse(data)

	return resStr, nil
}

func (g *GoogleGeocoder) extractAddressFromResponse(data []byte) string {
	res := make(map[string][]map[string]interface{}, 0)
	json.Unmarshal(data, &res)

	resStr := res["results"][0]["formatted_address"].(string)
	return resStr
}
