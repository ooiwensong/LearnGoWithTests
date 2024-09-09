package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2,2)
	expected := 4

	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}

// Testable examples are compiled whenever tests are executed. Examples
// are validated by the Go compiler, allowing developers to be confident
// that doc's examples always reflect current code behaviour
func ExampleAdd() {
	sum := Add(1,5)
	fmt.Println(sum)
	// Output: 6
}