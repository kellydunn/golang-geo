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

# what 

Geographical calculations in Go.

Still playing around with the language, seems fun so far! (￣︶￣)♫♪

# usage

Import from github, and get geomancin'

```
import("github.com/kellydunn/golang-geo")
```

Currently, `golang-geo` provides the following functionality:

  - Querying for points within a radius using your own SQL data tables
  - Calculate a point transposed from a distance at a specific bearing
  - Calculate the Great Circle Distance bewteen two points
  - Geocode an Address using Nominatim, the open street map resource (ala open.mapquestapi.com).

## Finding points within a radius

### Using SQL

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

Currently, `golang-geo` only makes use of the openstreetmap API, as provided by mapquest.  Future implementations aim to include Google Maps Geocoding, Bing Maps, and potentially others!

```
p := geo.Geocode("San Francisco International Airport")
```

## Reverse Geocoding

As mentioned above, `golang-geo` currently only supports Reverse Geocoding as it is offered by nominatim, the openstreetmap API.

```
address := geo.ReverseGeocode(p)
```

# notes

  - `golang-geo` currently only uses metric measurements to do calculations
  - You do not _need_ to use SQL in order to use this library.  Instead, you may import it and just use it on `Point` specific operations like `GreatCircleDistance` and `PointAtDistanceAndBearing`
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
  - More sane file structure!
  - More Tests!
  - Redis / NOSQL Mapper
  - Google Maps, Bing Maps!
  - Add an abstraction layer for PostgreSQL earthdistance / PostGIS
  - Declare your mapping service / api keys for Geocoding purposes

# testing

To test, be sure to provide a `config/geo.yml` file with your test environment database configuration, then run the following:

```
GO_ENV=test go test
```

# contributing
  - Fork
  - Create a topic branch
  - Make dem commits!
  - Write dem tests!
  - Submit Pull Request once Tests are Passing
  - do this (づ￣ ³￣)づ