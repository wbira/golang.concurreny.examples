package fibonacci

func GenerateFibonnaciSequence(n int) <-chan int {
	numbers := make(chan int)

	go func() {
		for i, f1, f2 := 0, 0, 1; i < n; i, f1, f2 = i+1, f1+f2, f1 {
			numbers <- f1
		}
		//close(numbers)
	}()

	return numbers
}
