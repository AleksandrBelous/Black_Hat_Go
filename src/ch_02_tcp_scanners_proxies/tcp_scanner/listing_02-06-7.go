package main

import (
	"fmt"
	"sync"
)

func workerWG(ports chan int, wg *sync.WaitGroup) {
	for port := range ports {
		fmt.Println(port)
		wg.Done()
	}
}

func main() {
	var wg sync.WaitGroup

	ports := make(chan int, 100)

	for i := 0; i < cap(ports); i++ {
		go workerWG(ports, &wg)
	}

	for port := 1; port <= 1024; port++ {
		wg.Add(1)
		ports <- port
	}

	wg.Wait()

	close(ports)
}
