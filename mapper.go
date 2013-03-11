package geo

// Provides a Queryable interface for finding Points via some Data Storage mechanism
type Mapper interface {
	PointsWithinRadius(p *Point, radius int) bool
}
