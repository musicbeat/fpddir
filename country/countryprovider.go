// Copyright 2014 Musicbeat.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package country implenments the methods of a stddata.Provider.
It provides searches against the data set of ISO 3166-2
country codes. Source data is declared in iso3166data.go
*/
package country

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"sort"
	"strings"
)

// CountryProvider implements the Provider interface.
type CountryProvider struct {
	loaded bool
	size int
	countryIndexes map[string]countryIndex
}

type countryIndex struct {
	countryMap map[string][]Country
	countryKeys []string
}

type Country struct {
	EnglishName	string
	Alpha2Code	string
	Alpha3Code				string
	NumericCode			string
}

// CountryResult is the interface{} that is returned from Search
type CountryResult struct {
	countries [][]Country
}

var englishNameMap map[string][]Country
var alpha2Map map[string][]Country
var alpha3Map map[string][]Country
var numericMap map[string][]Country

// Load implements the Loader interface
func (p *CountryProvider) Load() (n int, err error) {
	// initialize the maps:
	p.countryIndexes = make(map[string]countryIndex)
	englishNameMap = make(map[string][]Country)
	alpha2Map = make(map[string][]Country)
	alpha3Map = make(map[string][]Country)
	numericMap = make(map[string][]Country)


	reader := csv.NewReader(countrydata)
	reader.Comma = '\t'
	reader.FieldsPerRecord = 4
	reader.TrimLeadingSpace = true

	for {
		// read just one record, but we could ReadAll() as well
		record, err := reader.Read()
		// end-of-file is fitted into err
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
			return 0, err
		}

		var c Country
		c.EnglishName = record[0]
		c.Alpha2Code = record[1]
		c.Alpha3Code = record[2]
		c.NumericCode = record[3]

		// add the Country to the maps
		englishNameMap[c.EnglishName] = append(englishNameMap[c.EnglishName], c)
		alpha2Map[c.Alpha2Code] = append(alpha2Map[c.Alpha2Code], c)
		alpha3Map[c.Alpha3Code] = append(alpha3Map[c.Alpha3Code], c)
		numericMap[c.NumericCode] = append(numericMap[c.NumericCode], c)

	}
	p.loaded = true
	p.size = len(englishNameMap)
	p.storeData("name", englishNameMap)
	p.storeData("alpha2", alpha2Map)
	p.storeData("alpha3", alpha3Map)
	p.storeData("number", numericMap)
	return len(englishNameMap), err
}

func (p *CountryProvider) storeData(s string, m map[string][]Country) {
	// store the map
	var ci countryIndex
	ci.countryMap = m
	// extract the keys
	ci.countryKeys = make([]string, len(m))
	i := 0
	for k, _ := range m {
		ci.countryKeys[i] = k
		i++
	}
	// sort the keys
	sort.Strings(ci.countryKeys)
	// add to countryIndexes
	p.countryIndexes[s] = ci
}

// Search returns a collection as an interface{} and error. The collection
// contains an array of the results to the search. The value
// in index is used to choose the map of Country entities that will be searched.
// If the value in index does not match the name of a map, an error is returned.
// The keys in the map specified by index are searched using a regex-like 'q.*', and
// any matching Countries are returned in the result.
func (p *CountryProvider) Search(s string, q string) (result interface{}, err error) {
	// make sure the data is loaded
	if p.loaded != true {
		return nil, errors.New("this should be a 503 Service Unavailable by the time it gets to the client")
	}
	ci, found := p.countryIndexes[s]
	if !found {
		// search cannot be performed
		return nil, errors.New("this should be a 400 Bad Request by the time it gets to the client")
	}
	result = doSearch(ci, q)
	return result, nil
}
func doSearch(ci countryIndex, q string) (res CountryResult) {
	// prepare the response. allocate enough space for the response to be the 
	// entire data set.
	tmp := make([][]Country, len(ci.countryKeys))
	// brute force the sorted list of keys, looking for a match to 'q.*'.
	// add each match to the result array. The results are added in the
	// order of the sorted keys, so the results are sorted.
	i := 0
	for k := range ci.countryKeys {
		if strings.EqualFold(q, ci.countryKeys[k][0:len(q)]) {
			tmp[i] = ci.countryMap[ci.countryKeys[k]]
			i++
		}
	}
	res.countries = tmp[0:i]
	return res
}
