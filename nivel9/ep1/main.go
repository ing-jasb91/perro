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
	fmt.Println(
		"OS:\t", runtime.GOOS,
		"\nARCH:\t", runtime.GOARCH,
		"\nVersion\t", runtime.Version(),
		"\nCPU's\t", runtime.NumCPU(),
		"\nGoroutines:\t", runtime.NumGoroutine(),
	)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Algo1", i)
		}
		fmt.Println("\nGoroutines:\t", runtime.NumGoroutine())
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Algo2", i)
		}
		fmt.Println("\nGoroutines:\t", runtime.NumGoroutine())
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("\nGoroutines:\t", runtime.NumGoroutine())
	fmt.Println()
}
