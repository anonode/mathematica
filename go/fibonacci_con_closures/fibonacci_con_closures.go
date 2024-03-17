package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	return func(x int) int {
		var ret int
		if x == 0 {
			ret = 0
		} else if x < 2 {
			ret = 1
		} else {
			fib := make([]int, x)
			fib[0] = 1
			fib[1] = 1
			for i := 2; i < x; i++ {
				fib[i] = fib[i-1] + fib[i-2]
			}
			ret = fib[x-1]
		}
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
