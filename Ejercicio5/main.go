package main

import "fmt"

func main() {
	word := "hola mundo"
	fmt.Println("Word length:", len(word))

	for i, a := range word {
		fmt.Printf("%2d) %c \n", i, a)
	}

	/*
		a := make([]int, 5)
		fmt.Println(a, "\n")

		fmt-Println(a[1:4]);
	*/ /****CREA UN SLICE DE TAMAÑO 5 Y LOS COMPLETA CON EL VALOR NULO DEL TIPO DE DATO ****/
}
