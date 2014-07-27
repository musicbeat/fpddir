package country
// Keep reading: http://golang.org/doc/code.html#Testing
import (
	"fmt"
	"testing"

	. "github.com/musicbeat/stddata"
)
var p Provider

func TestCountryProvider(t *testing.T) {
	expected := 249
	fmt.Println("Test: CountryProvider.Load")
	p = new(CountryProvider)
	n, err := p.Load()
	if err != nil {
		t.Fatalf("Err %v\n", err)
	}
	if n != expected {
		t.Fatalf("Expected to load %d, loaded %d\n", expected, n)
	}
}
func TestNameSearch(t *testing.T) {
	_, err := p.Search("name", "A")
	if err != nil {
		t.Fatalf("Err %v\n", err)
	}
}
func TestNameSearchLowerCase(t *testing.T) {
	_, err := p.Search("name", "b")
	if err != nil {
		t.Fatalf("Err %v\n", err)
	}
}
func TestAlpha2Search(t *testing.T) {
	_, err := p.Search("alpha2", "C")
	if err != nil {
		t.Fatalf("Err %v\n", err)
	}
}
func TestAlpha3Search(t *testing.T) {
	_, err := p.Search("alpha3", "U")
	if err != nil {
		t.Fatalf("Err %v\n", err)
	}
}
func TestNumberSearch(t *testing.T) {
	_, err := p.Search("number", "1")
	if err != nil {
		t.Fatalf("Err %v\n", err)
	}
}

