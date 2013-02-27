package geo

import (
	"encoding/json"
	"fmt"
	_ "github.com/bmizerany/pq"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// A Geocoder that makes use of open street map's geocoding service
type MapQuestGeocoder struct {
	// TODO Figure out a better way to initialize mapquest geocoders
	//   - client ???
	//   - apiUrl ???
}

func (g *MapQuestGeocoder) Request(url string) ([]byte, error) {
	client := &http.Client{}

	fullUrl := fmt.Sprintf("http://open.mapquestapi.com/nominatim/v1/%s", url)
	// TODO Refactor into an api caller perhaps :P
	req, _ := http.NewRequest("GET", fullUrl, nil)
	resp, requestErr := client.Do(req)

	if requestErr != nil {
		panic(requestErr)
	}

	// TODO figure out a better typing for response
	data, dataReadErr := ioutil.ReadAll(resp.Body)

	if dataReadErr != nil {
		//panic(dataReadErr)
		return nil, dataReadErr
	}

	return data, nil
}

// Use MapQuest's open service for geocoding
// @param [String] str.  The query in which to geocode.
func (g *MapQuestGeocoder) Geocode(query string) (*Point, error) {
	url_safe_query := url.QueryEscape(query)
	data, err := g.Request(fmt.Sprintf("search.php?q=%s&format=json", url_safe_query))
	if err != nil {
		return nil, err
	}

	lat, lng := g.extractLatLngFromResponse(data)
	p := &Point{lat: lat, lng: lng}

	return p, nil
}

// private
// @param [[]byte] data.  The response struct from the earlier mapquest request as an array of bytes.
// @return [float64] lat.  The first point's latitude in the response. 
// @return [float64] lng.  The first point's longitude in the response. 
func (g *MapQuestGeocoder) extractLatLngFromResponse(data []byte) (float64, float64) {
	res := make([]map[string]interface{}, 0)
	json.Unmarshal(data, &res)

	lat, _ := strconv.ParseFloat(res[0]["lat"].(string), 64)
	lng, _ := strconv.ParseFloat(res[0]["lon"].(string), 64)

	return lat, lng
}

func (g *MapQuestGeocoder) ReverseGeocode(p *Point) (string, error) {
	data, err := g.Request(fmt.Sprintf("reverse.php?lat=%f&lon=%f&format=json", p.lat, p.lng))
	if err != nil {
		return "", err
	}

	resStr := g.extractAddressFromResponse(data)

	return resStr, nil
}

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
