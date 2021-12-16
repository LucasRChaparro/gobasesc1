package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println("La edad de Bejanmin es:", employees["Benjamin"])

	count := 0
	for i, employee := range employees {

		if employee > 21 {
			count++
			fmt.Println(i, "tiene mas de 21")
		}
	}

	fmt.Println("Hay:", count, "empleados mayores a 21 años")

	employees["Federico"] = 25
	delete(employees, "Pedro")

	fmt.Println(employees)

}
