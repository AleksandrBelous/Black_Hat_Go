package main

import (
	"fmt"
	"net"
	"sync"
)

func oneScanWG(port int, wg *sync.WaitGroup) {
	defer wg.Done()
	addressPort := fmt.Sprintf("scanme.nmap.org:%d", port)
	conn, err := net.Dial("tcp", addressPort)

	if err != nil {
		// порт закрыт или за файерволлом
		return
	}

	conn.Close()
	fmt.Println("Open", port)
}

func main() {
	var wg sync.WaitGroup
	for port := 1; port <= 1024; port++ {
		wg.Add(1)
		go oneScanWG(port, &wg)
	}
	wg.Wait()
}
