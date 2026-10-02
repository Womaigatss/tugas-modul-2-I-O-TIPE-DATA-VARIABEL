package main

import "fmt"

func main() {
	var jarijari, luas float64
	fmt.Print("masukan jari jari lingkaran (r):")
	fmt.Scanln(&jarijari)

	luas = 3.14 * (jarijari * jarijari)

	fmt.Printf("jadi luas lingkaran adalah :%v", luas)

}
