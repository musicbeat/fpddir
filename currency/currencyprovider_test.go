package currency
// Keep reading: http://golang.org/doc/code.html#Testing
import (
	"fmt"
	"testing"

	. "github.com/musicbeat/stddata"
)
var p Provider
var expected int = 180

func TestCurrencyProviderLoad(t *testing.T) {
	fmt.Println("Test: CurrencyProvider.Load")
	p = new(CurrencyProvider)
	n, err := p.Load()
	if err != nil {
		t.Fatal()
	}
	if n != expected {
		t.Fatalf("Expected to load %d, loaded %d\n", expected, n)
	}
}
func TestCountrySearch(t *testing.T) {
	names, err := p.Search("country", "A")
	if err != nil {
		t.Fatalf("Err %v\n", err)
	}
	fmt.Println("names %s\n", names)
}
func TestNameSearch(t *testing.T) {
	names, err := p.Search("name", "A")
	if err != nil {
		t.Fatalf("Err %v\n", err)
	}
	fmt.Println("names %s\n", names)
}
func TestNameSearchLowerCase(t *testing.T) {
	names, err := p.Search("name", "a")
	if err != nil {
		t.Fatalf("Err %v\n", err)
	}
	fmt.Println("names %s\n", names)
}
func TestCodeSearch(t *testing.T) {
	names, err := p.Search("code", "E")
	if err != nil {
		t.Fatalf("Err %v\n", err)
	}
	fmt.Println("names %s\n", names)
}
func TestNumberSearch(t *testing.T) {
	names, err := p.Search("number", "0")
	if err != nil {
		t.Fatalf("Err %v\n", err)
	}
	fmt.Println("names %s\n", names)
}
func BenchmarkNameSearch(b *testing.B) {
	p = new(CurrencyProvider)
	n, err := p.Load()
	if err != nil {
		b.Fatal()
	}
	if n != expected {
		b.Fatalf("Expected to load %d, loaded %d\n", expected, n)
	}
	b.ResetTimer()
	fmt.Printf("start runs...")
	for i := 0; i < b.N; i++ {
		_, err := p.Search("name", "a")
		if err != nil {
			b.Fatalf("Err %v\n", err)
		}
	}
}
