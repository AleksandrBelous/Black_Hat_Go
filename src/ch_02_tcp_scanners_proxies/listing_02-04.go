package main

import (
	"fmt"
	"net"
)

func oneScan(port int) {
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
	for port := 1; port <= 1024; port++ {
		go oneScan(port)
	}
}
