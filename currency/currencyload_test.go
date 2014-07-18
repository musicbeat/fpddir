package currency
// Keep reading: http://golang.org/doc/code.html#Testing
import (
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	fmt.Println("Test: currency.Load")
	currency := new(Currency)
	n, err := currency.Load()
	// check count or something.
	fmt.Printf("%d\n", n)
	fmt.Printf("%s\n", err)
}
