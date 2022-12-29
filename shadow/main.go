package main

import "fmt"

func main() {
	condition := true

	x := 1
	if condition {
		x := 2
		// x = 2
		fmt.Println("x = ", x)
	}

	// x = 1
	fmt.Println("x = ", x)
}
