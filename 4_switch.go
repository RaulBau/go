package main

import "fmt"

func main() {
	x := 5
 
	switch x {
	case 1:
		fmt.Println("x es igual a 1")
	case 2:
		fmt.Println("x es igual a 2")
	default:
		// Valor por defecto, si no hay coincidencias
		fmt.Println("x es igual a", x)
	}

	/*
	var t interface{}
	
	switch t := t.(type) {
		default:
			fmt.Println("Tipo desconocido")
		case bool:
			fmt.Println("Booleano")
		case int:
			fmt.Println("Entero")
	}
	*/
}