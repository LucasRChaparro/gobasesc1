package main

import "fmt"

func main1() {
	u1 := Usuario{
		nombre:     "Roberto",
		edad:       25,
		correo:     "roberto@gmail.com",
		contraseña: "ad2kklim",
	}

	fmt.Printf("%+v\n", u1)

	cambiarNombre(&u1)
	cambiarEdad(&u1)
	cambiarCorreo(&u1)
	cambiarContraseña(&u1)

	fmt.Printf("%+v\n", u1)
}

type Usuario struct {
	nombre     string
	edad       int
	correo     string
	contraseña string
}

func cambiarNombre(usuario *Usuario) {
	usuario.nombre = "Lucas"

}
func cambiarEdad(usuario *Usuario) {
	usuario.edad = 31

}
func cambiarCorreo(usuario *Usuario) {
	usuario.correo = "lucas@gmail.com"

}
func cambiarContraseña(usuario *Usuario) {
	usuario.contraseña = "nu3v4p455"

}
