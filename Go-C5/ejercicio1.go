package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {

	productos := [][]string{
		{"id", "precio", "cantidad"},
		{"1", "$1300", "20"},
		{"2", "$500", "500"},
		{"3", "$3000", "100"},
	}

	f, err := os.Create("ejercicio1.csv")
	defer f.Close()

	if err != nil {
		log.Fatalln("Fallo al abrir el archivo", err)
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, producto := range productos {
		if err := w.Write(producto); err != nil {
			log.Fatalln("Fallo al escribir el archivo", err)
		}
	}
}
