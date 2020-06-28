package word

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	count := Count("This is an example")

	if count != 4 {
		t.Error("Expected", 4, "got", count)
	}

}

func TestUseCount(t *testing.T) {
	m := UseCount("This is an example, This is an ejemplo")

	for k, v := range m {
		switch k {
		case "This":
			if v != 2 {
				t.Error("Expected", 2, "got", v)
			}
		case "is":
			if v != 2 {
				t.Error("Expected", 2, "got", v)
			}
		case "an":
			if v != 2 {
				t.Error("Expected", 2, "got", v)
			}
		case "example":
			if v != 1 {
				t.Error("Expected", 1, "got", v)
			}
		case "ejemplo":
			if v != 1 {
				t.Error("Expected", 1, "got", v)
			}
		}
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("This is an example")
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount("This is an example")
	}
}

func ExampleCount() {
	s := "This is an example."
	fmt.Println(Count(s))
	//Output:
	//4
}
