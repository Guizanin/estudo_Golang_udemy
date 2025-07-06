package main

import "fmt"

func somar(n1 int8, n2 int8) int8{
	return n1 + n2
}

func main(){
	soma := somar(10, 3)

	fmt.Println(soma)

	var f = func(txt string){
		fmt.Println(txt)
	}

	//funcao pode retornar mais de um resultado
	var f2 = func(n1,n2 int8) (int8, int8){
		soma := n1 + n2
		subtracao := n1 - n2

		return soma, subtracao
	}

	var f3 = func(n1, n2 int8, txt string) (int8, string){
		return n1 + n2, txt
	}

	soma, subtracao := f2(10, 33)
	fmt.Println(soma, subtracao)

	soma2, texto := f3(55, 20, "texto qualquer")
	fmt.Println("soma2: ", soma2)
	fmt.Println("texto: ", texto)

	//para ignorar um dos retornos, usar _
	_ , texto2 := f3(55, 20, "texto qualquer")
	fmt.Println("text2: ", texto2)


	f("texto teste da funcao em variavel")
}