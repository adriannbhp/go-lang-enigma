package main

import (
	"fmt"
)

func main() {
	//
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Scan()
	//input := scanner.Text()
	//
	//// Convert input to integer
	//count, err := strconv.Atoi(input)
	//if err != nil {
	//	fmt.Println("Input harus berupa angka.")
	//	return
	//}
	//
	//for i := count; i > 0; i-- {
	//	fmt.Printf("%v  I will become a go developer \n", i)
	//}

	var count int
	fmt.Scanln(&count)
	for i := count; i >= 1; i-- {
		fmt.Println(i, " I will become a go developer")
	}
}
