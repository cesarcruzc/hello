package main

import (
	"fmt"

	"./greet"
)

// Los paquetes son carpetas con una coleccion de archivos
func main() {
	fmt.Println("Greet")
	fmt.Println(greet.English())
	fmt.Println(greet.Italian())
	fmt.Println(greet.Spanish())

}
