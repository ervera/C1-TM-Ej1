package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var count int = 0

	for _, employee := range employees {
		if employee > 22 {
			count++
		}
	}

	fmt.Println("Hay", count, "empleados con edad mayor a 21")

	delete(employees, "Pedro")

	bs, _ := json.Marshal(employees)
	fmt.Printf("%v \n", string(bs))
}
