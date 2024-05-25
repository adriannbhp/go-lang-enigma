package main

import "fmt"

func main() {

	// Mendeklarasikan menggunakan var
	var name string = "Adrian"
	fmt.Println(name)

	// Mendeklarasikan menggunakan Type Inference
	alamat := "Bandung"
	fmt.Println(alamat)

	// Mendeklarasikan menggunakan multiple variable
	var a, b int = 1, 2
	fmt.Println(a, b)

	var isGoldMember bool = true
	fmt.Println(isGoldMember)

	// Default nya False
	var isGoldCard bool
	fmt.Println(isGoldCard)

	const fullName string = "Adrianbhp"
	// fullName = "Adriann Bimz" // ERROR : Cannot assign to fullName (Karena variabelnya const)
	fmt.Println(fullName)
}
