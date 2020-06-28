package main

import (
	"fmt"

	"gitlab.com/jasb91/ejercicios-ninja-curso-golang/nivel13/02-code-starting/quote"
	"gitlab.com/jasb91/ejercicios-ninja-curso-golang/nivel13/02-code-starting/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
