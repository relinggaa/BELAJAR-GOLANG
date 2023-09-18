package main

import "fmt"

func main() {
	// tentukan jumblah array
	var food [3]string
	food[0] = "geprek"
	food[1] = "rawon"
	food[2] = "betutu"
	

	// atau
	// jika jumblah array tidak tau gunakan [....]
	var minuman=[...]string{
		"es jeruk","esteh","marimas","milkshake","jus alpukat","jus jambu",
	}
	// membuat slice dan pointer nya
	var slice1=minuman[1:4]
	var semua =minuman[:]
	fmt.Println(slice1)
	fmt.Println(semua)
	// len mengukur panjang slice
	fmt.Println(len(slice1))
	// capasitas slice/array
	fmt.Println(cap(slice1))

	var slice2=minuman[1:]
	fmt.Println(slice2)
		// menambah data slice dengan benar
	var slice3=append(slice2,"Relingga")
	fmt.Println(slice3)
	// merubah slive
	slice3[1]="bukan marimas"
	fmt.Println(slice3)
	fmt.Println(minuman)
	// lansung membuat slice tanpa membuat array
	// len nya 2 cap,5
	newSlice:=make([]string,2,5)
	newSlice[0]="Relingga"
	newSlice[1]="Aditya"

	fmt.Println(newSlice)

	// copy slice
	copyslice:=make([]string,len(newSlice),cap(newSlice))
	copy(copyslice,newSlice)
	fmt.Println(copyslice)
}	