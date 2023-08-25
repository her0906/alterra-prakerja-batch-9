package main

import (
	"fmt"
)

func prima(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func main() {
	var number int
	fmt.Print("Masukkan sebuah bilangan = ")
	fmt.Scanln(&number)

	if prima(number) {
		fmt.Println(number, "adalah bilangan prima.")
	} else {
		fmt.Println(number, "bukan bilangan prima.")
	}
}
