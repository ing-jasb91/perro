package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incremento := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := incremento
			runtime.Gosched()
			v++
			incremento = v
			fmt.Println(incremento)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("El valor final de incremento es:", incremento)
}
