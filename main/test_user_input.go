package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//nameArgs := flag.String("name", "", "FullName")
	//addressArgs := flag.String("alamat", "", "Address")
	//emailArgs := flag.String("email", "", "Email")
	//
	//flag.Parse()
	//
	//fmt.Println("Halo, saya", *nameArgs, "Saya tinggal di", *addressArgs, "Alamat email saya adalah", *emailArgs)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nama := scanner.Text()
	scanner.Scan()
	alamat := scanner.Text()
	scanner.Scan()
	email := scanner.Text()
	fmt.Printf("Halo, saya %s. Saya tinggal di %s. Alamat email saya adalah %s.", nama, alamat, email)

}
