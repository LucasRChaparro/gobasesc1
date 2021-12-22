package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type Cliente struct {
	legajo    int
	nombre    string
	apellido  string
	dni       string
	telefono  string
	domicilio string
}

func main() {
	//creo objeto cliente
	c1 := Cliente{
		nombre:    "Luis",
		apellido:  "Sarasa",
		dni:       "12345",
		telefono:  "12345",
		domicilio: "calle falsa 123",
	}
	// le asigno legajo
	c1.GenerarId()
	//quiero abrir archivo
	LeerArchivo()
	//quiero crear el archivo
	CrearArchivo(c1)
	// reviso si el cliente tiene datos cargados
	VerificarCliente(c1)
}

type MiError struct {
	Mensaje string
}

func (c *Cliente) GenerarId() {

	id := rand.Intn(30-10) + 10
	if id == 0 {
		panic("No se pudo generar el identificador")
	}
	c.legajo = id
}

func CrearArchivo(c1 Cliente) {
	d1 := []byte(fmt.Sprint("ID;Nombre;Apellido;Dni;Telefono;Domicilio\n",
		c1.nombre, ";", c1.apellido, ";", c1.dni, ";", c1.telefono, ";", c1.domicilio, ";", "\n",
	))
	err := os.WriteFile("./Archivo.csv", d1, 0644)
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func LeerArchivo() {
	// quiero leer el archivo
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Ejecución finalizada \n")
	}()
	_, err := os.Open("customers.txt")
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado\n") // sale de LeerArchivo()
	}
}

func VerificarCliente(c Cliente) (bool, error) {
	if c.legajo == 0 || c.telefono == "" {
		return true, errors.New("No puede tener valores en 0")
	} else if c.nombre == "" || c.apellido == "" || c.domicilio == "" {
		return true, errors.New("No puede dejar campos vacíos")
	}
	return false, nil
}
