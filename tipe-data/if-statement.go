package main

import "fmt"

func main() {
	var kondisi="lapar"
	if kondisi=="lapar"{
		fmt.Println("Makanlah")
	}else{
		fmt.Println("mantap")
	}
	var ujian=10
	if ujian >=70{
		fmt.Println("KAMU LULUS")
	}else if ujian<=70{
		fmt.Println("Lulus Bersyarat")
	}else{
		fmt.Println("Tidak Lulus")
	}
	// if short statement
	nama:="Relingga"
	if length:=len(nama);length>5{
		fmt.Println("sudah betul")
	}else{
		fmt.Println("nama salah")
	}
	
}
