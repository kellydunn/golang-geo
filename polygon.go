// Part of this code originally comes from https://github.com/akavel/polyclip-go
// Also added other functions and some tests related to geo based polygons.
// Polygon format follows geoJSON order of polygon construction: http://geojson.org/geojson-spec.html

package geo

// A Polygon is carved out of a 2D plane by a set of (possibly disjoint) contours.
// It can thus contain holes, and can be self-intersecting.
type Polygon struct {
	Contours []*Contour
}

// Appends the passed in contour to the current Polygon.
func (p *Polygon) Add(c *Contour) {
	p.Contours = append(p.Contours, c)
}

// Returns whether or not the current Polygon contains the passed in Point.
// For geoJSON polygons, the first polygon is the outer polygon, all secondary
// polygons are internal cut outs. e.g. the centre of a donut.
func (p *Polygon) Contains(point *Point) bool {
	for i, c := range p.Contours {

		if i == 0 && !c.Contains(point) {
			return false
		}

		if i != 0 && c.Contains(point) {
			return false
		}
	}

	return true
}
