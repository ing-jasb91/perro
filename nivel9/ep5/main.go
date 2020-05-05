/*
Ejercicio práctico #5
Arreglar la race condition que creaste en el ejercicio #4 usando el paquete atomic
código: https://gitlab.com/eduar/go-programming
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {

	fmt.Println(
		"OS:\t", runtime.GOOS,
		"\nARCH:\t", runtime.GOARCH,
		"\nCPU's\t", runtime.NumCPU(),
		"\nGR:\t", runtime.NumGoroutine(),
	)

	var incremento int64
	const gs int = 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incremento, 1)
			runtime.Gosched()
			fmt.Println("Contador:", atomic.LoadInt64(&incremento))
			wg.Done()
		}()
		fmt.Println("Gorout:\t", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Incremento:", incremento)

}
