package stddata
// Keep reading: http://golang.org/doc/code.html#Testing
import (
	"fmt"
	"testing"
)

func TestLoadIso639(t *testing.T) {
	fmt.Println("Test: LoadIso639")
	LoadIso639()
	// check count or something.
}
