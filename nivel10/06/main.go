/*
Ejercicio Práctico #6
    • Escribe un programa que:
        ◦ Ponga 100 números en un canal
		◦ Extraiga los números del canal y los imprima
*/
package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {

		for i := 0; i <= 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

}
