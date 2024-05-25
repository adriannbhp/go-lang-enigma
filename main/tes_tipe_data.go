package main

import "fmt"

func main() {
	var x int
	var y float64

	x = 5
	y = 12.32

	result1 := x + int(y)
	fmt.Println(result1)

	result2 := float64(x) - y
	fmt.Printf("%.2f\n", result2)

	result3 := x * int(y)
	fmt.Println(result3)
}
