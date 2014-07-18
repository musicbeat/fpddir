/*
data loader component for ISO 3166-2 country codes.
source data is declared in iso3166data.go
*/
package country

import (
	"encoding/csv"
	"io"
	"log"
)

type Country struct {
	EnglishName	string
	Alpha2Code	string
	Alpha3Code				string
	NumericCode			string
}
type Countries struct {
	Countries []Country
}

var englishNameMap map[string]Country
var alpha2Map map[string]Country
var alpha3Map map[string]Country
var numericMap map[string]Country

// Load implements the Loader interface
func (c Country) Load() (n int, err error) {
	// initialize the maps:
	englishNameMap = make(map[string]Country)
	alpha2Map = make(map[string]Country)
	alpha3Map = make(map[string]Country)
	numericMap = make(map[string]Country)


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
		englishNameMap[c.EnglishName] = c
		alpha2Map[c.Alpha2Code] = c
		alpha3Map[c.Alpha3Code] = c
		numericMap[c.NumericCode] = c

	}
	return len(englishNameMap), err
}
