package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Tyler Durden")
	if s != "Hello there Tyler Durden" {
		t.Error("Expected", "Hello there Tyler Durden", "Got", s)
	}
}

// godoc -http=:808
func ExampleGreet() {
	fmt.Println(Greet("Tyler Durden"))
	// Output:
	// Hello there Tyler Durden
}

// go test -bench .
func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Tyler Durden")
	}
}
