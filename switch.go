package main

import "fmt"

func main() {
	nama := "Aditya"

	switch nama {
	case "Relingga":
		fmt.Println("Halo Relingga")
	case "Aditya":
		fmt.Println("Halo Aditya")
	default:
		fmt.Println("Kenalan Dong")
	}
	// short statement
	switch length:=len(nama); length > 5{
		case true:
			fmt.Println("sip")
		case false:
			fmt.Println("salah")

	}
	// switch tanpa expresion
	nilai:=100
	switch{
	case nilai ==100:
		fmt.Println("A")
	case nilai ==90:
		fmt.Println("B")
	case nilai ==80:
		fmt.Println("C")
	}
	
}
