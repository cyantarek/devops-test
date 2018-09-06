package main

import "fmt"

func main() {
	fmt.Println("the result is:", Add(5, 6))
}

//Add Func
func Add(x, y int) int {
	return x + y
}
