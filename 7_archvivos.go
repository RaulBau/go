package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	creaArchivo()
	leeArchivo()
}

func creaArchivo() {
	/*
		Se crea el archivo, si hay un
		error se notifica
	*/
	arch, err := os.Create("prueba.txt")

	if err != nil {
		log.Fatalf("Error al crear el archivo: %s\n", err)
	}

	// defer ayuda a cerrar el archivo despues de que se escribe
	defer arch.Close()

	// long guarda la longitud del archivo en bytes
	long, err := arch.WriteString("Este es un archivo de pruebas para el sistema de archivos de GO.")

	if err != nil {
		log.Fatalf("Error al escribir en el archvivo: %s\n", err)
	}

	// Name() regresa el nombre del archivo
	fmt.Printf("Nombre del archivo: %s\n", arch.Name())
	fmt.Printf("Longitud en bytes: %d bytes\n", long)
}

func leeArchivo() {
	fileName := "prueba.txt"

	//Se leen los datos del  archivo
	datos, err := ioutil.leeArchivo("prueba.txt")
	if err != nil {
		log.Panicf("Error leyendo los datos del archivo: %s", err)
	}
	fmt.Printf("Nombre del archivo: %s\n", fileName)
	fmt.Printf("Tamaño: %d bytes\n", len(datos))
	fmt.Printf("Datos: %s\n", datos)
}
