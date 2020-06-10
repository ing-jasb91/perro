/*
Crea un struct “errorPer” el cual implemente la interfaz builtin.error.
Crea una función “foo” que tenga un valor de tipo error como parámetro.
Crea un valor de tipo “errorPer” y pásalo a “foo”. Si necesitas un pista:
https://play.golang.org/p/ESmKdz-T3At
solución:
    • https://play.golang.org/p/AUpI7MCqPGl
    • assertion
        ◦ https://play.golang.org/p/l3IKqOu8mJo
    • conversion
        ◦ https://play.golang.org/p/i046JdkXH8x
video: 178
*/
package main

import "fmt"

type errorPer struct {
	info string
}

func (e errorPer) Error() string {
	return fmt.Sprintf("El error: %v", e.info)
}

func foo(e error) {
	fmt.Println("foo corrió", e)
}

func main() {
	e1 := errorPer{
		info: "Necesito dormir :/",
	}
	foo(e1)

}
