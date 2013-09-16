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
[![Build Status](https://travis-ci.org/kellydunn/go-art.png)](https://travis-ci.org/kellydunn/golang-geo)
# what 

This library provides convience functions for applying translations, geocoding, and finding distances between geographical points.  It is inspired by ruby's `geokit` and `geokit-rails`, and aims to help make dealing with geographical data a little bit easier.

# documentation

You can read the documentation [here](http://godoc.org/github.com/kellydunn/golang-geo)

# usage

Import from github, and get geomancin'

```
import("github.com/kellydunn/golang-geo")
```

Currently, `golang-geo` provides the following functionality:

  - Querying for points within a radius using your own SQL data tables.
  - Calculate a point transposed from a distance at a specific bearing.
  - Calculate the Great Circle Distance bewteen two points.
  - Geocode an Address using Google Maps API or Open Street Maps API.
  - Reverse Geocode a Point using the same services.

## Finding points within a radius

### Using SQL

Currently, the only function that relies on SQL is `PointsWithinRadius`.  You can configure your database acces by providing a `config/geo.yml` file at the root level of your project and connect to your database with the following line of code:

```
db, err := geo.HandleWithSQL()
```

Find all of the points of interest that are in a 5km radius of [42.333, 121,111]
You could also probably use PostgreSQL's built-in earth distance module :P 
http://www.postgresql.org/docs/8.3/static/earthdistance.html

```
p := &Point{lat: 42.3333, lng: 121.111}
res, _ := db.PointsWithinRadius(p, 5)
```

## Transposing points with a distance and bearing

You can also find a point after transposing another a certain distance(km) with a certain bearing(degrees)

```
p2 := p.PointAtDistanceAndBearing(7.9, 45)
```     

## Great Circle Distance

You can also find the GreatCircleDistance Distance between two points!  

```
distance := p.GreatCircleDistance(p2)
```

## Geocoding

There are now two possible Geocoders you can use with `golang-geo`

  - Google Maps 
  - Open Street Maps (as provided by MapQuest)

Both adhere to the Geocoder interface, which currently specifies a `Geocode` and `ReverseGeocode` method.  `Geocode` Accepts a string address and returns to you the first point found in the json response of each service:

```
g := &GoogleGeocoder{}
p, _ := g.Geocode("San Francisco International Airport")
```

## Reverse Geocoding

Reverse geocoding accepts a `Point`, and returns the address of the first point found in the json response of the services.

```
address, _ := g.ReverseGeocode(p)
```

# notes

  - `golang-geo` currently only uses metric measurements to do calculations
  - The `GO_ENV` environment variable it used to determine what environment should be used to query your database.  If you wish to run `golang-geo` in a different environment, please specify this variable by either exporting it, adding it to your profile, or prepending your command line executable with `GO_ENV=environment`

# SQL Configuration

Currently, `golang-geo` will attempt to read a `config/geo.yml` file in the root of your project.  If it does not exist, it will use a Default Server configuration with a user named "postgres" with a password "postgres".  If you want to supply a custom database conifguration, feel free to do so by using the template below:

```
// config/geo.yml
development:
  driver: postgres
  openStr: user=username password=password dbname=points sslmode=disable
  table: points
  latCol: lat
  lngCol: lng
```

You can currently configure which `table` the SQLMapper queries on, as well as the latitude and columns it uses to do all of its math (`latCol` and `lngCol`, respectively).

Keep in mind that `golang-geo` does not provision your database.  You must supply migrations, or otherwise manually alter your database to contain the table and columns provided in your SQL Configuration.

Thanks! ｡◕‿◕｡

# roadmap
  - More Tests!
  - Redis / NOSQL Mapper
  - Bing Maps?
  - Add an abstraction layer for PostgreSQL earthdistance / PostGIS

# testing

By default, `golang-geo` will attempt to run its test suite against a PostgreSQL database.  However, you may run the tests with mocked SQL queries by specifying that you want to run against a mocked database connection via the command line:

```
DB=mock go test
```

# contributing
  - Fork
  - Create a topic branch
  - Make dem commits!
  - Write dem tests!
  - Submit Pull Request once Tests are Passing
  - do this (づ￣ ³￣)づ