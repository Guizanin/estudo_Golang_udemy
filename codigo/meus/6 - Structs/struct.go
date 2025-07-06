package main

import "fmt"

type endereco struct{
	logradouro string
	numero int8
}

type usuario struct{
	nome string
	idade int8
	endereco endereco
}

func main(){
	fmt.Println("Arquivos struct")

	var u1 usuario
	u1.nome = "Guilherme"
	u1.idade = 36

	fmt.Println("struct u1 ", u1)


	u2 := usuario{"Bruna", 26, endereco{"Rua tal", 55}}
	fmt.Println("struct u2 ", u2)

	u3 := usuario{nome: "Bruna"}
	fmt.Println("struct u3 ", u3)

}