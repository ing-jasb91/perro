/*
Ejercicio Práctico #1
Comienza con este código https://play.golang.org/p/Bgads38iFgz.
En vez de usar el identificador blank (underscore), asegúrate
de que el código esté chequeando y manejando el error.

solución:
    • https://play.golang.org/p/tn8oiuL1Yn

*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type persona struct {
	Nombre   string
	Apellido string
	Frases   []string
}

func main() {
	p1 := persona{
		Nombre:   "James",
		Apellido: "Bond",
		Frases:   []string{"Shaken, not stirred", "¿Algún último deseo?", "Nunca digas nunca."},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(bs))
}
