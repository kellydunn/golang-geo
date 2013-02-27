package geo

import (
	"fmt"
	"testing"
)

// Seems brittle :\
func TestGreatCircleDistance(t *testing.T) {
	// Test that SEA and SFO are ~ 1091km apart, accurate to 100 meters.
	sea := &Point{lat: 47.44745785, lng: -122.308065668024}
	sfo := &Point{lat: 37.6160933, lng: -122.3924223}
	sfoToSea := 1090.7

	dist := sea.GreatCircleDistance(sfo)

	if !(dist < (sfoToSea+0.1) && dist > (sfoToSea-0.1)) {
		t.Error("Unnacceptable result.", dist)
	}
}

func TestPointAtDistanceAndBearing(t *testing.T) {
	sea := &Point{lat: 47.44745785, lng: -122.308065668024}
	p := sea.PointAtDistanceAndBearing(1090.7, 180)

	// Expected results of transposing point 
	// ~1091km at bearing of 180 degrees
	resultLat := 37.616572
	resultLng := -122.308066

	withinLatBounds := p.lat < resultLat+0.001 && p.lat > resultLat-0.001
	withinLngBounds := p.lng < resultLng+0.001 && p.lng > resultLng-0.001
	if !(withinLatBounds && withinLngBounds) {
		t.Error("Unnacceptable result.", fmt.Sprintf("[%f, %f]", p.lat, p.lng))
	}
}
