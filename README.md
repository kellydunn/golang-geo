```
              ___                                                              
             /\_ \                                                             
   __     ___\//\ \      __      ___      __               __      __    ___   
 /'_ `\  / __`\\ \ \   /'__`\  /' _ `\  /'_ `\  _______  /'_ `\  /'__`\ / __`\ 
/\ \L\ \/\ \L\ \\_\ \_/\ \L\.\_/\ \/\ \/\ \L\ \/\______\/\ \L\ \/\  __//\ \L\ \
\ \____ \ \____//\____\ \__/.\_\ \_\ \_\ \____ \/______/\ \____ \ \____\ \____/
 \/___L\ \/___/ \/____/\/__/\/_/\/_/\/_/\/___L\ \        \/___L\ \/____/\/___/ 
   /\____/                                /\____/          /\____/             
   \_/__/                                 \_/__/           \_/__/              

♫ around the world ♪
```
[![Build Status](https://drone.io/github.com/kellydunn/golang-geo/status.png)](https://drone.io/github.com/kellydunn/golang-geo/latest)
[![Coverage Status](https://coveralls.io/repos/kellydunn/golang-geo/badge.png?branch=develop)](https://coveralls.io/r/kellydunn/golang-geo?branch=develop)
[![GoDoc](https://godoc.org/github.com/kellydunn/golang-geo?status.svg)](http://godoc.org/github.com/kellydunn/golang-geo)
[![Join the chat at https://gitter.im/kellydunn/golang-geo](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/kellydunn/golang-geo?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

## what 

This library provides convenience functions for translating, geocoding, and calculating distances between geographical points.  It is inspired by ruby's `geokit` and `geokit-rails` gems, and aims to make working with geographical data a little bit easier in golang.

## documentation

You can read the documentation [here](http://godoc.org/github.com/kellydunn/golang-geo).

## usage

Import from github to get started!

```
package main

import("github.com/kellydunn/golang-geo"
       "fmt")

func main() {
     // Make a few points
     p := geo.NewPoint(42.25, 120.2)
     p2 := geo.NewPoint(30.25, 112.2)
     
     // find the great circle distance between them
     dist := p.GreatCircleDistance(p2)
     fmt.Printf("great circle distance: %d\n", dist)
}
```

Currently, `golang-geo` provides the following functionality:

  - Transposing a point for a given distance and bearing.
  - Calculating the Great Circle Distance bewteen two points.
  - Geocoding an address using Google Maps, Mapquest (OpenStreetMap data), OpenCage (OpenStreetMap, twofishes and other data sources) API.
  - Reverse Geocoding a Point using the same services.
  - Querying for points within a radius using your own SQL data tables.

Keep in mind that you do not need to use SQL in order to perform [simple Point operations](http://godoc.org/github.com/kellydunn/golang-geo#Point) and the only function that relies on SQL is [`PointsWithinRadius`](http://godoc.org/github.com/kellydunn/golang-geo#SQLMapper.PointsWithinRadius). 

### using SQL

As of `0.1.0`, `golang-geo` will shift its scope of responsiblity with SQL management.  The library will still support the functions exposed in its public API in the past, however, it will not concern itself so much with creating and maintaining `*sql.DB` connections as it has done in previous versions.  It is suggested that if you are using `geo.HandleWithSql` that you should instead consider creating a `geo.SQLMapper` yourself by calling the newly introduced `geo.NewSQLMapper` method, which accepts a `*sql.DB` connection and a filepath to the configuration file used to inform `golang-geo` of your particular SQL setup.

That being said, `geo.HandleWithSQL` is configured to connect to a SQL database by reading a `config/geo.yml` file in the root level of your project.  If it does not exist, it will use a Default SQL configuration that will use the postgres driver as described by [lib/pq](http://github.com/lib/pq).  The Default SQL configuration will attempt to connect as a user named "postgres" and with the password "postgres" to a database named "points".

#### examples of SQL database configurations

Here are some examples of valid config files that golang-geo knows how to process:

##### PostgreSQL
```
development:
  driver: postgres
  openStr: user=username password=password dbname=points sslmode=disable
  table: points
  latCol: lat
  lngCol: lng
```

##### MySQL
```
development:
  driver: mysql
  openStr: points/username/password
  table: points
  latCol: lat
  lngCol: lng  
```

## notes

  - `golang-geo` currently only uses metric measurements to do calculations
  - The `$GO_ENV` environment variable is used to determine which configuration group in `config.yml` is to be used.  For example, if you wanted to use the PostgreSQL configuration listed above, you could specify `GO_ENV=development` which would read `config.yml` and use the configuration under the root-level key `development`.

### installing older versions of golang-geo

With the advent of [gopkg.in](http://gopkg.in/), you can now install older versions of `golang-geo`!  Consult [CHANGELOG.md](https://github.com/kellydunn/golang-geo/blob/master/CHANGELOG.md) for the version you wish to build against.

## roadmap
  - More Tests!
  - Redis / NOSQL Mapper
  - Bing Maps?
  - Add an abstraction layer for PostgreSQL earthdistance / PostGIS

## testing

By default, `golang-geo` will attempt to run its test suite against a PostgreSQL database.  However, you may run the tests with mocked SQL queries by specifying that you want to do so on the command line:

```
DB=mock GO_ENV=test go test
```

The `$DB` environment variable is used to specify which database you'd like to run the tests against.  You may specify `postgres`, `mysql`, or `mock`.  The [Travis CI builds](https://travis-ci.org/kellydunn/golang-geo) for this project currently runs against all of these when running the test suite.

## contributing
  - Fork the project
  - Create a topic branch (preferably the in the `gitflow` style of `feature/`, `hotfix/`, etc)
  - Make your changes and write complimentary tests to ensure coverage.
  - Submit Pull Request once the full test suite is passing.
  - Pull Requests will then be reviewed by the maintainer and the community and hopefully merged!

Thanks!