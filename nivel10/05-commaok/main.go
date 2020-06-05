/*
Ejercicio Práctico #5
    • Demuestra el uso del idioma coma ok con este código https://play.golang.org/p/YHOMV9NYKK.
solución: https://play.golang.org/p/qh2ywLB5OG
video: 168
*/

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 1000
		close(c)
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	v, ok = <-c
	fmt.Println(v, ok)

}
