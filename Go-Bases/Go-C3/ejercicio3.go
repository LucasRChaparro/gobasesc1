package main

import "fmt"

func main3() {
	calcularSueldo(50, "A")

}

func calcularSueldo(minutos float64, categoria string) float64 {

	var horas float64 = minutos / 60

	switch categoria {
	case "C":
		var salario float64 = horas / 1000
		fmt.Println("El sueldo del empleado es:", salario)
		return salario
	case "B":
		var salario float64 = (horas * 1500) * 1.20
		fmt.Println("El sueldo del empleado es:", salario)
		return salario
	case "A":
		var salario float64 = (horas * 3000) * 1.50
		fmt.Println("El sueldo del empleado es:", salario)
		return salario
	}
	return 0
}
