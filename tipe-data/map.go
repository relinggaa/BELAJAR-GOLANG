package main

import "fmt"

func main() {
	var person=map[string]string{
		"nama":"Relingga",
		"address":"Bandung",
	}
	// menambahkan data
	person["umur"]="19"
	fmt.Println(person["nama"])
	fmt.Println(person["address"])
	fmt.Println(person["umur"])

	// membuat map tanpa data
	var Relingga map[string]string=make(map[string]string)
	Relingga["nama_lengkap"]="Relingga Aditya"
	fmt.Println(Relingga)
}
