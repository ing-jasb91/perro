package main

import (
	"fmt"

	"gitlab.com/jasb91/ejercicios-ninja-curso-golang/nivel13/01-bet/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
