package main

import "fmt"

func main1() {

	s1 := Estudiante{
		Nombre:   "Lucas",
		Apellido: "Chapa",
		DNI:      33333333,
		Fecha:    "26-05-1991",
	}

	s2 := Estudiante{}

	s2.Nombre = "Roberto"
	s2.Apellido = "Carlos"
	s2.DNI = 22222222
	s2.Fecha = "12-05-1991"

	s1.detalle()
	s2.detalle()
}

type Estudiante struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (stud Estudiante) detalle() {
	fmt.Println("Nombre:", stud.Nombre, "\nApellido:", stud.Apellido, "\nDNI:", stud.DNI, "\nFecha:", stud.Fecha)
}
