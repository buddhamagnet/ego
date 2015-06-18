### EGO

This repository is based on a prototype I built at work designed to decouple
our back end from a set of standalone front ends, and to allow us to plug in
additional back ends.

The API is wired up by reading a RAML (REST API Modelling Language) file
using a package I built.

#### SERVICE SETUP

1. Clone this repository into your GOPATH.
2. Run `godep restore` to get the minimal dependencies.
3. Run `make build` to drop the binary into the build folder.
4. Copy samples/.env into the build folder.
5. Run `build/ego` and hit localhost:9494.
6. Run `build/ego --port=<port>``` if you want another port.

### MAKEFILE

The Makefile provides the following:

* `build`: build the application binary and drop it in the build folder.
* `gomkbuild`: build the application binary.
* `gomkxbuild`: build all cross-platform binaries, using `gox`.
* `gomkclean`: clean the project directory.
* `vet`: run `go tool vet` on each source file.
* `lint`: run `golint` on each source file.
* `fmt`: run `go fmt` on the entire project.
* `test`: run `go test` for all packages in the project.
* `race`: run `go test` with race detection in all packages in the project.
* `cover`: run tests with coverage report in all pkgs in the project.
* `printvars`: print all variables defined in the Makefile.

### BINARIES

Run `pride` to get nicely colorized test output.

### SUGGESTIONS

Use [goconvey](http://goconvey.co) for browser-based test feedback.

#### FRONT END SETUP

This repository contains a small front end written in angular.js.

* ```cd``` into the frontend folder and run `make build`.
* Ensure the content service is running and run 'www/app'.
* Navigate to localhost:9494 and feel the love.
