package main

import "fmt"

func main(){
	//TIPAGEM NORMAL
	var variavel1 string = "variavel 1"
	
	//INFERENCIA DE TIPO
	variavel2 := "variavel 2"
	
	//MULTIPLAS VARIAVEIS TIPADAS
	var(
		variavel3 string = "v3"
		variavel4 string = "v4"
	)

	//MULTIPLAS VARIAVEIS COM INFERENCIA DE TIPO
	variavel5, variavel6 := "v5", "v6"

	//variaveis
	fmt.Println(variavel1);
	fmt.Println(variavel2);
	fmt.Println(variavel3, variavel4);
	fmt.Println(variavel5, variavel6)

	//constantes
	const constante1 string = "const1"
	
	fmt.Println(constante1)

}