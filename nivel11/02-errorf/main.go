/*
Ejercicio Práctico #2
Comienza con este código https://play.golang.org/p/kbeQfbe82Nl.
Crea un mensaje de error personalizado usando “fmt.Errorf”.
solución:
    • https://play.golang.org/p/p3xCMEpg5Hc
    • https://play.golang.org/p/repWMzfOOiu
	• https://play.golang.org/p/9RFq8BXHWm5

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

	bs, err := aJSON(p1)
	if err != nil {
		errJSON := fmt.Errorf("Algun error en el formato JSON: %v", err)
		log.Fatalln(errJSON)

	}
	fmt.Println(string(bs))

}

// aJSON también necesita retorna un error
func aJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, err
	}
	return bs, nil
}
