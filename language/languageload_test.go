package language
// Keep reading: http://golang.org/doc/code.html#Testing
import (
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	fmt.Println("Test: language.Load")
	l := new(Language)
	n, err := l.Load()
	// check count or something.
	fmt.Printf("%d\n", n)
	fmt.Printf("%s\n", err)
}
