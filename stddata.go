// Copyright 2014 Musicbeat.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package stddata provides a simple mechanism for making
information about "standard" data sets available
for lookups and queries. The standards are:
	Federal Reserve E-Payments Routing Directory
	ISO 639 Language Codes
	ISO 4217 Currency Codes
	ISO 3166-1 Country Codes (Officially Assigned)

Packages

The pieces of stddata's implementation are organized
as follows:
	stddata - interfaces, types, and functions for managing the data providers
	stddata/bank - Federal Reserve E-Payments Routing Directory
	stddata/country - ISO 3166-1 Country Codes (Officially Assigned)
	stddata/currency - ISO 4217 Currency Codes
	stddata/language - ISO 639 Language Codes

*/
package stddata

// Loader is the interface that wraps the Load method.
//
// Load loads the data, according to the needs of the particular
// implementation's standard data set. It returns the number of
// items it has loaded. If an error occurs, it returns that as well.
type Loader interface {
	Load() (n int, err error)
}

type Registration struct {
}
// Registrar is the interface that wraps the Register method.
//
// Register returns the Registration information for a
// provider. The Registration information allows the caller
// to set up the endpoints for serving the provider.
type Registrar interface {
	Register (r Registration, err error)
}

// Searcher is the interface that wraps the Search method.
//
// Search takes the query string, attempts to find a match,
// and returns an interface as the result. If an error occurs,
// it returns that as well. The value that is returned as v
// is intended to be marshaled as json -- it is expected to 
// be the collection of entities that match the search.
type Searcher interface {
	Search(q string) (v interface{}, err error)
}

// Provider is the interface that wraps all these parts
// together. A type that implements Provider's methods
// can be managed as a Standard Data Provider.
type Provider interface {
	Loader
	Registrar
	Searcher
}

