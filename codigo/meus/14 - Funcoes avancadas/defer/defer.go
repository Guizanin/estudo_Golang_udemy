package main

import "fmt"

func func1() {
	fmt.Println("Função 1...")
}

func func2() {
	fmt.Println("Função 2...")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado!")

	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {
	fmt.Println("Defer")

	defer func1()
	func2()

	fmt.Println(alunoEstaAprovado(6, 9))
}
