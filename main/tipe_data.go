package main

import "fmt"

/**
Tipe Data
*/

func main() {

	// Tipe Data Rune (int32)
	var karakter rune = 'A'
	fmt.Printf("Karakter %c ", karakter)

	// String Template (menggunakan ``)
	var strtemplatealamat string = `Jalan Sahari raya
    no.80
    Jakarta Pusat`
	fmt.Println(strtemplatealamat)

	var alamat string = "\n Jalan Sahari raya \n no.80 \n Jakarta Pusat "
	fmt.Print(alamat)

	var phi float64 = 3.14546457
	fmt.Printf("\nPhi %.2f", phi)

	var nama string = "Adrian"
	fmt.Printf("\nNama %s", nama)
}
