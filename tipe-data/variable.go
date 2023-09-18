package main

import "fmt"

func main(){
	// membuat variable
	var nama string
	nama="Relingga Aditya"
	fmt.Println(nama)
	var tanggal_lahir int32
	tanggal_lahir=19102004
	fmt.Println(tanggal_lahir)
	// gunakan var agar bisa tanpa mendeklarasikan variable nya
	var kuliah="telkomuniversity"
	println(kuliah)
	
	var umur=19
	print(umur)
	// ubah tipe datanya
	var umur_saya int8=19
	print(umur_saya)
	// jika tanpa menggunakan var
	country:="indonesia"
	print(country)
	country="america"
	print(country)
	// atau mutiple variable
	var(
		first_name="Relingga"
		last_name="Aditya"
	)
	fmt.Println(first_name)
	fmt.Println(last_name)
	// tipe data constant tidak bisa dirubah
	// tidak di pakai tidak masalah
	const saya="manusia"
	print(saya)
	// atau bisa juga multiple constant
	const(
		nama_depan="ucok"
		nama_belakang="jhonson"
	)
}	

