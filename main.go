package main

import (
	"fmt"
)

func main() {
	//variable
	var name string = "john kennedy"
	lastname := "wick"
	fmt.Println("Hello!", name, lastname)

	// constanta
	_ = "mmangiyeeloo"
	const dawir string = "orang gila"
	fmt.Println(dawir)

	// Multiple Variable
	var sempor, gombong, buayan string = "hildan", "rendy", "akmal"
	fmt.Println(sempor, gombong, buayan)

	// Type Declaration
	type ktp string
	var nomer ktp = "11111111111111"
	fmt.Println(nomer)

	//  Konversi Nilai
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	//
	var a = 10
	var b = 20
	var c = a * b
	fmt.Println(c)

	// Augmented Assignment
	var i = 18
	i += 20        // i = i + 20
	fmt.Println(i) //result -2

	var g = 18
	g -= 20        // i = i -20
	fmt.Println(g) //result -2

	//unary operator
	g++
	fmt.Println(g)

	// perbadningan

	if c == 23 {
		fmt.Println("Hasil sama")
	} else {
		fmt.Println("haisl tidak sama")
	}

	if c <= 200 {
		fmt.Println("Hasil Lebih Dari 23")
	} else {
		fmt.Println("Hasil Kurang Dari 23")
	}

	// berhenti di 11 PZN

}
