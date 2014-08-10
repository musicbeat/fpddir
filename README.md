# [stddata](https://github.com/musicbeat/stddata)

The [stddata](https://github.com/musicbeat/stddata) is a set of components implemented with [golang](https://golang.org) that serve searches of "standard" data sets via http, supplying json responses. The data sets are the Federal Reserve's ACH list, and ISO country, currency, and lanugage codes.

## Getting Started
 * [clone this repo](https://github.com/musicbeat/stddata) - data provider and search components
 * [clone this repo](https://github.com/musicbeat/stddata-cli) - main package with command line
 * go run stddata-cli.go
 * Serves searches at localhost:6060/bank, localhost:6060/country, localhost:6060/currency, and localhost:6060/language


