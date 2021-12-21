package main

import (
	"fmt"
)

func main3() {

	salary := 50000
	err := fmt.Errorf("el salario ingresado no alcanza el minimo imponible")

	if salary < 100000 {
		fmt.Printf("Error: %v", err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}

}
