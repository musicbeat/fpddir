package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Bank struct {
	RoutingNumber string // Length 9; Columns 1-9
	TelegraphicName string // Length 18; Columns 10-27
	CustomerName string // Length 36; Columns 28-63
	StateAbbreviation string // Length 2; Columns 64-65
	City string // Length 25; Columns 66-90
	FundsTransferStatus string // Length 1; Column 91
	FundsSettlementOnlyStatus string // Length 1; Column 92
	BookEntrySecuritiesTransferStatus string // Length 1; Column 93
	DateOfLastRevision string // Length 8; Columns 94-101
}
var rn = [...]int {0, 9}
var tn = [...]int {9, 27}
var cn = [...]int {27, 63}
var sa = [...]int {63, 65}
var ct = [...]int {65, 90}
var ft = [...]int {90, 91}
var fs = [...]int {91, 92}
var be = [...]int {92, 93}
var dt = [...]int {93, 101}

func main() {
	res, err := http.Get("http://www.fededirectory.frb.org/fpddir.txt")
	if err != nil {
		log.Fatal(err)
	}
	bio := bufio.NewReader(res.Body)
	for {
		line, err := bio.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		sline := strings.TrimRight(string(line), "\n")
		// fmt.Printf("%s\n", sline)
		//         "325280039MAC FCU           MAC FEDERAL CREDIT UNION            AKFT WAINWRIGHT            Y Y20120606\n"
		// fmt.Printf("01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890\n")
		// fmt.Printf("          1         2         3         4         5         6         7         8         9         0\n")
		var b Bank
		b.RoutingNumber = strings.TrimSpace(sline[rn[0]:rn[1]])
		b.TelegraphicName = strings.TrimSpace(sline[tn[0]:tn[1]])
		b.CustomerName = strings.TrimSpace(sline[cn[0]:cn[1]])
		b.StateAbbreviation = strings.TrimSpace(sline[sa[0]:sa[1]])
		b.City = strings.TrimSpace(sline[ct[0]:ct[1]])
		b.FundsTransferStatus = strings.TrimSpace(sline[ft[0]:ft[1]])
		b.FundsSettlementOnlyStatus = strings.TrimSpace(sline[fs[0]:fs[1]])
		b.BookEntrySecuritiesTransferStatus = strings.TrimSpace(sline[be[0]:be[1]])
		b.DateOfLastRevision = strings.TrimSpace(sline[dt[0]:dt[1]])

		j, err := json.MarshalIndent(b, "", "  ")
		if err == nil {
			fmt.Printf("%s\n", j)
		} else {
			fmt.Printf("gads: %s\n", err)
		}
	}
	defer res.Body.Close()
}
