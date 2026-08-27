package main

import (
	"fmt"
	"sync"
)

func worker(ports chan int, wg *sync.WaitGroup) {
	for port := range ports {
		fmt.Println(port)
		wg.Done()
	}
}

func main() {
	ports := make(chan int, 100)
	var wg sync.WaitGroup

	for i := 0; i < cap(ports); i++ {
		go worker(ports, &wg)
	}

	for port := 1; port <= 1024; port++ {
		wg.Add(1)
		ports <- port
	}

	wg.Wait()
	close(ports)
}
