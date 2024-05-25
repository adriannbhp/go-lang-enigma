package main

import "fmt"

func main() {

	jam := 25

	if jam >= 4 && jam <= 5 {
		fmt.Println("Bangun Pagi")
	} else if jam >= 6 && jam <= 7 {
		fmt.Println("Mandi dan Sarapan")
	} else if jam >= 8 && jam <= 11 {
		fmt.Println("Berangkat Sekolah")
	} else if jam == 12 {
		fmt.Println("Pulang Sekolah")
	} else if jam >= 13 && jam <= 15 {
		fmt.Println("Tidur Siang")
	} else if jam >= 16 && jam <= 17 {
		fmt.Println("Waktu Main")
	} else if jam >= 24 {
		fmt.Println("Hanya ada 24 jam dalam sehari")
	}

	if jam > 24 {
		fmt.Println("Hanya ada 24 jam dalam sehari")
	} else if jam == 4 || jam == 5 {
		fmt.Println("Bangun Pagi")
	} else if jam == 6 || jam == 7 {
		fmt.Println("Mandi dan Sarapan")
	} else if jam >= 8 && jam <= 11 {
		fmt.Println("Berangkat Sekolah")
	} else if jam == 12 {
		fmt.Println("Pulang Sekolah")
	}
}
