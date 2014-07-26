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
	"errors"
	"io"
	"log"
	"net/http"
	"sort"
	"strings"
)

// BankProvider implements the Provider interface.
type BankProvider struct {
	loaded   bool
	size     int
	bankIndexes map[string]bankIndex
}

type bankIndex struct {
	bankMap  map[string][]Bank
	bankKeys []string
}

// Bank is the information on one bank in the source data.
type Bank struct {
	RoutingNumber                     string // Length 9; Columns 1-9
	TelegraphicName                   string // Length 18; Columns 10-27
	CustomerName                      string // Length 36; Columns 28-63
	StateAbbreviation                 string // Length 2; Columns 64-65
	City                              string // Length 25; Columns 66-90
	FundsTransferStatus               string // Length 1; Column 91
	FundsSettlementOnlyStatus         string // Length 1; Column 92
	BookEntrySecuritiesTransferStatus string // Length 1; Column 93
	DateOfLastRevision                string // Length 8; Columns 94-101
}

// BankResult is the interface{} that is returned from Search
type BankResult struct {
	banks [][]Bank
}

// Column map:
var rn = [...]int{0, 9}
var tn = [...]int{9, 27}
var cn = [...]int{27, 63}
var sa = [...]int{63, 65}
var ct = [...]int{65, 90}
var ft = [...]int{90, 91}
var fs = [...]int{91, 92}
var be = [...]int{92, 93}
var dt = [...]int{93, 101}

var routingNumberMap map[string][]Bank
var telegraphicNameMap map[string][]Bank
var customerNameMap map[string][]Bank

// Load does the heavy lifting of retrieving the Fed's directory
// of banks, a fixed format text file served via http, and
// populating maps for searches.
func (p *BankProvider) Load() (n int, err error) {
	// Initialize the maps:
	p.bankIndexes = make(map[string]bankIndex)
	routingNumberMap = make(map[string][]Bank)
	telegraphicNameMap = make(map[string][]Bank)
	customerNameMap = make(map[string][]Bank)

	res, err := http.Get("http://www.fededirectory.frb.org/fpddir.txt")
	if err != nil {
		log.Fatal(err)
		return 0, err
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
			log.Fatal(err)
			return 0, err
		}
		sline := strings.TrimRight(string(line), "\n")
		// fmt.Printf("%s\n", sline)
		// fmt.Printf("325280039MAC FCU           MAC FEDERAL CREDIT UNION            AKFT WAINWRIGHT            Y Y20120606\n")
		// fmt.Printf("01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890\n")
		// fmt.Printf("          1         2         3         4         5         6         7         8         9         0\n")

		b.RoutingNumber = strings.TrimSpace(sline[rn[0]:rn[1]])
		b.TelegraphicName = strings.TrimSpace(sline[tn[0]:tn[1]])
		b.CustomerName = strings.TrimSpace(sline[cn[0]:cn[1]])
		b.StateAbbreviation = strings.TrimSpace(sline[sa[0]:sa[1]])
		b.City = strings.TrimSpace(sline[ct[0]:ct[1]])
		b.FundsTransferStatus = strings.TrimSpace(sline[ft[0]:ft[1]])
		b.FundsSettlementOnlyStatus = strings.TrimSpace(sline[fs[0]:fs[1]])
		b.BookEntrySecuritiesTransferStatus = strings.TrimSpace(sline[be[0]:be[1]])
		b.DateOfLastRevision = strings.TrimSpace(sline[dt[0]:dt[1]])

		// add the Bank to the maps:
		routingNumberMap[b.RoutingNumber] = append(routingNumberMap[b.RoutingNumber], b)
		telegraphicNameMap[b.TelegraphicName] = append(telegraphicNameMap[b.TelegraphicName], b)
		customerNameMap[b.CustomerName] = append(customerNameMap[b.CustomerName], b)

	}
	p.loaded = true
	p.size = len(routingNumberMap)
	p.storeData("number", routingNumberMap)
	p.storeData("name", customerNameMap)
	p.storeData("tele", telegraphicNameMap)
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
func (p *BankProvider) Search(s string, q string) (result interface{}, err error) {
	// make sure the data is loaded
	if p.loaded != true {
		return nil, errors.New("this should be a 503 Service Unavailable by the time it gets to the client")
	}
	bi, found := p.bankIndexes[s]
	if !found {
		// search cannot be performed
		return nil, errors.New("this should be a 400 Bad Request by the time it gets to the client")
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
		if strings.EqualFold(q, bi.bankKeys[k][0:len(q)]) {
			tmp[i] = bi.bankMap[bi.bankKeys[k]]
			i++
		}
	}
	res.banks = tmp[0:i]
	return res
}
