package main

import "fmt"

func main1() {

	nombre := "Lucas"

	fmt.Println("cantidad de letras del nombre:", len(nombre))

	for i, letra := range nombre {
		fmt.Println(i+1, string(letra))
	}
}
