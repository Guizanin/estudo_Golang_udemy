package main

import "fmt"

func main() {
	fmt.Println("arrays internos: ")

	//usa make ele aloca espaco na memoria
	//1 parametro -> tipo a ser criado
	//2 parametro -> quantidade ocupado
	//3 paramtro -> quantidade maxima**
	slice1 := make([]int, 3, 5)
	fmt.Println("slice1: ", slice1)
	fmt.Println("tamanho:", len(slice1))
	fmt.Println("capacidade:", cap(slice1))

	fmt.Println("modificando o slice1 -------------")
	//Golang reconhece que vai estourar o tamanho do slice e dobra usa capacidade
	slice1 = append(slice1, 10, 30, 40, 50)
	fmt.Println("slice1: ", slice1)
	fmt.Println("tamanho:", len(slice1))
	fmt.Println("capacidade:", cap(slice1))

}
