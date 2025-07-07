package main

import "fmt"


func main(){
	fmt.Println("Funções anonimas: ")

	retorno := func(texto string) string {
		return fmt.Sprintf("Retorno da função: %s", texto)
	}("Passando parametro")

	fmt.Println(retorno)

}