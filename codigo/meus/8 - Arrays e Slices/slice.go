package main

import "fmt"

func main() {
	fmt.Println("Slices: ")

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice)

	//add mais coisa no slice, por não saber a posicao final usa append
	slice = append(slice, 259)
	fmt.Println(slice)

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice2 := array[1:3]
	fmt.Println("slice2: ", slice2)

}
