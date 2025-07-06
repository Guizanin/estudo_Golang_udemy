package main

import "fmt"

/*
função com retorno nomeado
tipo inferido na declaracao do metodo não precisa inferir na atribuição
*/
func calculosMatematicas1(n1, n2 int) (soma, subtracao int){
	soma = n1 + n2
	subtracao = n1 - n2

	return
}
//mas pode ser feito do modo tradiconal tbm, sem nomear retorno
func calculosMatematicas2(n1, n2 int) (int, int){
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main(){
	nomeado, nomeado2 := calculosMatematicas1(1,2)
	semNomear, semNomear2 := calculosMatematicas2(12,11)

	fmt.Println("retorno nomeado: ", nomeado, nomeado2 )
	fmt.Println("retorno sem nomear: ", semNomear, semNomear2 )
}