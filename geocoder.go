package geo

// Geocoder interface
type Geocoder interface {
	Geocode(query string) (*Point, error)
	ReverseGeocode(p *Point) (string, error)
}
