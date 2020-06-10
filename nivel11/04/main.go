/*
Ejercicio Práctico #4
Comenzando con este código https://play.golang.org/p/jjfF_jsGplU,
usa el struct raiz.Error como valor de tipo error.
Si quieres usa estos valores para tu coordenada:
    • lat "50.2289 N"
    • long "99.4656 W"
solución:
	• https://play.golang.org/p/Ezz-YX3tCcv
*/
package main

import (
	"fmt"
	"log"
	"math"
)

type raizError struct {
	lat  string
	long string
	err  error
}

func (re raizError) Error() string {
	return fmt.Sprintf("error matemático en %v %v %v", re.lat, re.long, re.err)
}

func main() {

	_, err := raiz(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func raiz(f float64) (float64, error) {

	if f < 0 {
		// Escribe tu código aquí
		e := fmt.Errorf("El valor es negativo: %v", f)

		return 0.0, raizError{
			lat:  "10.10 N",
			long: "66.85 W",
			err:  e,
		}

	}
	return math.Sqrt(f), nil
}
