package main

import (
	"fmt"
	"os"
)

func main1() {

	defer func() {
		fmt.Printf("Ejecucion finalizada \n")
	}()

	_, err := os.Open("customer.txt")
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}

}
