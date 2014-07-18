package stddata

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Currency struct {
	// XMLName		xml.Name	`xml:"CcyNtry"`
	CountryName		string	`xml:"CtryNm"`
	CurrencyName	string	`xml:"CcyNm"`
	CurrencyCode	string	`xml:"Ccy"`
	CurrencyNumber	string	`xml:"CcyNbr"`
	MinorUnits	string	`xml:"CcyMnrUnts"`
}
type Currencies struct {
	XMLName		xml.Name	`xml:"ISO_4217"`
	Currencies	[]Currency	`xml:"CcyTbl>CcyNtry"`
}

func LoadIso4217() {
	res, err := http.Get("http://www.currency-iso.org/dam/downloads/table_a1.xml")
	if err != nil {
		log.Fatal(err)
	}
	currencyBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var c Currencies
	err = xml.Unmarshal([]byte(currencyBody), &c)
	if err != nil {
		fmt.Printf("gads: %s\n", err)
	}

	j, err := json.MarshalIndent(c, "", "  ")
	if err == nil {
		fmt.Printf("%s\n", j)
	} else {
		fmt.Printf("gads: %s\n", err)
	}

	defer res.Body.Close()
}
