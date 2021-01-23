package main

import (
	"fmt"
)

var num = 5

func main() {
	fmt.Println("Hello World")
	num1 := 2
	num2 := 4
	fmt.Println(add(num1, num2))
	square1, square2 := square(num1, num2)
	fmt.Println(square1, square2)
	for i := 1; i <= 5; i++ {
		fmt.Println("Loop")
	}
}

func add(x int, y int) int {
	var result = x + y + num
	return result
}

func square(x int, y int) (int, int) {
	var result1 = x * x
	var result2 = y * y
	return result1, result2
}
