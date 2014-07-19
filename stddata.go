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

Other standards for future work have been identified, including:
	US Postal Codes and Cities
	... to be seen
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

// Server is the interface that wraps the Serve method.
//
// Serve sets up the HTTP server endpoint for search.
// It uses the supplied port. It returns the full
// search url. If an error occurs, it returns that as well.
type Server interface {
	Serve(port string) (err error)
}

// Searcher is the interface that wraps the Search method.
//
// Search takes the query string, attempts to find a match,
// and returns a JSON result. If an error occurs, it returns 
// that as well.
type Searcher interface {
	Search(q string) (r string, err error)
}



