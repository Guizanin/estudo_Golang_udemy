package main

import "fmt"

func main(){
	fmt.Println("Switch:")

	numero := 1;

	dia := diaSemana(numero)
	fmt.Println("Usando switch dia da semana: " + dia)

	mensagemDia(dia)
}

func diaSemana(numero int) string{
	/*
	Forma comum de utilizar switch
	*/
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	default:
		return "Numero invalido"
	}
}

func mensagemDia(dia string){
	/*
	Segunda forma de utilizar switch, colocando a condição no CASE
	pode ser feito com um chainOfResponsability
	*/
	switch{
	case dia == "Domingo":
		fmt.Println(dia + ", é o primeiro dia da semana")
	case dia == "Segunda":
		fmt.Println(dia + ", é o segundo dia da semana")
	default:
		fmt.Printf("Dia inválido")
	}
}