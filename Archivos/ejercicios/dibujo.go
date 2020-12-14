package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

var img = image.NewNRGBA(image.Rect(0,0,800,800))
var col color.Color

func lineaHorizontal(x1, y, x2 int){
	for ; x1 <= x2; x1++{
		img.Set(x1, y ,col)
	}
}

func rectangulo(x1, y1, x2 ,y2 int){
	for; y1 <= y2 ; y1++{
		lineaHorizontal(x1, y1, x2)
	}
}

func cuadrado(x1,y1,l int){
	rectangulo(x1,y1,x1 + l, y1 + l)
}

func main(){
	
	c := 0

	for i := 0; i <= 8; i++{
		for j:=0; j <= 8; j++{			
			if c == 0{
				c = 1
				col = color.RGBA{255,255,255,255}//blanco
			} else {
				c = 0
				col = color.RGBA{0, 0, 0, 255} // negro
			}
			cuadrado(i * 100,j * 100,100)
		}
	}	

	f, err := os.Create("dibujo.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}