package bank
// Keep reading: http://golang.org/doc/code.html#Testing
import (
	"fmt"
	"testing"

	. "github.com/musicbeat/stddata"
)

func TestBankprovider(t *testing.T) {
	expected := 8648
	fmt.Println("Test: bank.Load")
	var p Provider
	p = new(BankProvider)
	n, err := p.Load()
	if err != nil {
		t.Fatal()
	}
	if n != expected {
		t.Fatalf("Expected to load %d, loaded %d\n", expected, n)
	}
	// name search:
	names, err := p.Search("name", "AB")
	if err != nil {
		t.Fatalf("Err %v\n", err)
	}
	fmt.Println("names %s\n", names)
	// name search:
	names, err = p.Search("name", "ab")
	if err != nil {
		t.Fatalf("Err %v\n", err)
	}
	fmt.Println("names %s\n", names)
	// number search:
	numbers, err := p.Search("number", "123")
	if err != nil {
		t.Fatalf("Err %v\n", err)
	}
	fmt.Println("numbers %s\n", numbers)
}
