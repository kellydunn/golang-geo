OB# golang-geo roadmap

## [1.0.0] Desired Featureset

  <a name = "namespace"/>
  - **Cleaner top-level namespace** It seems rather messy to have the entire impelentation of golang-geo to be in the same root-level directory.  It seems like it would be better if users could import "github.com/kellydunn/golang-geo/sql" for sql funcitonality.

  <a name = "geocoders" />
  - **Extract geocoder implementations into seperate libraries** In the future, golang-geo's responsibilities will not be to handle API implementation logic.  This is better defined and contained in seperate implementations as they will be easier to maintain than to contribute upstream to golang-geo core.

  <a name = "api-names" />
  - **Rename API methods to be the exact name of the corresponding mathematical functions** The current implementation hides domain knowledge by convoluting the name of the API methods and what they actually do.

  <a name = "point-namespace" />
  - **All point logic should be within a `point` namespace level and not struct methods of points** It feels rather awkward to calculate haversine distances by issuing a struct-level method.

  <a name = "exported" />
  - **Come to a consensus on exported vs unexported fields for Points and Polygons** Admittedly, the implementation of geo.Point has been influenced by traditional OOP visibility principles and does not mix well with golang's notion of exported and unexported fields.  There are some advantages and disadvantages of both approaches, but it is currently unclear if clients should have direct visibility to Point and Polygon internals. 