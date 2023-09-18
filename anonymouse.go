package main

import "fmt"

func main() {
	a := func(name string) string {
		return "halo" + name
	}
	fmt.Println(a("rel"))
}
