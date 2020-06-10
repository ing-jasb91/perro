/*
Ejercicio Práctico #1
Crea un paquete perro. El paquete perro debería tener una función exportada “Edad”
el cuál toma años humano y los retorna como años de perro (1 año humano = 7 años de perro).
Documenta tu código con comentarios. Usa este código en func main.

		• Corre el programa y asegúrate de que funciona.
		• Corre un servidor local con godoc y revisa la documentación.

solución: https://gitlab.com/eduar/go-programming
*/
package main

import (
	"fmt"

	"gitlab.com/jasb91/ejercicios-ninja-curso-golang/nivel12/01-perro/perro"
)

func main() {
	fmt.Println(perro.Edad(4))

}
