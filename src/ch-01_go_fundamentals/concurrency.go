package main

import (
	"fmt"
	"time"
)

func f() {
	fmt.Println("f function")
}

func strlen(s string, c chan int) {
	c <- len(s)
}

func main() {
	go f()
	time.Sleep(1 * time.Second)
	fmt.Println("main function")

	c := make(chan int)
	go strlen("Salutations", c)
	go strlen("World", c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
