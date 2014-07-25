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
	if n != expected {
		t.Fatalf("Expected to load %d, loaded %d\n", expected, n)
	}
	// check count or something.
	fmt.Printf("%d\n", n)
	fmt.Printf("%s\n", err)
}
