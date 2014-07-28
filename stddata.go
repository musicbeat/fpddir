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
		A handy, fixed format text file available at the Fed's website
	stddata/country - ISO 3166-1 Country Codes (Officially Assigned)
		ISO charges for access to this information through their website, but
		Wikipedia has a table of these codes. A data set was extracted from
		Wikipedia's website for use in this package.
	stddata/currency - ISO 4217 Currency Codes
		A handy xml document available from iso.org's website.
	stddata/language - ISO 639 Language Codes
		A handy, pipe-delimited csv file.

*/
package stddata

// Provider is the interface that wraps all the interfaces
// together. A type that implements Provider's methods
// can be managed as a Standard Data Provider.
type Provider interface {
	// Load loads the data, according to the needs of the particular
	// implementation's standard data set. It returns the number of
	// items it has loaded. If an error occurs, it returns that as well.
	Load() (n int, err error)
	// Search takes the name of the index to be searched, and the value
	// to match in that index. It returns an interface and an error.
	// The value that is returned as v is intended to be marshaled as 
	// json -- it is expected to be the collection of entities that 
	// match the search.
	Search(index string, q string) (v interface{}, err error)
}

