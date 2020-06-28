package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	yDog := Years(3)
	if yDog != 21 {
		t.Error("Expected", 21, "Got", Years(3))
	}
}

func TestYearsTwo(t *testing.T) {
	yDog := YearsTwo(3)
	if yDog != 21 {
		t.Error("Expected", 21, "Got", YearsTwo(3))
	}
}

func ExampleYears() {
	fmt.Println(Years(4))
	//Output:
	//28
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(4))
	//Output:
	//28
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
