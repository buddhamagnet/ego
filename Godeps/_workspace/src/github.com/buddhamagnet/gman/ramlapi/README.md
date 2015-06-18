###RAML API

####SAMPLE USAGE:
```go
package main

import (

	"net/http"
	"github.com/buddhamagnet/gman/ramlapi"
	"github.com/gorilla/mux"
)

var router *mux.Router

func main() {
    router = mux.NewRouter().StrictSlash(true)
    api, _ := ramlapi.ProcessRAML("api.raml")
    ramlapi.Build(api, routerFunc)
    log.Fatal(http.ListenAndServe(port, nil))
}

func routerFunc(verb, resourcepath, handler string) {
    router.
        Methods(verb).
        Path(resourcepath).
        Handler(RouteMap[handler])
}
```
