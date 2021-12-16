package main

import "fmt"

func main2() {

	var precio float64 = 540
	var descuento float64 = 20

	var descuentodelproducto float64 = precio * (descuento / 100)

	fmt.Println(descuentodelproducto)
}
