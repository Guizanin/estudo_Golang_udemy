package main

import "fmt"

type usuario struct {
	nome  string
	idade int
}

func (u usuario) pegaNome() {
	fmt.Printf("O nome do usuario recuperado é: %s, idade: %d \n", u.nome, u.idade)
}
func (u *usuario) fazerAniversario() {
	u.idade++
}

/*
Não tem suporte a metodos dentro do struct como fazemos em OO,
porém referencia o struct como parametro da func e usa ele dentro da funcao
*/
func main() {
	usuario1 := usuario{"Guilherme Zanin", 37}
	usuario2 := usuario{"Bruna", 27}

	usuario1.pegaNome()
	usuario1.fazerAniversario()

	usuario2.pegaNome()

	fmt.Println("Usuario 1 fez aniversário::")
	usuario1.pegaNome()
}
