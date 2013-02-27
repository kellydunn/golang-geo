package geo

import("math")

// Represents a Physical Point in geographic notation [lat, lng]
// @field [float64] lat. The geographic latitude representation of this point.
// @field [float64] lng. The geographic longitude representation of this point.
type Point struct {
	lat float64
	lng float64
}

// Original Implementation from: http://www.movable-type.co.uk/scripts/latlong.html
// @param [float64] dist.  The arc distance in which to transpose the origin point (in meters).
// @param [float64] bearing.  The compass bearing in which to transpose the origin point (in degrees).
// @return [*Point].  Returns a Point struct populated with the lat and lng coordinates
//                    of transposing the origin point a certain arc distance at a certain bearing.
func (p *Point) PointAtDistanceAndBearing(dist float64, bearing float64) *Point {
	// Earth's radius ~= 6,356.7523km
	// TODO Constantize
	dr := dist / 6356.7523

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

// Original Implementation from: http://www.movable-type.co.uk/scripts/latlong.html
// Calculates the Haversine distance between two points.
// @param [*Point].  The destination point.
// @return [float64].  The distance between the origin point and the destination point.
func (p *Point) GreatCircleDistance(p2 *Point) float64 {
	r := 6356.7523 // km
	dLat := (p2.lat - p.lat) * (math.Pi / 180.0)
	dLon := (p2.lng - p.lng) * (math.Pi / 180.0)

	lat1 := p.lat * (math.Pi / 180.0)
	lat2 := p2.lat * (math.Pi / 180.0)

	a1 := math.Sin(dLat/2) * math.Sin(dLat/2)
	a2 := math.Sin(dLon/2) * math.Sin(dLon/2) * math.Cos(lat1) * math.Cos(lat2)

	a := a1 + a2

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return r * c
}