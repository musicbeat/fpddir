package stddata
/*
 * data loader component for ISO 639.2 language codes.
 * source is US Library of Congress
 * http://www.loc.gov/standards/iso639-2/ISO-639-2_utf-8.txt
 * These are notes from:
 * http://www.loc.gov/standards/iso639-2/ascii_8bits.html
 * "These files may be used to download the list of
 * language codes with their language names, for example into a
 * database. To read the files, please note that one line of text
 * contains one entry. An alpha-3 (bibliographic) code, an alpha-3
 * (terminologic) code (when given), an alpha-2 code (when given),
 * an English name, and a French name of a language are all separated
 * by pipe (|) characters. If one of these elements is not applicable
 * to the entry, the field is left empty, i.e., a pipe (|) character
 * immediately follows the preceding entry. The Line terminator is
 * the LF character."
 */

import (
	"encoding/json"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
)

type Language struct {
	Alpha3bibliographic	string
	Alpha3terminologic	string
	Alpha2				string
	EnglishName			string
	FrenchName			string
}
type Languages struct {
	Languages []Language
}

var LanguageMaps map[string] map[string]Language
var AlphaMap map[string]Language
var EnglishNameMap map[string]Language

func LoadIso639() {
	// initialize the maps:
	AlphaMap = make(map[string]Language)
	EnglishNameMap = make(map[string]Language)
	LanguageMaps = make(map[string] map[string]Language)
	LanguageMaps["Alpha"] = AlphaMap
	LanguageMaps["English"] = EnglishNameMap
	
	res, err := http.Get("http://www.loc.gov/standards/iso639-2/ISO-639-2_utf-8.txt")
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(res.Body)
	reader.Comma = '|'
	reader.FieldsPerRecord = 5
	reader.TrimLeadingSpace = true

	defer res.Body.Close()

	var languages Languages
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

		var l Language
		l.Alpha3bibliographic = record[0]
		l.Alpha3terminologic = record[1]
		l.Alpha2 = record[2]
		l.EnglishName = record[3]
		l.FrenchName = record[4]

		languages.Languages = append(languages.Languages, l)

		AlphaMap[l.Alpha3bibliographic] = l
		EnglishNameMap[l.EnglishName] = l

	}
	for key, value := range LanguageMaps {
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
	j, err := json.MarshalIndent(languages, "", "  ")
	if err == nil {
		fmt.Printf("%s\n", j)
	} else {
		fmt.Printf("gads: %s\n", err)
	}

}
