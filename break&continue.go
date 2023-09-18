package main

import "fmt"

// break memperhentikan perulanagan
func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		if i==5{
			break
		}
	}
// skip perulangan
	for i := 2; i <= 5; i++ {
		fmt.Println(i)
		if i%2==0{
			continue
		}
	}
	
}
