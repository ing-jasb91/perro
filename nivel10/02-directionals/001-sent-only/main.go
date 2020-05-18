/*
Ejercicio Práctico #2(a)
    • Haz que este código funcione:
        ◦ https://play.golang.org/p/oB-p3KMiH6
			▪ Solución:  https://play.golang.org/p/isnJ8hMMKg
Respuesta: https://play.golang.org/p/QHxrG8UEiuq
*/

package main

import (
	"fmt"
)

func main() {
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
