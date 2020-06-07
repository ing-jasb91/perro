/*
Ejercicio Práctico #7
    • Escribe un programa que:
        ◦ Lance 10 gorutinas
            ▪ Cada gorutina agrega 10 números a un canal
        ◦ Extrae los números del canal e imprímelos
soluciones:
    • https://play.golang.org/p/bkAy9ECzVpd
	• https://play.golang.org/p/83Vh5ymfhS2
*/
package main

import "fmt"

func main() {
	ch := make(chan int)
	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				ch <- i
			}
		}()
	}
	for k := 0; k < 100; k++ {
		fmt.Println(k, <-ch)

	}

}
