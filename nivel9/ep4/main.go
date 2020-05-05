/*
Ejercicio práctico #4
Arregla la race condition que creaste en el ejercicio anterior usando un mutex
●	Tiene sentido eliminar runtime.Gosched()
código: https://gitlab.com/eduar/go-programming
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mux sync.Mutex

func main() {

	fmt.Println(
		"OS:\t", runtime.GOOS,
		"\nARCH:\t", runtime.GOARCH,
		"\nCPU's\t", runtime.NumCPU(),
		"\nGoroutines:\t", runtime.NumGoroutine(),
	)

	incremento := 0
	const gs int = 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mux.Lock()

			v := incremento
			//runtime.Gosched()
			v++
			incremento = v
			mux.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:\t", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Incremento:\t", incremento)

}
