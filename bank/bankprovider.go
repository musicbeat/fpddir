// Copyright 2014 Musicbeat.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package bank implements the methods of a stddata.Provider.
It provides searches against the data set retrieved from
the Federal Reserve E-Payments Routing Directory.
*/
package bank

import (
	"bufio"
	"io"
	"net/http"
	"sort"
	"strings"

	"github.com/musicbeat/stddata"
)

// BankProvider implements the Provider interface.
type BankProvider struct {
	loaded      bool
	size        int
	bankIndexes map[string]bankIndex
}

type bankIndex struct {
	bankMap  map[string][]Bank
	bankKeys []string
}

// Bank is the information on one bank in the source data.
type Bank struct {
	Routing               string // Length 9; Columns 1-9
	OfficeCode            string // Length 1; Columns 10
	ServicingFRBNumber    string // Length 9; Columns 11-19
	RecordTypeCode        string // Length 1; Columns 20
	ChangeDate            string // Length 6; Columns 21-26
	NewRoutingNumber      string // Length 9; Columns 27-35
	CustomerName          string // Length 36; Columns 36-71
	Address               string // Length 36; Columns 72-107
	City                  string // Length 20; Columns 108-127
	StateCode             string // Length 2; Columns 128-129
	Zipcode               string // Length 5; Columns 130-134
	ZipcodeExtension      string // Length 4; Columns 135-138
	TelephoneAreaCode     string // Length 3; Columns 139-141
	TelephonePrefixNumber string // Length 3; Columns 142-144
	TelephoneSuffixNumber string // Length 4; Columns 145-148
	InstitutionStatusCode string // Length 1; Columns 149
	DataViewCode          string // Length 1; Columns 150
}

// BankResult is the interface{} that is returned from Search
type BankResult struct {
	Banks [][]Bank
}

// Column map:
var rn = [...]int{0, 9}
var oc = [...]int{9, 10}
var sf = [...]int{10, 19}
var rt = [...]int{19, 20}
var cd = [...]int{20, 26}
var nr = [...]int{26, 35}
var cn = [...]int{36, 71}
var ad = [...]int{71, 107}
var ci = [...]int{107, 127}
var sc = [...]int{127, 129}
var zc = [...]int{129, 134}
var z4 = [...]int{134, 138}
var ac = [...]int{138, 141}
var tp = [...]int{141, 144}
var ts = [...]int{144, 148}
var is = [...]int{148, 149}
var dv = [...]int{149, 150}

var routingNumberMap map[string][]Bank
var customerNameMap map[string][]Bank

var fedurl = "http://www.fededirectory.frb.org/FedACHdir.txt"

// Load does the heavy lifting of retrieving the Fed's directory
// of banks, a fixed format text file served via http, and
// populating maps for searches.
func (p *BankProvider) Load() (n int, err error) {
	// Initialize the maps:
	p.bankIndexes = make(map[string]bankIndex)
	routingNumberMap = make(map[string][]Bank)
	customerNameMap = make(map[string][]Bank)

	res, err := http.Get(fedurl)
	if err != nil {
		msg := "Failed to retrieve " + fedurl + ". " + err.Error()
		return 0, &stddata.ServiceError{msg, http.StatusServiceUnavailable}
	}
	defer res.Body.Close()

	bio := bufio.NewReader(res.Body)
	for {
		var b Bank
		line, err := bio.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, &stddata.ServiceError{err.Error(), http.StatusServiceUnavailable}
		}
		sline := strings.TrimRight(string(line), "\n")

		b.Routing = strings.TrimSpace(sline[rn[0]:rn[1]])
		b.OfficeCode = strings.TrimSpace(sline[oc[0]:oc[1]])
		b.ServicingFRBNumber = strings.TrimSpace(sline[sf[0]:sf[1]])
		b.RecordTypeCode = strings.TrimSpace(sline[rt[0]:rt[1]])
		b.ChangeDate = strings.TrimSpace(sline[cd[0]:cd[1]])
		b.NewRoutingNumber = strings.TrimSpace(sline[nr[0]:nr[1]])
		b.CustomerName = strings.TrimSpace(sline[cn[0]:cn[1]])
		b.Address = strings.TrimSpace(sline[ad[0]:ad[1]])
		b.City = strings.TrimSpace(sline[ci[0]:ci[1]])
		b.StateCode = strings.TrimSpace(sline[sc[0]:sc[1]])
		b.Zipcode = strings.TrimSpace(sline[zc[0]:zc[1]])
		b.ZipcodeExtension = strings.TrimSpace(sline[z4[0]:z4[1]])
		b.TelephoneAreaCode = strings.TrimSpace(sline[ac[0]:ac[1]])
		b.TelephonePrefixNumber = strings.TrimSpace(sline[tp[0]:tp[1]])
		b.TelephoneSuffixNumber = strings.TrimSpace(sline[ts[0]:ts[1]])
		b.InstitutionStatusCode = strings.TrimSpace(sline[is[0]:is[1]])
		b.DataViewCode = strings.TrimSpace(sline[dv[0]:dv[1]])

		// add the Bank to the maps:
		routingNumberMap[b.Routing] = append(routingNumberMap[b.Routing], b)
		customerNameMap[b.CustomerName] = append(customerNameMap[b.CustomerName], b)

	}
	p.storeData("number", routingNumberMap)
	p.storeData("name", customerNameMap)
	p.size = len(routingNumberMap)
	p.loaded = true
	return len(routingNumberMap), err
}

func (p *BankProvider) storeData(s string, m map[string][]Bank) {
	// store the map
	var bi bankIndex
	bi.bankMap = m
	// extract the keys
	bi.bankKeys = make([]string, len(m))
	i := 0
	for k, _ := range m {
		bi.bankKeys[i] = k
		i++
	}
	// sort the keys
	sort.Strings(bi.bankKeys)
	// add to bankIndexes
	p.bankIndexes[s] = bi
}

// Search returns a collection as an interface{} and error. The collection
// contains an array of the results to the search. The value
// in index is used to choose the map of Bank entities that will be searched.
// If the value in index does not match the name of a map, an error is returned.
// The keys in the map specified by index are searched using a regex-like 'q.*', and
// any matching Banks are returned in the result.
func (p *BankProvider) Search(index string, q string) (result interface{}, err error) {
	// make sure the data is loaded
	if p.loaded != true {
		return 0, &stddata.ServiceError{err.Error(), http.StatusServiceUnavailable}
	}
	bi, found := p.bankIndexes[index]
	if !found {
		// search cannot be performed
		msg := "No index on " + index
		return nil, &stddata.ServiceError{msg, http.StatusBadRequest}
	}
	result = doSearch(bi, q)
	return result, nil
}
func doSearch(bi bankIndex, q string) (res BankResult) {
	// prepare the response. allocate enough space for the response to be the
	// entire data set.
	tmp := make([][]Bank, len(bi.bankKeys))
	// brute force the sorted list of keys, looking for a match to 'q.*'.
	// add each match to the result array. The results are added in the
	// order of the sorted keys, so the results are sorted.
	i := 0
	for k := range bi.bankKeys {
		if len(bi.bankKeys[k]) >= len(q) {
			if strings.EqualFold(q, bi.bankKeys[k][0:len(q)]) {
				tmp[i] = bi.bankMap[bi.bankKeys[k]]
				i++
			}
		}
	}
	res.Banks = tmp[0:i]
	return res
}
