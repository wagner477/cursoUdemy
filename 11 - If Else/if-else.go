package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("É maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outronumero := numero; outronumero > 0 {
		fmt.Println("Número é maior que 0")
	} else if numero < -10 {
		fmt.Println("Número é maior que -10")
	} else {
		fmt.Println("Entre 0 e -10")
	}

}
