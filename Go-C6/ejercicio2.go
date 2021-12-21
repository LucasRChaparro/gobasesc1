package main

import "fmt"

func main2() {

	producto1 := nuevoProducto("Torta", 3000)
	producto2 := nuevoProducto("Facturas", 1000)

	usuario1 := new(User)
	usuario1.nombre = "Lucas"
	usuario1.apellido = "Chapa"
	usuario1.correo = "lucas@gmail.com"

	fmt.Printf("%+v\n", producto1)
	fmt.Printf("%+v\n", producto2)
	fmt.Printf("%+v\n", usuario1)

}

type User struct {
	nombre    string
	apellido  string
	correo    string
	productos []Producto
}

type Producto struct {
	nombre   string
	precio   int
	cantidad int
}

func nuevoProducto(nombre string, precio int) *Producto {
	p := new(Producto)
	p.nombre = nombre
	p.precio = precio

	return p

}
