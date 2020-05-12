/*
Ejercicio práctico #6
Crea un programa que imprima tu OS y ARCH. Usa los siguientes comandos para correrlo
    • go run
    • go build
    • go install
código: https://gitlab.com/eduar/go-programming
*/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print(
		"\n************************************",
		"\nOperating System:\t", runtime.GOOS,
		"\nCPU Arch:\t\t", runtime.GOARCH,
		"\nGo Version:\t\t", runtime.Version(),
		"\nGo Root:\t\t", runtime.GOROOT(),
		"\n************************************\n\n",
	)

}
