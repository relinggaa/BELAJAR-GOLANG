package main

import "fmt"

func main() {
	var perjalanan = [...]string{
		"banyuwangi", "surabaya", "solo", "semarang", "yogyakarta", "cirebon", "bandung",
	}
	var slice = perjalanan[2:5]
	fmt.Println(cap(slice))
	fmt.Println(slice)
	var sliceappend=append(slice,"jakarta" )
	fmt.Println(sliceappend)

}
