package main

import "fmt"

func main() {
	// membuat alias
	type ktp string
	var nik ktp = "41424234324"
	fmt.Println(nik)
}
