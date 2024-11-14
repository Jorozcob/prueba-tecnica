/*
- crear endponts para la api
- consumir la api, https://randomuser.me/documentation y devolver 15000 registros.
- tratar los datos para que no se repitan, de los cuales solo tomar en el cual el uuid no debe de ser repetido.

    genero,
    primer nombre,
    primer apellido,
    email y
    uuid

la api debe de responder en menos de 2.25 segundos.
*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Name struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

type User struct {
	Gender string `json:"gender"`
	Name   Name   `json:"name"`
	Email  string `json:"email"`
	UUID   string `json:"uuid"`
}

type RandomUserResponse struct {
	Results []User `json:"results"`
}

var userMap = make(map[string]User)

func getUsuarios(c *gin.Context) {

	const totalusr = 150
	const lotes = 50

	start := time.Now()
	for i := 0; i < totalusr/lotes; i++ {
		url := fmt.Sprintf("https://randomuser.me/api/?results=%d", lotes)
		res, err := http.Get(url)
		if err != nil {
			log.Fatalf("Error al obtener datos: %v", err)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		if res.StatusCode != http.StatusOK {
			log.Fatalf("Respuesta fallida con código: %d y\ncuerpo: %s\n", res.StatusCode, body)
		}

		var apiResponse RandomUserResponse
		if err := json.Unmarshal(body, &apiResponse); err != nil {
			log.Fatalf("Error al parsear respuesta JSON: %v", err)
		}

		for _, user := range apiResponse.Results {
			if _, exists := userMap[user.UUID]; !exists {

				userMap[user.UUID] = user
			}
		}
	}

	users := []User{}
	for _, user := range userMap {
		users = append(users, user)
	}

	if len(users) > totalusr {
		users = users[:totalusr]
	}

	elapsed := time.Since(start)
	if elapsed.Seconds() > 2.25 {
		log.Printf("La respuesta tardó más de 2.25 segundos: %.2f segundos", elapsed.Seconds())
	}

	c.JSON(http.StatusOK, users)
}

func main() {
	router := gin.Default()

	router.GET("/usuarios", getUsuarios)

	router.Run(":8080")
}
