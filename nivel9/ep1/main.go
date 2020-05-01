// Ejercicio práctico #1
//     • Además de la gorutina principal, lanza dos gorutinas adicionales
//         ◦ Cada gorutina debe imprimir algo en pantalla
//     • Usa waitgroups para asegurarte que cada gorutina finalize antes de que el programa termine

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("OS:\t\t", runtime.GOOS)
	fmt.Println("ARCH:\t\t", runtime.GOARCH)
	fmt.Println("CPU:\t\t", runtime.NumCPU())
	fmt.Println("Go Routine:\t", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("Iteration N°:", i)
		}
		wg.Done()
	}()
	fmt.Println("Go Routine 1:", runtime.NumGoroutine())

	go func() {
		fmt.Println("Primera Goroutine")
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println("Even N°:", i)
			}
		}
		wg.Done()
	}()
	fmt.Println("Go Routine 2:", runtime.NumGoroutine())

	go func() {
		fmt.Println("Tercera Goroutine")
		for i := 1; i <= 10; i++ {
			if i%2 != 0 {
				fmt.Println("Odd N°:", i)
			}
		}
		wg.Done()
	}()
	fmt.Println("Go Routine 3:", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("Go Routine Final:", runtime.NumGoroutine())
}
