package currency

import (
	"encoding/xml"
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
	// XMLName		xml.Name	`xml:"ISO_4217"`
	Currencies	[]Currency	`xml:"CcyTbl>CcyNtry"`
}

var countryNameMap map[string][]Currency
var currencyNameMap map[string][]Currency
var currencyCodeMap map[string][]Currency
var currencyNumberMap map[string][]Currency

func (c Currency) Load() (n int, err error) {
	// Initialize the maps:
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
	
	for _, c := range currencies.Currencies {
		countryNameMap[c.CountryName] = append(countryNameMap[c.CountryName], c)
		currencyNameMap[c.CurrencyName] = append(currencyNameMap[c.CurrencyName], c)
		currencyCodeMap[c.CurrencyCode] = append(currencyCodeMap[c.CurrencyCode], c)
		currencyNumberMap[c.CurrencyNumber] = append(currencyNumberMap[c.CurrencyNumber], c)
	}

	return len(currencyCodeMap), err
}

