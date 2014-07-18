/*
data loader component for ISO 3166-2 country codes.
source data is declared in iso3166data.go
*/
package stddata

import (
	"encoding/json"
	"encoding/csv"
	"fmt"
	"io"
	"sort"
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

var CountryMaps map[string] map[string]Country

func LoadIso3166() {
	var englishNameMap map[string]Country
	var alpha2Map map[string]Country
	var alpha3Map map[string]Country
	var numericMap map[string]Country
	// initialize the maps:
	englishNameMap = make(map[string]Country)
	alpha2Map = make(map[string]Country)
	alpha3Map = make(map[string]Country)
	numericMap = make(map[string]Country)
	CountryMaps = make(map[string] map[string]Country)

	CountryMaps["Alpha2"] = alpha2Map
	CountryMaps["Alpha3"] = alpha3Map
	CountryMaps["Numeric"] = numericMap
	CountryMaps["EnglishName"] = englishNameMap

	reader := csv.NewReader(iso3166data)
	reader.Comma = '\t'
	reader.FieldsPerRecord = 4
	reader.TrimLeadingSpace = true

	var countries Countries
	for {
		// read just one record, but we could ReadAll() as well
		record, err := reader.Read()
		// end-of-file is fitted into err
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}

		var c Country
		c.EnglishName = record[0]
		c.Alpha2Code = record[1]
		c.Alpha3Code = record[2]
		c.NumericCode = record[3]

		countries.Countries = append(countries.Countries, c)
		englishNameMap[c.EnglishName] = c
		alpha2Map[c.Alpha2Code] = c
		alpha3Map[c.Alpha3Code] = c
		numericMap[c.NumericCode] = c

	}
	for key, value := range CountryMaps {
		fmt.Println("Map: ", key)
		var keys []string
		for k := range value {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range(keys) {
			fmt.Println("Key: ", k, "Value: ", value[k])
		}
	}
	fmt.Println("Json:")
	j, err := json.MarshalIndent(countries, "", "  ")
	if err == nil {
		fmt.Printf("%s\n", j)
	} else {
		fmt.Printf("gads: %s\n", err)
	}

}
