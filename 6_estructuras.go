package main

import "fmt"

type Libro struct {
   titulo string
   autor string
   categoria string
   id int
}

func main() {
   var Libro1 Libro    /* Declare Libro1 of type Libro */
   var Libro2 Libro    /* Declare Libro2 of type Libro */
 
   /* definición de Libro 1 */
   Libro1.titulo = "Libro 1"
   Libro1.autor = "Autor 1"
   Libro1.categoria = "Programación"
   Libro1.id = 123

   /* definición de Libro 2 */
   Libro2.titulo = "Go"
   Libro2.autor = "	Robert Griesemer, Rob Pike y Ken Thompson"
   Libro2.categoria = "Programación"
   Libro2.id = 124
 
   imprimeLibro(Libro1)

   imprimeLibro(Libro2)
}

func imprimeLibro( Libro Libro ) {
   fmt.Printf( "Libro titulo : %s\n", Libro.titulo);
   fmt.Printf( "Libro autor : %s\n", Libro.autor);
   fmt.Printf( "Libro categoria : %s\n", Libro.categoria);
   fmt.Printf( "Libro id : %d\n\n", Libro.id);
}