package main

import (
	"errors"
	"fmt"
	"os"
)

type myError struct{}

func (e myError) Error() string {
	return "Ocurrio Erorr"
}

func main() {

	a, err := salarioMensual(70, 10050)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Println(a)

	calcularAguinaldo(10000, 20000, 10000, 30000)

}

func salarioMensual(horas int, valor int) (calculo int, err error) {
	total := horas * valor
	e := fmt.Errorf("el trabajador no puede haber trabajado menos de 80 hrs mensuales")

	if horas < 80 {
		fmt.Println("Error", e)

	}

	if total > 150000 {
		total := total - (10 / 100)
		return total, errors.New("se le desconto el 10%")
	}

	return total, nil

}

func calcularAguinaldo(sueldos ...int) (aguinaldo int, err error) {

	var max int = sueldos[0]
	for i := 1; i < len(sueldos); i++ {

		if sueldos[i] < 0 {
			e1 := myError{}
			e2 := fmt.Errorf("e1: %w", e1)
			fmt.Println(errors.Unwrap(e1))
			fmt.Println(errors.Unwrap(e2))

		}

		if sueldos[i] > max {
			max = sueldos[i]
		}

	}
	aguinaldo = max / 2
	return aguinaldo, nil
}
