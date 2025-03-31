package main

import "fmt"

type pessoa struct{
	nome string
	idade int8
}

//herança só declara o type sozinho no struct que herda
type estudante struct{
	pessoa
	curso string
}

func main(){
	fmt.Println("Herança")

	e1 := estudante{
		pessoa{"Gui", 36}, "Sistemas de Info"}

	fmt.Println("estudante: ", e1)
	
}