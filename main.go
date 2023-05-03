package main

import "fmt"

func main() {

	// var judul string = "judul"
	// var category [3]string
	// var sinopsis string = "sinopsis"

	// fmt.Println("Selamat Datang Di Komikmu")
	// fmt.Print("Masukkan Judul Yang Ingin Dicari :")
	// fmt.Scan(&judul)
	// fmt.Print("Masukkan Category Yang Ingin Dicari :")
	// fmt.Scan(&category[0])
	// fmt.Print("Masukkan Category Yang Ingin Dicari :")
	// fmt.Scan(&category[1])
	// fmt.Print("Masukkan Category Yang Ingin Dicari :")
	// fmt.Scan(&category[2])
	// fmt.Print("Sinopsis :")
	// fmt.Scan(&sinopsis)

	// fmt.Printf("Judul yang dicari %s\n ", judul)
	// fmt.Println("Kateogri yang di ingingkan", category)
	// for i := 0; i < len(category); i++ {
	// 	fmt.Println(category[i])
	// }
	// fmt.Println("Isi Sipnosis", sinopsis)

	var choice int

	fmt.Println("Selamat Datang DI Komikmu")
	fmt.Print("Masukkan pilihan Anda (1-5): ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Rimuru")
		fmt.Println("Genre : fantasy, action, romance")
	} else if choice == 2 {
		fmt.Println("Luffy")
		fmt.Println("Genre : fantasy, action, comedy")
	} else if choice == 3 {
		fmt.Println("Naruto")
		fmt.Println("Genre : fantasy, action, romance, comedy")
	} else if choice == 4 {
		fmt.Println("Power Rangers")
		fmt.Println("Genre : action, Mecha, School")
	} else if choice == 5 {
		fmt.Println("Ultraman	")
		fmt.Println("Genre : action , Mecha, Monster")
	} else {
		fmt.Println("Pilihan tidak tersedia")
	}

}
