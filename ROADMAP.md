# `golang-geo v1.0.0` Proposed Featureset

After a year of normal use by external parties, a few desired features have been mentioned to me in person, as well as in the form of commentary on github and google groups.

This document provides a proposal for various features that aim to satisfy a majority of these use cases and concerns while mitigating backwards incompatibilities as much as possible.

## Contents:

- [ ] [Sub Packages](#sub-packages)
  - [ ] [Geocoders](#geocoders)
  - [ ] [Indexers](#indexers)
  - [ ] [Point](#point)
- [ ] [API Changes](#api-changes)
  - [ ] [Point](#api-changes-point)
    - [ ] [Exported Fields](#point-exported-fields)
  - [ ] [Geocoders](#api-changes-geocoders)
    - [ ] [Methods](#geocoder-methods)

<a href="sub-packages" />
## Sub Packages

<a href="geocoders"/>
### Geocoders

The `geocoders` package will be used to contain all supported `Geocoder` implementations.  The aim of this package is to provide a default set of `Geocoder`s that make the other functionality provided by `Point` to be interesting and useful "right-out-of-the-box" for a majority of use cases.

Pull Requests to implement new `Geocoder`s are welcome, but not required to work with the library.  Much in the spirit of `mymysql` and `libpq` in relation to golang's `database/sql` package; other `Geocoder` implementations can be made outside of the library and pulled in with an `import` statement.

There will be three default supported `Geocoder`s upon the release of `1.0.0`:
  - Google Geocoder
  - Mapquest Geocoder
  - Opencage Geocoder

<a href="indexers" />
### Indexers
The `indexers` package provides implementation for various `Indexer`s.  `Indexer`s are analogous with `Mapper`s from previous packages; they have been re-named to illustrate their responsibility:  To index and find `Point`s with various types of Spatial Access patterns.  This change is being introduced to address the use cases of indexing points with other types of spatial access patterns, like RTrees, GIS, or other methods.  This means that an `Indexer` interface will be introduced, which will be able to find points within a bounding box, radius, and search for a single point.

There will be one supported `Indexer` upon the release of `1.0.0`:
  - `MySQLIndexer` (previously known as `SQLMapper`)

<a href="point" />
### Point
The `point` package will provide all interesting `Point` operations.

***

<a href="api-changes" />
## API Changes

<a href="api-changes-point" />
### Point

The following parts of the `Point` in this library will be changed wit the release of `1.0.0`:

<a href="point-exported-fields">
#### Exported Fields

- **Lat**: Previously unexported as `lat`, this field will now be exported.
- **Lng**: Previously unexported as `lng`, this field will now be exported.

<a href="api-changes-geocoders" />
### Geocoders

The following illustrates the proposed changes `Geocoders` will receive with the release of `1.0.0`:

<a href="geocoder-methods">
#### Methods

  - `Geocode(string) ([]Point, error)`: `Geocoder`s will now return a slice of `Point` or an `error` when geocoding.
    - Previously, Geocode would only ever return the first point from the response endpoint.  This would return all points.
  - `ReverseGeocode(*Point) ([]string, err)`: `Geocoder`s will now return a slice of `string` or an `error` when reverse geocoding.
    - Previously, ReverseGeocode would only ever return the first address from the response endpoint.  This would return all addresses.