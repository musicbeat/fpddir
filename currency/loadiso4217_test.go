package stddata
// Keep reading: http://golang.org/doc/code.html#Testing
import (
	"fmt"
	"testing"
)

func TestLoadIso4217(t *testing.T) {
	fmt.Println("Test: LoadIso4217")
	LoadIso4217()
	// check count or something.
}
