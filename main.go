package main

import (
	"fmt"

	"golang.concurrency/fibonacci"
)

func main() {
	for num := range fibonacci.GenerateFibonnaciSequence(20) {
		fmt.Println(num)
	}
	fmt.Println("Done!")
}
