package main

import "fmt"

type pessoa struct {
	nome  string
	idade int8
	cpf   int64
	endereco
	escola
}

type endereco struct {
	logradouro string
	numero     int8
}

type escola struct {
	nomeEscola string
	endereco
}

func main() {

	endereco1 := endereco{logradouro: "Rua joão Tavora", numero: 59}

	enderecoEscola := endereco{logradouro: "Rua Martim Soares Moreno", numero: 80}

	escola1 := escola{nomeEscola: "Rubens Moreira da Rocha", endereco: enderecoEscola}

	pessoa1 := pessoa{
		nome:     "Wagner",
		idade:    20,
		cpf:      123456789,
		endereco: endereco1,
		escola:   escola1,
	}

	fmt.Println(pessoa1)

}
