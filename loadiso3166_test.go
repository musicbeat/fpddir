package stddata
// Keep reading: http://golang.org/doc/code.html#Testing
import (
	"fmt"
	"testing"
)

func TestLoadIso3166(t *testing.T) {
	fmt.Println("Test: LoadIso3166")
	LoadIso3166()
	// check count or something.
}
