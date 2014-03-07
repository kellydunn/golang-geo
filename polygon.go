// Part of this code originally comes from https://github.com/akavel/polyclip-go
// Also added other functions and some tests related to geo based polygons.
// Polygon format follows geoJSON order of polygon construction: http://geojson.org/geojson-spec.html

package geo

import (
	"math"
)

// Contour represents a sequence of vertices connected by line segments, forming a closed shape.
type Contour []Point


// Add is a convenience method for appending a point to a contour.
func (c *Contour) Add(p Point) {
	*c = append(*c, p)
}

// Polygon is carved out of a 2D plane by a set of (possibly disjoint) contours.
// It can thus contain holes, and can be self-intersecting.
type Polygon []Contour

// Add is a convenience method for appending a contour to a polygon.
func (p *Polygon) Add(c Contour) {
	*p = append(*p, c)
}


// Checks if a point is inside a contour using the "point in polygon" raycast method.
// This works for all polygons, whether they are clockwise or counter clockwise,
// convex or concave.
// See: http://en.wikipedia.org/wiki/Point_in_polygon#Ray_casting_algorithm
// Returns true if p is inside the polygon defined by contour.
func (c Contour) Contains(p Point) bool {
	// Cast ray from p.x towards the right
	intersections := 0
	for i := range c {
		curr := c[i]
		ii := i + 1
		if ii == len(c) {
			ii = 0
		}
		next := c[ii]

		if (p.lng >= next.lng || p.lng <= curr.lng) &&
			(p.lng >= curr.lng || p.lng <= next.lng) {
			continue
		}
		// Edge is from curr to next.
		if p.lat >= math.Max(curr.lat, next.lat) ||
			next.lng == curr.lng {
			continue
		}

		// Find where the line intersects...
		xint := (p.lng-curr.lng)*(next.lat-curr.lat)/(next.lng-curr.lng) + curr.lat
		if curr.lat != next.lat && p.lat > xint {
			continue
		}

		intersections++
	}

	return intersections%2 != 0
}

// For geoJSON polygons, the first polygon is the outer polygon, all secondary
// polygons are internal cut outs. e.g. the centre of a donut.
func (poly Polygon) Contains(p Point) bool {
	for i, c:= range poly {
		if i == 0 && !c.Contains(p) {
			return false
		}
		if i != 0 && c.Contains(p) {
			return false
		}
	}
	return true
}