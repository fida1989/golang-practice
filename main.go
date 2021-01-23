package main

import (
	"fmt"
	"math"
)

var num = 5

func main() {
	fmt.Println("Hello World")
	//
	var sum = add(5, 6)
	fmt.Printf("Ourput is %d\n",sum)
	//
	con :=6
	if(con > 5){
		fmt.Println("yes")
	}else{
		fmt.Println("no")
	}
	//
	num1 := 2.0
	num2 := 4.0
	square1, square2 := square(num1, num2)
	fmt.Println(square1, square2)
	//
	for i := 1; i <= 5; i++ {
		fmt.Println("Loop")
	}
}

func add(x int, y int) int {
	var result = x + y + num
	return result
}

func square(x float64, y float64) (float64, float64) {
	var result1 = math.Sqrt(x)
	var result2 = y * y
	return result1, result2
}
