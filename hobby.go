package main

import "fmt"

func main() {

	var name string
	var email string
	var hobby [2]string

	fmt.Println("Selamat Datang di Signup Account")
	fmt.Print("Masukkan Nama Anda :")
	fmt.Scan(&name)
	fmt.Print("Masukkan Hobby Anda :")
	fmt.Scan(&hobby[0])
	fmt.Print("Masukkan Hobby Anda :")
	fmt.Scan(&hobby[1])
	fmt.Print("Masukkan Email :")
	fmt.Scan(&email)

	fmt.Printf("Nama Anda adalah %s\n ", name)
	fmt.Println("Hobby Anda adalah ", hobby)
	for i := 0; i < len(hobby); i++ {
	}
	fmt.Println("Email Anda ", email)

}
