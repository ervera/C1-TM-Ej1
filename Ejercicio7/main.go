package main

import "fmt"

func main() {

	var selectedMounth int = 0
	var months = map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio",
		7: "Julio", 8: "Agosto", 9: "Septiembre", 10: "Octubre", 11: "Noviembre", 12: "Diciembre"}

	month, ok := months[selectedMounth]

	if ok {
		fmt.Printf("%d) %v \n", selectedMounth, month)
	} else {
		fmt.Printf("mes ingresado incorrecto \n")
	}
}
