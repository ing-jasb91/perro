/*
Este ejercicio te ayudará a reforzar el concepto de method sets:
    • Crea un tipo struct persona
    • Adjunta el método hablar usando un receptor de tipo puntero
        ◦ *persona
    • Crea un tipo interfaz humano
        ◦ Para implícitamente implementar esa interfaz, un tipo humano debe tener el método hablar
    • Crea la función “diAlgo”
        ◦ Haz que tome un humano como parámetro
        ◦ Haz que llame al método hablar
    • Muestra lo siguiente en tu código
        ◦ PUEDES pasar un valor de tipo *persona a diAlgo
        ◦ NO puedes pasar un valor de tipo persona a diAlgo
	• Aquí hay una pista si necesitas ayuda
		◦ https://play.golang.org/p/JQd6LsU_L-K
	Receptores       Valores
	-----------------------------------------------
	(t T)           T y *T
	(t *T)          *T

	Solución: https://gitlab.com/eduar/go-programming

*/
package main

import "fmt"

type persona struct {
	Nombre   string
	Apellido string
	Edad     int
}

func (h *persona) hablar() {
	fmt.Printf("Esta persona %s %s habla\n", h.Nombre, h.Apellido)
}

type humano interface {
	hablar()
}

func diAlgo(h humano) {
	h.hablar()
}

func main() {
	persona1 := persona{
		Nombre:   "Jhon",
		Apellido: "Serrano",
		Edad:     28,
	}
	// La interfaz humano solo permite valores de tipo *persona
	// diAlgo(persona1)
	// Este valor es permitido porque es una dirección de memoria, apunta a un valor.
	diAlgo(&persona1)
	// Aunque el método hablar recibe parámetros de tipo *persona, puede recibir valores normales
	// Mientras sean directamente desde una variable y no una interface.
	persona1.hablar()
}
