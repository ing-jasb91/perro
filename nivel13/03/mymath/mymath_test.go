package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type myInts struct {
		data   []int
		answer float64
	}

	tests := []myInts{
		{[]int{1, 4, 6, 8, 100}, 6},
		{[]int{0, 8, 10, 1000}, 9},
		{[]int{9000, 4, 10, 8, 6, 12}, 9},
		{[]int{123, 744, 140, 200}, 170},
	}

	for _, v := range tests {
		if CenteredAvg(v.data) != v.answer {
			t.Error("Expected", v.answer, "Got", CenteredAvg(v.data))
		}
	}

}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{9000, 4, 10, 8, 6, 12})
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{9000, 4, 10, 8, 6, 12}))
	// Output:
	// 9
}
