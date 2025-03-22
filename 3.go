package main

import "fmt"

func fibolagifibolagi(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 6
	} else if n == 2 {
		return 7
	} else {
		return fibolagifibolagi(n-1) + fibolagifibolagi(n-2) + fibolagifibolagi(n-3)
	}
}

func main() {
	var masukan, hasil int
	fmt.Scan(&masukan)
	hasil = fibolagifibolagi(masukan)
	fmt.Println(hasil)
}
