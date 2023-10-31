package main

import (
	"fmt"
	"main/singleton"
)

func main() {
	instance1 := counter.GetInstance()
	instance2 := counter.GetInstance()

	fmt.Println(instance1.Increment()) // Output: 1
	fmt.Println(instance2.Increment()) // Output: 2
	fmt.Println(instance1.Increment()) // Output: 3
}
