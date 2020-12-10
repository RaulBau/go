package main

import "fmt"

func main() {
	//Definicion de la variable
	var dato int = 10
	var cmp int = 15;

	if( dato < cmp ) {
		fmt.Printf("dato es menor que %d", cmp)
	 }
	 else{
		fmt.Printf("dato es mayor que %d", cmp)
	 }
}