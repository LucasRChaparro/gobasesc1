package main

import "fmt"

type miError struct {
	msg string
}

func main1() {

	salary := 50000
	err := &miError{msg: "el salario ingresado no alcanza el minimo imponible"}

	if salary < 100000 {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}

}

func (e *miError) Error() string {
	return fmt.Sprintf("%v", e.msg)
}
