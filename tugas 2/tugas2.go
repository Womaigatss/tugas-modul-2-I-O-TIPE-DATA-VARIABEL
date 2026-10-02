package main

import "fmt"

func main() {
	var  nama, nim, kelas, prodi string
	fmt.Print("masukan nama:")
	fmt.Scanln(&nama)
	fmt.Print("masukan nim:")
	fmt.Scanln(&nim)
	fmt.Print("masukan kelas:")
	fmt.Scanln(&kelas)
	fmt.Print("prodi:")
	fmt.Scanln(&prodi)

	fmt.Printf("Perkenalkan nama saya adalah %v dan saya adalah maba dari prodi %v nim saya adalah %v dan saya dari kelas %v", nama, prodi, nim, kelas)
}
