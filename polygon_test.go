package geo

import (
	"encoding/json"
	"os"
	"testing"
)

// Tests a point is in a real geo polygon
func TestPointInPolygon(t *testing.T) {
	// Contour is the outline polygon of Brunei made up of points: (Long, Lat)
	brunei, err := polygonFromFile("test/data/brunei.json")
	if err != nil {
		t.Error("brunei json file failed to parse: ", err)
	}

	// See if the capital city of brunei is inside the Brunei polygon?
	point := Point{lng: 114.9480600, lat: 4.9402900}
	if !brunei.Contains(&point) {
		t.Error("Expected the capital of Brunei to be in Brunei, but it wasn't.")
	}
}

// Ensures that the polygon logic can correctly identify if a polygon
// does not contain a point.
func TestPointNotInPolygon(t *testing.T) {
	// Contour is the outline polygon of Brunei made up of points: (Long, Lat)
	brunei, err := polygonFromFile("test/data/brunei.json")
	if err != nil {
		t.Error("brunei json file failed to parse: ", err)
	}

	// Seattle, WA should not be inside of Brunei
	point := NewPoint(47.45, 122.30)
	if brunei.Contains(point) {
		t.Error("Seattle, WA [47.45, 122.30] should not be inside of Brunei")
	}

	// A point just outside of the successful bounds in Brunei
	// Should not be contained in the Polygon
	precision := NewPoint(114.659596, 4.007636)
	if brunei.Contains(precision) {
		t.Error("A point just outside of Brunei should not be contained in the Polygon")
	}

}

// Tests a point is in a real geo polygon that has a hole in it, e.g. a donut
func TestPointInPolygonWithHole(t *testing.T) {
	nsw, err := polygonFromFile("test/data/nsw.json")
	if err != nil {
		t.Error("nsw json file failed to parse: ", err)
	}

	act, err := polygonFromFile("test/data/act.json")
	if err != nil {
		t.Error("act json file failed to parse: ", err)
	}

	// Look at two contours
	canberra := Point{lng: 149.128684300000030000, lat: -35.2819998}
	isnsw := nsw.Contains(&canberra)
	isact := act.Contains(&canberra)
	if !isnsw && !isact {
		t.Error("Canberra should be in NSW and also in the sub-contour ACT state")
	}

	// Using NSW as a multi-contour polygon
	nswmulti := &Polygon{}
	for _, p := range nsw.Points() {
		nswmulti.Add(p)
	}

	for _, p := range act.Points() {
		nswmulti.Add(p)
	}

	isnsw = nswmulti.Contains(&canberra)
	if isnsw {
		t.Error("Canberra should not be in NSW as it falls in the donut contour of the ACT")
	}

	sydney := Point{lng: 151.209, lat: -33.866}

	if !nswmulti.Contains(&sydney) {
		t.Error("Sydney should be in NSW")
	}

	losangeles := Point{lng: 118.28333, lat: 34.01667}
	isnsw = nswmulti.Contains(&losangeles)

	if isnsw {
		t.Error("Los Angeles should not be in NSW")
	}

}

type Points struct {
	Points []*Point
}

// Open a JSON file and unpack the polygon
func polygonFromFile(filename string) (*Polygon, error) {
	p := &Polygon{}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	points := new(Points)
	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(&points); err != nil {
		return nil, err
	}

	for _, point := range points.Points {
		p.Add(point)
	}

	return p, nil
}
