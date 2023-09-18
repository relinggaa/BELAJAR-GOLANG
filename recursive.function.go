package main

import "fmt"

func factorial(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}
// menggunakan factorial recursive
func factorialRecursive(value int)int{
	if value==1{
		return 1
	}else{
		return value *  factorialRecursive(value -1)
	}
}

func main() {
	fmt.Println(factorial(2*2))
	fmt.Println(factorialRecursive(2*2))
}
