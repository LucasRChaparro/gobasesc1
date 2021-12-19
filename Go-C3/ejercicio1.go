package main

import "fmt"

func main1() {

	calcularImpuesto(40000)
	calcularImpuesto(55000)
	calcularImpuesto(151000)

}

func calcularImpuesto(sueldo float64) {

	if sueldo > 50000 {
		var impuesto float64 = sueldo * 0.17
		if sueldo > 150000 {
			impuesto += (sueldo - impuesto) * 0.10
			fmt.Println("El impuesto del empleado es de:", impuesto)
			return
		}
		fmt.Println("El impuesto del empleado es de:", impuesto)

	}
	return
}
