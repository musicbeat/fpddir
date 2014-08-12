# [stddata](https://github.com/musicbeat/stddata)

The [stddata](https://github.com/musicbeat/stddata) is a set of components implemented with [golang](https://golang.org) that serve searches of "standard" data sets via http, supplying json responses. The data sets are the Federal Reserve's ACH list, and ISO country, currency, and lanugage codes.

Everything here is a proof-of-concept. Nothing is ready for production use. And the code could benefit from considerably more refactoring to eliminate redundancy, improve error handling, and generally become more idiomatic [go](https://golang.org) code.

## Ingredients
 * [stddata](https://github.com/musicbeat/stddata) - this package, which comprises the data providers.
 * [stddata-cli](https://github.com/musicbeat/stddata-cli) - the ```main``` package, used to launch the http server.
 * [stddata-build](https://github.com/musicbeat/stddata-build) - some components that construct a [docker](https://docker.com)-ized deployment of the server.

## Getting Started
 * [clone this repo](https://github.com/musicbeat/stddata) - data provider and search components
 * [clone this repo](https://github.com/musicbeat/stddata-cli) - main package with command line
 * go run stddata-cli.go
 * Serves searches at localhost:6060/bank, localhost:6060/country, localhost:6060/currency, and localhost:6060/language

## More
 * Check out [stddata-build](https://github.com/musicbeat/stddata-build) to explore the use of [docker](https://docker.com) with the stddata server.
