package main

import "fmt"

func fibonacci() func(int) int {
	var n_1 = 0
	var n_2 = 1

	return func(x int) int {
		if x == 0 || x == 1 {
			return x
		}
		var result = n_1 + n_2
		n_1 = n_2
		n_2 = result

		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
