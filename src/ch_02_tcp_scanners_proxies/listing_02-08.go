package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, results chan int) {
	for port := range ports {
		addressPort := fmt.Sprintf("scanme.nmap.org:%d", port)
		conn, err := net.Dial("tcp", addressPort)

		if err != nil {
			// порт закрыт или за файерволлом
			results <- 0
			continue
		}

		conn.Close()
		results <- port
	}
}

func main() {
	var openPorts []int

	ports := make(chan int, 1000)
	results := make(chan int)

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for port := 1; port <= 1024; port++ {
			ports <- port
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results

		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	close(ports)
	close(results)

	sort.Ints(openPorts)

	for _, port := range openPorts {
		fmt.Println("open", port)
	}
}
