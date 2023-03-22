package main

import "fmt"

func main() {

	var angkaPertama int
	fmt.Println("Masukan angka Pertama: ")
	fmt.Scanln(&angkaPertama)

	var angkaKedua int
	fmt.Println("Masukan angka Kedua: ")
	fmt.Scanln(&angkaKedua)

	var operator string
	fmt.Println("Masukan operator: (+,-,x,:)")
	fmt.Scanln(&operator)

	if operator == "+" {
		fmt.Println("hasilnya adalah ", angkaPertama + angkaKedua)
	}else if operator == "-" {
		fmt.Println("hasilnya adalah", angkaPertama - angkaKedua)
	}else if operator == "x" {
		fmt.Println("hasilnya adalah", angkaPertama * angkaKedua)
	}else if operator == ":" {
		fmt.Println("hasilnya adalah", angkaPertama / angkaKedua)
	}else{
		fmt.Println("Kalkulator Rusak")
	}	
}