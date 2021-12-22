package main

import (
	"fmt"
)

func main2() {

	promedio(2, 2, 3, 3, -1)
	promedio(2, 2, 3, 3)

}

func promedio(values ...int) int {
	var resultado int
	var promedio int
	for _, value := range values {
		if value <= 0 {

			fmt.Println("Un numero es negativo")
			return value
		} else {
			resultado += value
			promedio = resultado / len(values)

		}
	}
	fmt.Println("El promedio de los valores es :", (promedio))
	return promedio
}
