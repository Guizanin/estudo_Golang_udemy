package main

import (
	"errors"
	"fmt"
	pastaqualquer "modulosCriados/pastaQualquer"
)

func main(){
	// int int8 int16 int32 int64
	// int -> infere o tipo do int de acordo com a aquitetura do pc
	// uint -> unsigned int é um int sem sinal, somente positivo
	// rune (alias) quando utilizado para int que descrevem caracteres da tabela ask32
	// byte (alias int8)

	var n1 int = 1000
	var n2 byte = 200

	fmt.Println(n1)
	fmt.Println(n2)


	//float32 float64 

	// Não existe CHAR, o GOlang transforma o caractere para o numero correspondente
	// da tabela ASC, entre aspas simples
	char := 'B'
	fmt.Println(char)

	//Boleano
	var booleano1 bool = true
	var booleano2 bool

	fmt.Println(booleano1)
	fmt.Println(booleano2)

	//ERRO é um tipo em Golang
	var erro1 error

	var erro2 error = errors.New("novo erro")

	fmt.Println(erro1)
	fmt.Println(erro2)


	fmt.Println(pastaqualquer.Retorna("teste 2 parametros"))


}