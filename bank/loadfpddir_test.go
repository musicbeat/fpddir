package stddata
// Keep reading: http://golang.org/doc/code.html#Testing
import (
	"fmt"
	"testing"
)

func TestLoadFpddir(t *testing.T) {
	fmt.Println("Test: LoadFpddir")
	LoadFpddir()
	// check count or something.
}
