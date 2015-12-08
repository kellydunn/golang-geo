package geo

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"math"
)

// Represents a Physical Point in geographic notation [lat, lng].
type Point struct {
	lat float64
	lng float64
}

const (
	// According to Wikipedia, the Earth's radius is about 6,371km
	EARTH_RADIUS = 6371
)

// Returns a new Point populated by the passed in latitude (lat) and longitude (lng) values.
func NewPoint(lat float64, lng float64) *Point {
	return &Point{lat: lat, lng: lng}
}

// Returns Point p's latitude.
func (p *Point) Lat() float64 {
	return p.lat
}

// Returns Point p's longitude.
func (p *Point) Lng() float64 {
	return p.lng
}

// Returns a Point populated with the lat and lng coordinates
// by transposing the origin point the passed in distance (in kilometers)
// by the passed in compass bearing (in degrees).
// Original Implementation from: http://www.movable-type.co.uk/scripts/latlong.html
func (p *Point) PointAtDistanceAndBearing(dist float64, bearing float64) *Point {

	dr := dist / EARTH_RADIUS

	bearing = (bearing * (math.Pi / 180.0))

	lat1 := (p.lat * (math.Pi / 180.0))
	lng1 := (p.lng * (math.Pi / 180.0))

	lat2_part1 := math.Sin(lat1) * math.Cos(dr)
	lat2_part2 := math.Cos(lat1) * math.Sin(dr) * math.Cos(bearing)

	lat2 := math.Asin(lat2_part1 + lat2_part2)

	lng2_part1 := math.Sin(bearing) * math.Sin(dr) * math.Cos(lat1)
	lng2_part2 := math.Cos(dr) - (math.Sin(lat1) * math.Sin(lat2))

	lng2 := lng1 + math.Atan2(lng2_part1, lng2_part2)
	lng2 = math.Mod((lng2+3*math.Pi), (2*math.Pi)) - math.Pi

	lat2 = lat2 * (180.0 / math.Pi)
	lng2 = lng2 * (180.0 / math.Pi)

	return &Point{lat: lat2, lng: lng2}
}

// Calculates the Haversine distance between two points in kilometers.
// Original Implementation from: http://www.movable-type.co.uk/scripts/latlong.html
func (p *Point) GreatCircleDistance(p2 *Point) float64 {
	dLat := (p2.lat - p.lat) * (math.Pi / 180.0)
	dLon := (p2.lng - p.lng) * (math.Pi / 180.0)

	lat1 := p.lat * (math.Pi / 180.0)
	lat2 := p2.lat * (math.Pi / 180.0)

	a1 := math.Sin(dLat/2) * math.Sin(dLat/2)
	a2 := math.Sin(dLon/2) * math.Sin(dLon/2) * math.Cos(lat1) * math.Cos(lat2)

	a := a1 + a2

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return EARTH_RADIUS * c
}

// Calculates the initial bearing (sometimes referred to as forward azimuth)
// Original Implementation from: http://www.movable-type.co.uk/scripts/latlong.html
func (p *Point) BearingTo(p2 *Point) float64 {

	dLon := (p2.lng - p.lng) * math.Pi / 180.0

	lat1 := p.lat * math.Pi / 180.0
	lat2 := p2.lat * math.Pi / 180.0

	y := math.Sin(dLon) * math.Cos(lat2)
	x := math.Cos(lat1)*math.Sin(lat2) -
		math.Sin(lat1)*math.Cos(lat2)*math.Cos(dLon)
	brng := math.Atan2(y, x) * 180.0 / math.Pi

	return brng
}

// Calculates the midpoint between 'this' point and the supplied point.
// Original implementation from http://www.movable-type.co.uk/scripts/latlong.html
func (p *Point) MidpointTo(p2 *Point) *Point {
	lat1 := p.lat * math.Pi / 180.0
	lat2 := p2.lat * math.Pi / 180.0

	lon1 := p.lng * math.Pi / 180.0
	dLon := (p2.lng - p.lng) * math.Pi / 180.0

	bx := math.Cos(lat2) * math.Cos(dLon)
	by := math.Cos(lat2) * math.Sin(dLon)

	lat3Rad := math.Atan2(
		math.Sin(lat1)+math.Sin(lat2),
		math.Sqrt(math.Pow(math.Cos(lat1)+bx, 2)+math.Pow(by, 2)),
	)
	lon3Rad := lon1 + math.Atan2(by, math.Cos(lat1)+bx)

	lat3 := lat3Rad * 180.0 / math.Pi
	lon3 := lon3Rad * 180.0 / math.Pi

	return NewPoint(lat3, lon3)
}

// Renders the current point to a byte slice.
// Implements the encoding.BinaryMarshaler Interface.
func (p *Point) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, p.lat)
	if err != nil {
		return nil, fmt.Errorf("unable to encode lat %v: %v", p.lat, err)
	}
	err = binary.Write(&buf, binary.LittleEndian, p.lng)
	if err != nil {
		return nil, fmt.Errorf("unable to encode lng %v: %v", p.lng, err)
	}

	return buf.Bytes(), nil
}

func (p *Point) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)

	var lat float64
	err := binary.Read(buf, binary.LittleEndian, &lat)
	if err != nil {
		return fmt.Errorf("binary.Read failed: %v", err)
	}

	var lng float64
	err = binary.Read(buf, binary.LittleEndian, &lng)
	if err != nil {
		return fmt.Errorf("binary.Read failed: %v", err)
	}

	p.lat = lat
	p.lng = lng
	return nil
}

// Renders the current Point to valid JSON.
// Implements the json.Marshaller Interface.
func (p *Point) MarshalJSON() ([]byte, error) {
	res := fmt.Sprintf(`{"lat":%v, "lng":%v}`, p.lat, p.lng)
	return []byte(res), nil
}

// Decodes the current Point from a JSON body.
// Throws an error if the body of the point cannot be interpreted by the JSON body
func (p *Point) UnmarshalJSON(data []byte) error {
	// TODO throw an error if there is an issue parsing the body.
	dec := json.NewDecoder(bytes.NewReader(data))
	var values map[string]float64
	err := dec.Decode(&values)

	if err != nil {
		log.Print(err)
		return err
	}

	*p = *NewPoint(values["lat"], values["lng"])

	return nil
}
