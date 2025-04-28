package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := 15

	if numero < 10 {
		fmt.Println("numero menor 10")
	} else {
		fmt.Println("numero maior 10")
	}

	if outroNumero := numero; outroNumero < 10 {
		fmt.Println("outro menor  10")
	} else {
		fmt.Println("outro maior 10")
	}
}
