package main

import (
	"fmt"
	"net"
)

func main() {
	for port := 1; port <= 1024; port++ {
		addressPort := fmt.Sprintf("scanme.nmap.org:%d", port)
		conn, err := net.Dial("tcp", addressPort)

		if err != nil {
			// порт закрыт или за файерволлом
			continue
		}

		conn.Close()
		fmt.Println("Open", port)
	}
}
