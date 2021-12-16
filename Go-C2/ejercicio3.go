package main

import "fmt"

func main() {

	edad := 23
	tieneempleo := true
	sueldo := 80000

	if edad > 22 && tieneempleo {
		fmt.Println("Cumple con los requisitos")
		if sueldo > 100000 {
			fmt.Println("No se le cobra interes")
		} else {
			fmt.Println("Se le cobra interes")
		}
	} else {
		fmt.Println("No cumple con los requisitos")
	}

}
