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
	Name string
}

func (p *persona) toTalk() string {
	return fmt.Sprintln("Hello,", p.Name)
}

type human interface {
	toTalk() string
}

func saySomething(h human) {
	fmt.Println(h.toTalk())
}

func main() {
	p1 := persona{
		Name: "Wilda",
	}
	saySomething(&p1)
	//saySomething(p1)
	fmt.Println(p1.toTalk())
}
