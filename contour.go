package geo

import (
	"math"
)

// Contour represents a sequence of vertices connected by line segments, forming a closed shape.
type Contour struct {
	Points []*Point
}

// Add is a convenience method for appending a point to a contour.
func (c *Contour) Add(p *Point) {
	c.Points = append(c.Points, p)
}

// Checks if a point is inside a contour using the "point in polygon" raycast method.
// This works for all polygons, whether they are clockwise or counter clockwise,
// convex or concave.
// See: http://en.wikipedia.org/wiki/Point_in_polygon#Ray_casting_algorithm
// Returns true if p is inside the polygon defined by contour.
func (c Contour) Contains(p *Point) bool {
	// Cast ray from p.x towards the right
	intersections := 0
	for i := range c.Points {
		curr := c.Points[i]
		ii := i + 1
		if ii == len(c.Points) {
			ii = 0
		}
		next := c.Points[ii]

		if (p.Lng() >= next.Lng() || p.Lng() <= curr.Lng()) &&
			(p.Lng() >= curr.Lng() || p.Lng() <= next.Lng()) {
			continue
		}
		// Edge is from curr to next.
		if p.Lat() >= math.Max(curr.Lat(), next.Lat()) ||
			next.Lng() == curr.Lng() {
			continue
		}

		// Find where the line intersects...
		xint := (p.Lng()-curr.Lng())*(next.Lat()-curr.Lat())/(next.Lng()-curr.Lng()) + curr.Lat()
		if curr.Lat() != next.Lat() && p.lat > xint {
			continue
		}

		intersections++
	}

	return intersections%2 != 0
}
