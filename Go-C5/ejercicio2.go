package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	f, err := os.Open("ejercicio2.csv")

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)

	for {
		productos, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)

		}

		fmt.Println(productos[0], "\t", productos[1], "\t", productos[2])

	}

}
