package main

import (
	"errors"
	"fmt"
)

func main() {

	valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorMaximo := maxFunc(1, 2, 3, 4, 1, 2, 4, 5)
	valorPromedio := promFunc(0)

	minFunc, error := estadisticas(valorMinimo)
	promFunc, error := estadisticas(valorPromedio)
	maxFunc, error := estadisticas(valorMaximo)

	if error != nil {
		fmt.Println("Hubo un error:", error)
	} else {
		fmt.Println("El numero minimo es:", minFunc)
		fmt.Println("El promedio es:", promFunc)
		fmt.Println("El numero maximo es:", maxFunc)
	}
}

func minFunc(values ...int) int {
	var min int = values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < min {
			min = values[i]
		}
	}

	return min
}

func maxFunc(values ...int) int {
	var max int = values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > max {
			max = values[i]
		}
	}

	return max
}

func promFunc(values ...int) int {
	var resultado int
	var promedio int

	for _, value := range values {
		resultado += value
		promedio = resultado / len(values)
	}

	return promedio
}

func estadisticas(calculo int) (int, error) {
	if calculo == 0 {
		return 0, errors.New("No se encontro valor asignado")
	} else {
		return calculo, nil
	}
}
