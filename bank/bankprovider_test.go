package bank

// Keep reading: http://golang.org/doc/code.html#Testing
import (
	"fmt"
	"testing"

	. "github.com/musicbeat/stddata"
)

var p Provider

// number of banks (at least for now)
var expected = 19817

func TestBankProviderLoad(t *testing.T) {
	fmt.Println("Test: bank.Load")
	p = new(BankProvider)
	n, err := p.Load()
	if err != nil {
		t.Fatal()
	}
	if n != expected {
		t.Fatalf("Expected to load %d, loaded %d\n", expected, n)
	}
}
func TestBankNameSearch(t *testing.T) {
	// name search:
	matches, err := p.Search("name", "AB")
	if err != nil {
		t.Fatalf("Err %v\n", err)
	}
	fmt.Println("matches %s\n", matches)
}
func TestBankNameSearchLowerCase(t *testing.T) {
	// name search:
	matches, err := p.Search("name", "ab")
	if err != nil {
		t.Fatalf("Err %v\n", err)
	}
	fmt.Println("matches %s\n", matches)
}
func TestBankNumberSearch(t *testing.T) {
	// number search:
	numbers, err := p.Search("number", "123")
	if err != nil {
		t.Fatalf("Err %v\n", err)
	}
	fmt.Println("numbers %s\n", numbers)
}
func BenchmarkNameSearch(b *testing.B) {
	p = new(BankProvider)
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
		_, err := p.Search("name", "ab")
		if err != nil {
			b.Fatalf("Err %v\n", err)
		}
	}
}
