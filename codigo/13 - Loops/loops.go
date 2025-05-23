package main

import (
	"fmt"
)

func main(){
	fmt.Println("Loops: ")

	// i := 0

	//FORMA 1
	// for i < 10 {
	// 	fmt.Println("Incrementando i: ", i)
	// 	i++
	// 	time.Sleep(time.Second)
	// }

	//FORMA 2
	// for j:= 0; j < 10; j++ {
	// 	fmt.Println("incremento forma 2 for: ", j)
	// }

	//FORMA 3
		//iterando em arrays
	arrayNomes := [3]string{"jão", "Davi", "Golias"}
	for indice, nome := range arrayNomes{
		fmt.Println("indice: ", indice, "; nome: ", nome)
	}

		//iterando strings
		palavra := "PALAVRA"
		for indice, letra := range palavra {
			fmt.Println("iterando palavra: ", indice, "codigo asc: ", letra, "| agora sim letra: ", string(letra) )
		}

		//iterando com map
		maps := map[string] string {
			"nome": "Guilherme",
			"sobrenome": "Zanin",
		}
		for chave, valor := range maps {
			fmt.Println("iterando sobre map: chave: ", chave, " valor: ", valor)
		}

}