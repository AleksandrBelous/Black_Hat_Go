package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Black Hat Gophers!")

	var count = int(42)
	fmt.Println(&count)
	ptr := &count
	fmt.Println(*ptr)
	*ptr = 100
	fmt.Println(count)
}
