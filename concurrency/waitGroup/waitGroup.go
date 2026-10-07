package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printName(name string) {
	fmt.Println(name)
	wg.Done()
}

func main() {
	wg.Add(1)
	go printName("Wilson")
	wg.Add(1)
	go printName("Patyal")

	wg.Wait()
	fmt.Println("End")
}
