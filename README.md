# WeGo
____

WeGo is an experimental, tiny MVC for creating APIs. Some features of WeGo:
  - Easy to define routes (routes.json)
  - Ability to serve static or dynamic responses
  - Easily deployable


### Routing
Routing is configured through a JSON file (routes.json), with several options available for customization:
```json
{
  "path" : "/fetch",
  "method" : "GET",
  "mode" : "static",
  "file" : "fetch.json"
}
```

### Test Driving

Installing is easy:
```sh 
$ go get github.com/drewlong/wego
$ go build wego
```
You'll want your directory tree to look like this:
```
MyApp/
- wego
- routes.json
- templates/
-- test.json (example)
-- fetch.json (example)
```

To run:
```
$ ./wego -p 8080
```
Wego defaults to port 7000, but you can set the port with the -p flag. 

### Questions, Ideas?

Feel free to contact me, fork as you'd like. Enjoy!
