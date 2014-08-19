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

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// Provider is the interface for a Standard Data Provider. A type that 
// implements Provider's methods can be managed as a Standard Data Provider.
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

// Service is used to handle http access to the stddata providers' data.
type Service struct {
	Provider   Provider
	Count      int
	EntityName string
}

// LoadProvider is used to prepare the data.
//
// Implementations retrieve their source data, and index it for
// searching.
func (s *Service) LoadProvider(p Provider, e string) (err error) {
	s.Provider = p
	s.EntityName = e
	n, err := s.Provider.Load()
	if err != nil {
		log.Printf("Provider for %s failed to load. %s\n", e, err)
		return errors.New("Searches will get 503 Service Unavailable for this provider")
	}
	s.Count = n
	return nil
}

// ServeHTTP is the Service's implementation for searching.
//
// After some basic validation of the search request, the
// Provider's Search() implementation is called. The response
// from Search() is marshalled into json.
func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// get the index and query values
	var index, query string
	var err error
	if index, query, err = getQuery(r.URL.RawQuery); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := s.Provider.Search(index, query)
	if err != nil {
		if serr, ok := err.(*ServiceError); ok {
			w.WriteHeader(serr.Code)
			return
		}
		log.Printf("Error %v\n", err)
		return
	}
	// convert result to json
	j, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	io.WriteString(w, fmt.Sprintf("%s\n", j))
}

// Get the "index=query" parts of the request, for example, "name=Abc".
// Or for a dump of an index, "name=_dump". Or error.
func getQuery(u string) (query string, index string, err error) {
	v := strings.Split(u, "=")
	if len(v) < 2 {
		err := errors.New("Malformed request")
		return index, query, err
	}
	index = v[0]
	if len(index) < 1 {
		err := errors.New("Malformed request")
		return index, query, err
	}
	query = v[1]
	if len(query) < 1 {
		err := errors.New("Malformed request")
		return index, query, err
	}
	return index, query, err
}
// ServiceError combines an http status code and an
// application error message.
type ServiceError struct {
	Msg  string // description of error
	Code int    // http status constant
}

// Error implements the built-in error interface on ServiceError.
func (e *ServiceError) Error() string {
	return e.Msg
}
