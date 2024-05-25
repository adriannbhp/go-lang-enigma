package main

import "fmt"

func main() {
	var nama string
	fmt.Printf("Nama Lengkap : ")
	fmt.Scanf("%s", &nama)
	fmt.Println("Nama", nama, ". Alamat", nama)

	//fmt.Printf("Alamat : ")
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Scan()
	//homeAddress := scanner.Text()
	//fmt.Println("Alamat Anda \n ", homeAddress)

	// Menggunakan Os Argument

	//argsWithProg := os.Args
	//fmt.Println(argsWithProg[1])
	//fmt.Println(argsWithProg[2])
	//fmt.Println(argsWithProg[3])

	//fullNameArg := flag.String("Name", "", "User Name")
	//homeAddressArgs := flag.String("address", "", "User Address")
	//isGoldMember := flag.Bool("Goldm", false, "if user Gold Member")
	//
	//flag.Parse()
	//
	//fmt.Println("Full Name ", *fullNameArg)
	//fmt.Println("Home Address ", *homeAddressArgs)
	//fmt.Println("Gold Member ", *isGoldMember)
}
