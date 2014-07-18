package country
// Keep reading: http://golang.org/doc/code.html#Testing
import (
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	fmt.Println("Test: country.Load")
	country := new(Country)
	n, err := country.Load()
	// check count or something.
	fmt.Printf("%d\n", n)
	fmt.Printf("%s\n", err)
}
