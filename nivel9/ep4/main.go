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

func main() {
	var wg sync.WaitGroup
	var incremento int
	gs := 60
	fmt.Println("N° CPU:", runtime.NumCPU())
	fmt.Println("N° Goroutine:", runtime.NumGoroutine())
	
	
	wg.Add(gs)
	fmt.Println("N° Goroutine 1:", runtime.NumGoroutine())
	var mx sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mx.Lock()
			v := incremento
			//runtime.Gosched()
			v++
			incremento = v
			fmt.Println(incremento)
			mx.Unlock()
			wg.Done()
		}()
	}
	fmt.Println("N° Goroutine 2:", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("N° Goroutine:", runtime.NumGoroutine())
	fmt.Println("El valor final del incremento es:", incremento)
}