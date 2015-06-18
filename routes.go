package main

import (
	"github.com/buddhamagnet/gman/ramlapi"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var router *mux.Router

var RouteSet = map[string]string{
	"blog listings": "/content/blogs/{blog_path}",
}

var RouteMap = map[string]http.HandlerFunc{
	"RootHandler":    RootHandler,
	"ContentHandler": ContentHandler,
}

// NewRouter creates a mux.Router and registers routes with it.
func NewRouter() *mux.Router {
	router = mux.NewRouter().StrictSlash(true)
	assembleMiddleware(router)
	api, err := ramlapi.ProcessRAML(os.Getenv("RAML_FILE_PATH"))
	if err != nil {
		log.Fatal(err)
	}
	ramlapi.Build(api, routerFunc)
	return router
}

// assembleMiddleware sets up the middleware stack for gref.
func assembleMiddleware(r *mux.Router) {
	http.Handle("/",
		JsonMiddleware(
			LoggingMiddleware(
				RecoverMiddleware(r))))
}

func routerFunc(verb, resourcepath, handler string) {
	router.
		Methods(verb).
		Path(resourcepath).
		Handler(RouteMap[handler])
}
