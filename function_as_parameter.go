package main

import "fmt"

func salam(nama string,filter func(string)string){
	sensor := filter(nama)
	fmt.Println("halo",sensor)
}
func sensor(nama string)string{
	if nama=="anjing"{
		return "****"
	}else{
		return nama
	}
}

func main() {
	salam("r",sensor)
	salam("anjing",sensor)
}
