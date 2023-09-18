package main

import "fmt"

func main() {
	// tentukan jumblah array
	var food [3]string
	food[0] = "geprek"
	food[1] = "rawon"
	food[2] = "betutu"
	
	fmt.Println(food[0])
	fmt.Println(food[1])
	fmt.Println(food[2])

	// atau
	// jika jumblah array tidak tau gunakan [....]
	var minuman=[...]string{
		"es jeruk","esteh","marimas","milkshake",
	}
	fmt.Println(minuman)
}
	