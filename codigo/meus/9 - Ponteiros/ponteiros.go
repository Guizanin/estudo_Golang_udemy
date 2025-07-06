package main

import "fmt"

func main() {
	fmt.Println("____Ponteiros")

	numero := 10
	//para criar o ponteiro & antes da variavel alvo
	ponteiro := &numero
	//para ver o valor deste ponteiro * antes de usar a variavel ponteiro
	//da-se o nome de desreferenciacao
	fmt.Println(numero, *ponteiro)
}
