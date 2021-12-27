// package main

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type Productos struct {
// 	Id        string `json:"id"`
// 	Nombre    string `json:"nombre"`
// 	Color     string `json:"color"`
// 	Precio    int    `json:"precio"`
// 	Stock     int    `json:"stock"`
// 	Codigo    string `json:"codigo"`
// 	Publicado bool   `json:"publicado"`
// 	Fecha     string `json:"fecha"`
// }

// func main() {
// 	r := gin.Default()

// 	r.GET("/bienven/:name", func(c *gin.Context) {
// 		name := c.Param("name")
// 		c.JSON(200, gin.H{
// 			"message": "Hola " + name,
// 		})
// 	})

// 	datosJson, err := ioutil.ReadFile("producto.json")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	productos := Productos{}
// 	err = json.Unmarshal(datosJson, &productos)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	r.GET("/productos", func(c *gin.Context) {
// 		c.JSON(200, productos)

// 	})

// 	r.GET("producto/:id", func(c *gin.Context) {
// 		id := c.Query("id")
// 		for _, p := range id {
// 			if productos.Id == "id" {
// 				c.JSON(http.StatusOK, p)
// 				return
// 			}
// 		}
// 		c.JSON(http.StatusNotFound, gin.H{"message": "producto no encontrado"})
// 	})

// 	r.Run(":8080")
// }
