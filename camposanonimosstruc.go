package main

import "fmt"

type Animal struct {
	Color  string
	Nombre string
}

type Perro struct {
	// Animal, como atributo de Perro, es un atributo anonimo
	Animal // Nuestra estructura Perro, ahora tiene todos los atributos de la estructura Animal
}

func main() {
	d := Perro{}
	d.Color = "Negro"
	fmt.Printf("%s\n", d.Color)
}
