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


```

# what 

A set of convience interfaces and methods that makes geo-related calculations easier for Go.

Also just an simple experiement for me to play around with in Go.

# usage

Import from github, and get geomancin'

```
import( _ "github.com/kellydunn/golang-geo")

func main() {
     // Read below for more information on how to configure your SQL setup.
     db, err := geo.HandleWithSql()

     ...

     // Find all of the points of interest that are in a 5km radius of [42.333, 121,111]
     // You could also probably use PostgreSQL's built-in earth distance module :P 
     // http://www.postgresql.org/docs/8.3/static/earthdistance.html
     p := &Point{lat: 42.3333, lng: 121.111}
     res, _ := db.PointsWithinRadius(p, 5)

     ...

     // You can also find a point after transposing another a certain distance(km) with a certain bearing(degrees)
     p2 := p.PointAtDistanceAndBearing(7.9, 45)
     
     // Inspect the point!
     fmt.Printf("LAT: %f\n", p2.lat)
     fmt.Printf("LNG: %f\n", p2.lng)

     ...


     // You can also find the GreatCircleDistance Distance between two points
     distance := p.GreatCircleDistance(p2)

     ...
}
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
  - Redis / NOSQL Mapper
  - Add an abstraction layer for PostgreSQL earthdistance / PostGIS
  - Declare your mapping service / api keys for Geocoding purposes

