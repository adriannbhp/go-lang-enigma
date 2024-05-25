package main

import "fmt"

/*
*
Pada tugas ini kamu diminta melengkapi coding berikut untuk menukar nilai variabel a dan b
lalu tampilkanlah pada console.

Contoh:

a = Camp, b = Enigma

maka nilai a menjadi Enigma dan b menjadi Camp.
*/

func main() {
	//var a, b string = "Camp", "Enigma"
	//
	//temp := a
	//a = b
	//b = temp
	//
	//fmt.Println(a)
	//fmt.Println(b)

	var a, b string
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	// Tulis kode disini
	temp := a
	a = b
	b = temp

	fmt.Print(a, b)
}
