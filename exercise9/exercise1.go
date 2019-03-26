package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("begin CPU", runtime.NumCPU())
	fmt.Println("begin GS", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("hello from thing one")
		wg.Done()
	}()
	go func() {
		fmt.Println("hello from anything")
		wg.Done()
	}()

	fmt.Println("mid CPU", runtime.NumCPU())
	fmt.Println("mid GS", runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("Going to exit!")
	fmt.Println("end CPU", runtime.NumCPU())
	fmt.Println("end GS", runtime.NumGoroutine())
}
