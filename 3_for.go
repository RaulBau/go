package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ { 
		fmt.Println(i)
	}

	letras := []string{"A", "B", "C"}
	for key, value := range letras {
  		fmt.Println(key," - ",value)
	}

	for key := range letras {
  		fmt.Println(key)
	}

	/*
	//Ciclo infinito while
	for true { 
  		fmt.Println("Gopher")
	}

	//Ciclo infinito while
	for {
		fmt.Println("Gopher")
	}
	*/
}