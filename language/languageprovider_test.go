package language
// Keep reading: http://golang.org/doc/code.html#Testing
import (
	"fmt"
	"testing"

	. "github.com/musicbeat/stddata"
)
var p Provider
var expected int = 486

func TestLanguageProviderLoad(t *testing.T) {
	fmt.Println("Test: LanguageProvider.Load")
	p = new(LanguageProvider)
	n, err := p.Load()
	if err != nil {
		t.Fatal()
	}
	if n != expected {
		t.Fatalf("Expected to load %d, loaded %d\n", expected, n)
	}
}
func TestAlphaSearch(t *testing.T) {
	matches, err := p.Search("alpha", "fr")
	if err != nil {
		t.Fatalf("Err %v\n", err)
	}
	fmt.Println("matches %s\n", matches)
}
func TestNameSearch(t *testing.T) {
	matches, err := p.Search("name", "an")
	if err != nil {
		t.Fatalf("Err %v\n", err)
	}
	fmt.Println("matches %s\n", matches)
}
func TestNameSearchLowerCase(t *testing.T) {
	matches, err := p.Search("name", "en")
	if err != nil {
		t.Fatalf("Err %v\n", err)
	}
	fmt.Println("matches %s\n", matches)
}
func BenchmarkNameSearch(b *testing.B) {
	p = new(LanguageProvider)
	n, err := p.Load()
	if err != nil {
		b.Fatal()
	}
	if n != expected {
		b.Fatalf("Expected to load %d, loaded %d\n", expected, n)
	}
	b.ResetTimer()
	fmt.Printf("start runs...")
	for i := 0; i < b.N; i++ {
		_, err := p.Search("name", "a")
		if err != nil {
			b.Fatalf("Err %v\n", err)
		}
	}
}
