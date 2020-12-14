package main

/*
	Pasar al programa 10 numeros como argumentos de la ejecucion.

	Los numeros se guardan en un arreglo.
	Se deberan de ordenar de menor a mayor y mostrarlos en pantalla (imprimir el arreglo).
	Se tiene que verificar si la cantidad de numeros es la correcta antes de ejecutar el programa.
*/

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	switch l := len(args); {
	case l > 10:
		fmt.Println("Ingresa un maximo de 10 numeros")
		return
	}

	var nums [10]int64

	// Se llena el arreglo con los numeros convertidos a entero
	for i, n := range args {
		n, err := strconv.ParseInt(n, 10, 64)
		if err != nil {
			//Ignora el numero y continua si no es valido
			continue
		}

		//Se guarda el numero en el arreglo
		nums[i] = n
	}

	//Se ordenan los numeros mediante el algoritmo de burbuja
	for range nums {
		for i, n := range nums {
			if i < len(nums)-1 && n > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}

	fmt.Println(nums)
}
