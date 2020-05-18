/*
Ejercicio Práctico #1
    • Haz que este código funcione:
        ◦ Usando una función literal, también conocida como, función anónima autoejecutable
            ▪ Solución: https://play.golang.org/p/SHr3lpX4so
        ◦ Un canal con búfer
			▪ Solución: https://play.golang.org/p/Y0Hx6IZc3U

			Respuesta: https://play.golang.org/p/cmBIx9_9gfH
*/

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
