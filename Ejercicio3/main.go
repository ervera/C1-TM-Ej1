package main

import "fmt"

func main() {
	/** var 1nombre string DEBE EMPEZAR CON STRING
	var apellido string
	var int edad ESTA AL REVEZ
	1apellido := 6 DEBE EMPEZAR CON STRING y EL TIPO DE DATO ES UN STRING y no tener WALRUS
	var licencia_de_conducir = true
	var estatura de la persona int SIN ESPACIOS
	cantidadDeHijos := 2 **/

	var nombre string
	fmt.Println(nombre)

	var apellido string
	fmt.Println(apellido)

	var edad int
	fmt.Printf("%d \n", edad)

	apellido = "6"
	fmt.Println(apellido)

	var licencia_de_conducir = true
	fmt.Println(licencia_de_conducir)

	var estaturaDeLaPersona int
	fmt.Printf("%d \n", estaturaDeLaPersona)

	cantidadDeHijos := 2
	fmt.Printf("%d \n", cantidadDeHijos)

	fmt.Println(nombre)
	fmt.Println(apellido)
	fmt.Println(licencia_de_conducir)
	fmt.Printf("%d \n", edad)
	fmt.Println(apellido)
	fmt.Printf("%b \n", estaturaDeLaPersona)
	fmt.Printf("%d \n", cantidadDeHijos)
}
