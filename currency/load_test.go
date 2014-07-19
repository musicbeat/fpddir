package currency
// Keep reading: http://golang.org/doc/code.html#Testing
import (
	"fmt"
	"testing"
)
var expected int = 180

func TestLoad(t *testing.T) {
	fmt.Println("Test: currency.Load")
	currency := new(Currency)
	if n, err := currency.Load(); n != expected {
		t.Errorf("currency.Load n=%d, want %d\n", n, expected)
		if err != nil {
			t.Errorf("currency.Load err %s\n", err)
		}
	}
}
