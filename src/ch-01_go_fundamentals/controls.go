package main

import (
	"fmt"
	"math/rand"
)

func foo(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("unknown type", v)
	}
}

func main() {
	var x = rand.Int()

	if x == 1 {
		fmt.Println("X is equal to 1")
	} else {
		fmt.Println("X is not equal to 1")
	}

	switch x {
	case 1:
		fmt.Println("x is 1")
	case 2:
		fmt.Println("x is 2")
	default:
		fmt.Println("Default case")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	nums := []int{1, 2, 3, 4, 5}
	for idx, val := range nums {
		fmt.Println(idx, val)
	}

}
