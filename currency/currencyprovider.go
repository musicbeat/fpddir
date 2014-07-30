// Copyright 2014 Musicbeat.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package currency implements the methods of a stddata.Provider.
It provides searches against the data set retrieved from
currency-iso.org.
*/
package currency

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
)

// CurrencyProvider implements the Provider interface.
type CurrencyProvider struct {
	loaded bool
	size int
	currencyIndexes map[string]currencyIndex
}

type currencyIndex struct {
	currencyMap map[string][]Currency
	currencyKeys []string
}

// Currency is the information on one currency in the source data.
type Currency struct {
	// XMLName		xml.Name	`xml:"CcyNtry"`
	CountryName		string	`xml:"CtryNm"`
	CurrencyName	string	`xml:"CcyNm"`
	CurrencyCode	string	`xml:"Ccy"`
	CurrencyNumber	string	`xml:"CcyNbr"`
	MinorUnits	string	`xml:"CcyMnrUnts"`
}
type Currencies struct {
	// XMLName		xml.Name	`xml:"ISO_4217"`
	Currencies	[]Currency	`xml:"CcyTbl>CcyNtry"`
}

// CurrencyResult is the interface{} that is returned from Search
type CurrencyResult struct {
	Currencies [][]Currency
}

var countryNameMap map[string][]Currency
var currencyNameMap map[string][]Currency
var currencyCodeMap map[string][]Currency
var currencyNumberMap map[string][]Currency

// Load does the heavy lifting of retrieving the iso.org
// web site's handy XML file. The file is retrieved and
// parsed into structs, and loaded into maps and indexes
// to support searches.
func (p *CurrencyProvider) Load() (n int, err error) {
	// Initialize the maps:
	p.currencyIndexes = make(map[string]currencyIndex)
	countryNameMap = make(map[string][]Currency)
	currencyNameMap = make(map[string][]Currency)
	currencyCodeMap = make(map[string][]Currency)
	currencyNumberMap = make(map[string][]Currency)

	res, err := http.Get("http://www.currency-iso.org/dam/downloads/table_a1.xml")
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	defer res.Body.Close()

	currencyBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	var currencies Currencies
	err = xml.Unmarshal([]byte(currencyBody), &currencies)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	// add the currency entities to the maps:
	for _, c := range currencies.Currencies {
		countryNameMap[c.CountryName] = append(countryNameMap[c.CountryName], c)
		currencyNameMap[c.CurrencyName] = append(currencyNameMap[c.CurrencyName], c)
		currencyCodeMap[c.CurrencyCode] = append(currencyCodeMap[c.CurrencyCode], c)
		currencyNumberMap[c.CurrencyNumber] = append(currencyNumberMap[c.CurrencyNumber], c)
	}

	p.storeData("country", countryNameMap)
	p.storeData("name", currencyNameMap)
	p.storeData("code", currencyCodeMap)
	p.storeData("number", currencyNumberMap)
	p.size = len(currencyCodeMap)
	p.loaded = true
	return len(currencyCodeMap), err
}

func (p *CurrencyProvider) storeData(s string, m map[string][]Currency) {
	// store the map
	var ci currencyIndex
	ci.currencyMap = m
	// extract the keys
	ci.currencyKeys = make([]string, len(m))
	i := 0
	for k, _ := range m {
		ci.currencyKeys[i] = k
		i++
	}
	// sort the keys
	sort.Strings(ci.currencyKeys)
	p.currencyIndexes[s] = ci
}

// Search returns a collection as an interface{} and error. The collection
// contains an array of the results to the search. The value
// in index is used to choose the map of Currency entities that will be searched.
// If the value in index does not match the name of a map, an error is returned.
// The keys in the map specified by index are searched using a regex-like 'q.*', and
// any matching Currency entities are returned in the result.
func (p *CurrencyProvider) Search(s string, q string) (result interface{}, err error) {
	// make sure the data is loaded
	if p.loaded != true {
		return nil, errors.New("this should be a 503 Service Unavailable by the time it gets to the client")
	}
	ci, found := p.currencyIndexes[s]
	if !found {
		// search cannot be performed
		return nil, errors.New("this should be a 400 Bad Request by the time it gets to the client")
	}
	result = doSearch(ci, q)
	return result, nil
}
func doSearch(ci currencyIndex, q string) (res CurrencyResult) {
	// prepare the response. allocate enough space for the response to be the 
	// entire data set.
	tmp := make([][]Currency, len(ci.currencyKeys))
	// brute force the sorted list of keys, looking for a match to 'q.*'.
	// add each match to the result array. The results are added in the
	// order of the sorted keys, so the results are sorted.
	i := 0
	for k := range ci.currencyKeys {
		if len(ci.currencyKeys[k]) >= len(q) {
			if strings.EqualFold(q, ci.currencyKeys[k][0:len(q)]) {
				tmp[i] = ci.currencyMap[ci.currencyKeys[k]]
				i++
			}
		}
	}
	res.Currencies = tmp[0:i]
	return res
}
