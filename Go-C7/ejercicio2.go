package main

import (
	"errors"
	"fmt"
)

func main2() {

	salary := 50000

	if salary < 100000 {
		fmt.Println(errors.New("el salario ingresado no alcanza el minimo imponible"))
	} else {
		fmt.Println("Debe pagar impuesto")
	}

}
