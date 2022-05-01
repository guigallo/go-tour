package main

import "fmt"

// func add(x, y int) int { // can omit on same types
func add(x int, y int) int {
	return x + y
}

// a function can return any number of results
func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	// return x, y // named return
	return // naked return x, y vars
}

func main() {
	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}
