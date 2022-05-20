package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {

	if nota := (n1 + n2) / 2; nota < 5 {
		return false
	}

	return true

}

func main() {
	defer funcao1()
	// Defer = Adiar

	funcao2()

	fmt.Println(alunoEstaAprovado(3, 3))
}
