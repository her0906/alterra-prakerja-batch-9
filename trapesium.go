package main

import (
	"fmt"
)

func trapesium(alas1, alas2, tinggi float64) float64 {
	return 0.5 * (alas1 + alas2) * tinggi
}

func main() {
	var alas1, alas2, tinggi float64

	fmt.Print("Masukkan panjang alas 1 = ")
	fmt.Scanln(&alas1)

	fmt.Print("Masukkan panjang alas 2 = ")
	fmt.Scanln(&alas2)

	fmt.Print("Masukkan tinggi trapesium = ")
	fmt.Scanln(&tinggi)

	luas := trapesium(alas1, alas2, tinggi)
	fmt.Printf("Luas trapesium adalah: %.2f\n", luas)
}
