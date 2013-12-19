GoValidateJson
==============

Convert a JSON request to a map, then validate and convert the request to a defined model (struct). GoValidateJson makes validating your JSON requests, as well as accessing the data afterwards extremely easy. 

Note: While this can be used to convert json to structs, the idea is this will fail if the struct can't be filled instead of mapping a default value to the field. 

####Run:
```
  go test -v 
```
to view the results after running.

####Struct Definition Example:

```
type Model struct {
  Name  string  `request:"name"`
  Email string  `request:"password"`
  Age   float64 `request:"age"`
}
```

You must define a "request" tag on each field of your struct, as this is used to get the JSON key from the map. 

#### Types
As you may know JSON only has 5 main types: String, Float, Bool, List, and Object. In our case Float is mapped to float64, String to string, and Bool to bool. If we hit a map we repeat the same process and run again. Lists are still being implemented
