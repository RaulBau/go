package main5

import "fmt"

func main() {
	/*Definicion de variables*/
	var a int = 100
	var b int = 200
	var ret int
	
	//Llamando la función para obtener el valor maximo
	ret = max(a, b)

	fmt.Printf( "El valor maximo es: %d\n", ret )

	fmt.Println(a, b)

	//Lamando a la funcion de intercambio de datos
	a, b = cambio(a, b)

	fmt.Println(a, b)
}

//Funcion que retorna el maximo de 2 numeros
func max(n1, n2 int) int {
	//Declaracion de variables
	var res int
 
	if (n1 > n2) {
	   res = n1
	} else {
	   res = n2
	}
	return res 
}

//Funcion para intercambiar valores
func cambio(x, y int) (int, int) {
	return y, x
}