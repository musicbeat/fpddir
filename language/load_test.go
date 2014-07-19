package language
// Keep reading: http://golang.org/doc/code.html#Testing
import (
	"fmt"
	"testing"
)

var expected = 486

func TestLoad(t *testing.T) {
	fmt.Println("Test: language.Load")
	l := new(Language)
	if n, err := l.Load(); n!= expected {
		t.Errorf("language.Load n=%d, want %d\n", n, expected)
		if err != nil {
			t.Errorf("language.Load err %s\n", err)
		}
	}
}
