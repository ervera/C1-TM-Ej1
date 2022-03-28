package main

import "fmt"

func main() {
	var years int = 23
	var employed bool = true
	var employeeSeniority int = 3
	var salary int = 180000

	if years > 22 && employed && employeeSeniority > 0 {
		if salary > 100000 {
			fmt.Println("Su prestamo es aprobado y el mismo no posee intereses.")
		} else {
			fmt.Println("Su prestamo es aprobado pero el mismo posee intereses.")
		}
	} else {
		fmt.Println("Su prestamo no fue aprobado.")
	}

}

/*Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. Para ello tiene ciertas reglas para saber a
 qué cliente se le puede otorgar. Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más
 de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no les cobrará interés a los que su sueldo es mejor a $100.000.
Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.*/
