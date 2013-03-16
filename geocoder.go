package geo

type Geocoder interface {
	Geocode(query string) (*Point, error)
	ReverseGeocode(p *Point) (string, error)
}
