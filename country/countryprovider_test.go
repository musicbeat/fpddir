package country
// Keep reading: http://golang.org/doc/code.html#Testing
import (
	"fmt"
	"testing"

	. "github.com/musicbeat/stddata"
)

func TestCountryProvider(t *testing.T) {
	expected := 249
	fmt.Println("Test: CountryProvider.Load")
	var p Provider
	p = new(CountryProvider)
	n, err := p.Load()
	if err != nil {
		t.Fatal()
	}
	if n != expected {
		t.Fatalf("Expected to load %d, loaded %d\n", expected, n)
	}
	// name search:
	names, err := p.Search("name", "A")
	if err != nil {
		t.Fatalf("Err %v\n", err)
	}
	fmt.Println("names %s\n", names)
	// name search:
	names, err = p.Search("name", "b")
	if err != nil {
		t.Fatalf("Err %v\n", err)
	}
	fmt.Println("names %s\n", names)
	// alpha2 search:
	alpha2, err := p.Search("alpha2", "C")
	if err != nil {
		t.Fatalf("Err %v\n", err)
	}
	fmt.Println("alpha2 %s\n", alpha2)
	// alpha3 search:
	alpha3, err := p.Search("alpha3", "U")
	if err != nil {
		t.Fatalf("Err %v\n", err)
	}
	fmt.Println("alpha3 %s\n", alpha3)
	// number search:
	numbers, err := p.Search("number", "1")
	if err != nil {
		t.Fatalf("Err %v\n", err)
	}
	fmt.Println("numbers %s\n", numbers)
}

