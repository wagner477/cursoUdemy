package main

import (
	"errors"
	"fmt"
)

func main() {
	numero := -222222222222222
	fmt.Println(numero)

	var numero2 uint32 = 1000 // unsygned int
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE
	var numero3 rune = 12345
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 121212131.55
	fmt.Println(numeroReal2)

	numeroReal3 := 12131.44
	fmt.Println(numeroReal3)

	// Fim números reais

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	// Fim strings

	var texto string
	fmt.Println(texto)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
