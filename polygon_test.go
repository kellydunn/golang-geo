package geo

import (
	"testing"
)

// Tests a point is in a real geo polygon
func TestPointInPolygon(t *testing.T) {
	// Contour is the outline polygon of Brunei made up of points: (Long, Lat)
	brunei, err := json2contour("test/data/brunei.json")
	if err != nil {
		t.Error("brunei json file failed to parse: ", err)
	}
	// See if the capital city of brunei is inside the Brunei polygon?
	point := Point{lng: 114.9480600, lat: 4.9402900}
	if !brunei.Contains(&point) {
		t.Error("Expected the capital of Brunei to be in Brunei, but it wasn't.")
	}
}

// Tests a point is in a real geo polygon that has a hole in it, e.g. a donut
func TestPointInPolygonWithHole(t *testing.T) {
	nsw, err := json2contour("test/data/nsw.json")
	if err != nil {
		t.Error("nsw json file failed to parse: ", err)
	}
	act, err := json2contour("test/data/act.json")
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
	nswmulti := new(Polygon)
	nswmulti.Add(nsw)
	nswmulti.Add(act)
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
