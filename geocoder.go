package geo

// This interface describes a Geocoder, which provides the ability to Geocode and Reverse Geocode geographic points of interest.
// Geocoding should accept a string that represents a street address, and returns a pointer to a Point that most closely identifies it.
// Reverse geocoding should accept a pointer to a Point, and return the street address that most closely represents it.
type Geocoder interface {
	Geocode(query string) (*Point, error)
	ReverseGeocode(p *Point) (string, error)
}
