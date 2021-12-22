package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type Productos struct {
	Id        string `json:"id"`
	Nombre    string `json:"nombre"`
	Color     string `json:"color"`
	Precio    int    `json:"precio"`
	Stock     int    `json:"stock"`
	Codigo    string `json:"codigo"`
	Publicado bool   `json:"publicado"`
	Fecha     string `json:"fecha"`
}

func main() {
	r := gin.Default()

	r.GET("/bienven/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "Hola " + name,
		})
	})

	datosJson, err := ioutil.ReadFile("producto.json")
	if err != nil {
		log.Fatal(err)
	}

	productos := Productos{}
	err = json.Unmarshal(datosJson, &productos)
	if err != nil {
		log.Fatal(err)
	}

	r.GET("/productos", func(c *gin.Context) {
		c.JSON(200, productos)
	})

	r.Run(":8080")
}
