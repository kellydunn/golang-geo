# [0.1.0](https://github.com/kellydunn/golang-geo/tree/v0.0.2) January 25, 2014

  - Introducing `geo.NewSQLMapper`, which creates and returns a pointer to a new `geo.SQLMapper`.  This solved issues where users had to create extraneous `*sql.DB` in order to perform Mapper operations.  The introduction of this method signature marks `geo.HandleWithSQL` a canidate for removal in a Major Patch verison.
  - Introducing `geo.SQLMapper.SqlDBConn`, which returns the database connection of the `geo.SQLMapper` for inspection purposes.
  - Introduces `geo.GetSQLConfFromFile`, which accepts the pathname to the desired configuration file.  This was necessary in order for `NewSQLMapper` to allow users to supply different pathnames to create new `SQLMapper`s.
  - Various Bugfixes including some mismatched `EARTH_RADIUS` logic in querying points in SQL databases.

# [0.0.2](https://github.com/kellydunn/golang-geo/tree/v0.0.2) November 28, 2013

  - Change `EARTH_RADIUS` to comply with the information published in [wikipedia](http://en.wikipedia.org/wiki/Earth_radius).
  - Added some more documentation to publicly available structs and methods.

# [0.0.1](https://github.com/kellydunn/golang-geo/tree/v0.0.1) November 28, 2013

  - First tagged release.
