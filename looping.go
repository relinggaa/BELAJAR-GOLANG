package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println("Looping ke",i) 
		i++
	}
	//for dengan statement
	
	for i := 1;  i <= 10; i++ {
		fmt.Println("Looping for statement ke",i) 
	}
	// looping slice
	slice:=[]string{"Relingga","Rose","Lisa"}

	for i=0; i < len(slice); i++{
		fmt.Println(slice[i])
	}
	// perulangan for range
	for i,value :=range slice{
		fmt.Println(i,value)
	}
	// ubah dengan underscore jika tidak ingin ada index
	for _,value :=range slice{
		fmt.Println(value)
	}
	// looping map
	data:=make(map[string]string)
	data["nama"]="Relingga"
	data["umur"]="19"
	
	for key,value := range data{
		fmt.Println(key,"=",value)
	}


}