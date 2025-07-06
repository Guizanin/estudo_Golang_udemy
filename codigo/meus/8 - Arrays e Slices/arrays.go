package main

import "fmt"

//ARRAYS JÁ TEM TAMANHO PRÉ-FIXADOS EM SUA CRIAÇÃO
//SENDO NADA FLEXÍVEIS
func main(){
	fmt.Println("Arrays e Slices")
	
	//ARRAYS
	//nome [TAMANHO]TIPO -> obrigatório o tamanho
	var array1 [5]int
	fmt.Println("array1: ", array1)

	array1[0] = 99
	fmt.Println("array1 preenchido primeira posição: ", array1)

//----

  array2 := [5]string{"Guilherme", "Bruna", "Carol", "jane"}
	fmt.Println("array2: ", array2)

	//----

	//o tamanho é baseado na quatidade de entradas ao criar
  array3 := [...]int{1,2,3,4,5}
	fmt.Println("array3: ", array3)


	

}