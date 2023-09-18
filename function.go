package main

import (
	"fmt"
)

func helloworld(){
	fmt.Println("Hello World")
}
func salam(nama string){
	fmt.Println("Selamat Siang",nama)
}
func tambah(number,numberb int8)int8{
	result:=number+numberb
	return result
}
func welcome(nama string)string{
	if nama==""{
		return "Selamat Datang bro"
	}else{
		return "Selamat Datang"+nama
	}
}
// returning mutiple values
func makanan()(string,string,string){
	return "Ayam Geprek","Gurami","sate"
}
// named return value
func minuman()(minuman1 string,minuman2 string,minuman3 string){
	minuman1="esteh"
	minuman2="esjeruk"
	minuman3="esdegan"
	return
}
// Variadic Function/parameter menemrima banyak data
func jumblah(angka ...int)int{
	total:=0
	for _,nomer:= range angka{
		total+=nomer
	}
	return total
}
// function value
func goodbye(nama string)string{
	return "good bye"+nama
}

func main(){
	helloworld()
	salam("relingga")
	fmt.Println(tambah(1,2))
	fmt.Println(welcome(""))

	ayam,sate,gurami:=makanan()
	fmt.Println(ayam)
	fmt.Println(sate)
	fmt.Println(gurami)
	// return 1 data
	ayam,_,_=makanan()
	fmt.Println(ayam)

	minuman1,minuman2,minuman3:=minuman()
	fmt.Println(minuman1)
	fmt.Println(minuman2)
	fmt.Println(minuman3)
	// vardic function
	
	// menggunakan slice
	slice:=[]int{
		10,20,
	}
	fmt.Println(jumblah(slice...))
	// function value
	bye:=goodbye
	fmt.Println(bye(" r"))
}



