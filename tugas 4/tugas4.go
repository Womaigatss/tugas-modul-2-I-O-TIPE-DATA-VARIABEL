package main

import "fmt"

func main ()  {
	var fahrenheit,celcius int

	
	
    fmt.Print("masukkan suhu(F):  ")
	fmt.Scanln(&fahrenheit)

	celcius = (fahrenheit-32)*5/9

	fmt.Printf("jadi suhu adalah %v celcius",celcius)

}