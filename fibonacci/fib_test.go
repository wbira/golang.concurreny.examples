package fibonacci_test

import (
	"fmt"
	"testing"

	"golang.concurrency/fibonacci"
)

func TestGenerationFibonacci(t *testing.T) {
	for num := range fibonacci.GenerateFibonnaciSequence(10) {
		fmt.Println(num)
	}
}
