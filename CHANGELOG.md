# golang-geo changelog

## [0.7.0](https://github.com/kellydunn/golang-geo/tree/v0.7.0) December 8, 2015

  - Adds the ability for points to be marshalled an unmarshalled via a binary protocol (gob).

## [0.6.1](https://github.com/kellydunn/golang-geo/tree/v0.6.1) June 24, 2015

  - Ensures that the Google Geocoder operates over HTTPS.

## [0.6.0](https://github.com/kellydunn/golang-geo/tree/v0.6.0) May 22, 2015

  - Introduces the `(p *Point) MidpointTo(other *Point)` function.

## [0.5.4](https://github.com/kellydunn/golang-geo/tree/v0.5.4) May 10, 2015

  - Removes unneeded private methods that were a bit clunky to test
  - Adding in resliency to Geocoding and Reverse Geocoding methods by testing query building methods
  - API key support is now available for all geocoder providers

## [0.5.3](https://github.com/kellydunn/golang-geo/tree/v0.5.3) May 10, 2015

  - Implments a Mapquest Data Transfer Object for decoding JSON responses from the MapQuest API.

## [0.5.2](https://github.com/kellydunn/golang-geo/tree/v0.5.2) May 10, 2015

  - Removes some uncessary pointer logic when unmarshalling various structs from provider responses.

## [0.5.1](https://github.com/kellydunn/golang-geo/tree/v0.5.1) March 14, 2015

  - Resolves an issue with the Google Geocoder that would panic when attempting to Reverse Geocode some Points with no results.

## [0.5.0](https://github.com/kellydunn/golang-geo/tree/v0.5.0) January 14, 2015

  - Exposes `GoogleGeocoder.HttpClient` so that clients may be able to swap out underlying http client implementations.

## [0.4.1](https://github.com/kellydunn/golang-geo/tree/v0.4.1) December 1, 2014

  - Improves geocoder testing.

## [0.4.0](https://github.com/kellydunn/golang-geo/tree/v0.4.0) November 2, 2014

  - Introduces the OpenCage Geocoder

## [0.3.3](https://github.com/kellydunn/golang-geo/tree/v0.3.3) September 1, 2014

  - Fixes some inconsistent documentation.

## [0.3.2](https://github.com/kellydunn/golang-geo/tree/v0.3.2) August 11, 2014

  - Resolves an issue where Reverse Geocoding with a google Geocoder was panicing unexpectedly.
  - Fixes some test conditions for marshalling Points as JSON

## [0.3.1](https://github.com/kellydunn/golang-geo/tree/v0.3.1) June 1, 2014

  - Cleans up some implementation details of how `geo.GoogleGeocoder` query results are handled.

## [0.3.0](https://github.com/kellydunn/golang-geo/tree/v0.3.0) April 29, 2014

  - Introduces `geo.Polygon`, which is composed of many `geo.Points`. (Thanks, @mish15!)
  - Introduces the ability to create a `geo.Polygon` with `NewPolygon` by passing in `[]*geo.Point`
  - Introduces the ability to figure out of a point is contained in a polygon with `*geo.Polygon.Contains`
  - Introduces `*geo.Polygon.IsClosed` which determines if a polygon is a closed shape or not.
  - Improves documentation and testing coverage.
  - Indicates that consumers should use [gopkg.in](http://gopkg.in) in order to download older versions of the library.
  - Increases testing flexibilty by giving Geocoders the ability to specify their own base URL (Thanks, @adams-sarah!)
  - Points now implement the `json.Marshaler` and `json.Unmarshaler` interface!

## [0.2.1](https://github.com/kellydunn/golang-geo/tree/v0.2.1) Februrary 24, 2014

  - Introduces some bugfixes for google maps and mapquest api error handling
  - Improved some documentation

## [0.2.0](https://github.com/kellydunn/golang-geo/tree/v0.2.0) Februrary 10, 2014

  - Introduces `*Point.BearingTo`, which finds the initial bearing (or forward azimuth) from one point to another

## [0.1.0](https://github.com/kellydunn/golang-geo/tree/v0.1.0) January 25, 2014

  - Introducing `geo.NewSQLMapper`, which creates and returns a pointer to a new `geo.SQLMapper`.  This solved issues where users had to create extraneous `*sql.DB` in order to perform Mapper operations.  The introduction of this method signature marks `geo.HandleWithSQL` a canidate for removal in a Major Patch verison.
  - Introducing `geo.SQLMapper.SqlDBConn`, which returns the database connection of the `geo.SQLMapper` for inspection purposes.
  - Introduces `geo.GetSQLConfFromFile`, which accepts the pathname to the desired configuration file.  This was necessary in order for `NewSQLMapper` to allow users to supply different pathnames to create new `SQLMapper`s.
  - Various Bugfixes including some mismatched `EARTH_RADIUS` logic in querying points in SQL databases.

## [0.0.2](https://github.com/kellydunn/golang-geo/tree/v0.0.2) November 28, 2013

  - Change `EARTH_RADIUS` to comply with the information published in [wikipedia](http://en.wikipedia.org/wiki/Earth_radius).
  - Added some more documentation to publicly available structs and methods.

## [0.0.1](https://github.com/kellydunn/golang-geo/tree/v0.0.1) November 28, 2013

  - First tagged release.
