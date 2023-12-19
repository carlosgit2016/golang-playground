/*

Publish routes

/home
/about
/dashboard

Handle headers

GET HTTP 1.1 Content-Type=application/json

Handle POST / PUT / DELETE
Handle params as propertities
Handle query

/home GET
/about GET
/dashboard GET
/create/:person POST
/update/:person PUT
/delete/:person DELETE

/create?person=:person POST
/update?person=:person PUT
/delete?person=:person DELETE


Handle different types of body request

query / www-form-x-url / form data

*/

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type person struct {
	name    string `json:"name"`
	surname string `json:"surname"`
	age     int    `json:"age"`
}

func (ref *person) getName() string {
	return ref.name
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postRoutes(c *gin.Context, router *gin.Router) {
	router.POST("")
	router.POST("")
	router.POST("")
	router.POST("")
}

func postGenericHandler(c *gin.Context) {
	var person person

	/*
		{
			"name": "felps",
			"surname": "felps",
			"age": 15
		}
	*/
	if err := c.BindJSON(&person); err != nil {
		return
	}

	fmt.Print(person)
	c.IndentedJSON(http.StatusCreated, person)

}

func main() {
	router := gin.Default()
	router.GET("/home", getAlbums)

	router.POST("/create/person", postGenericHandler)
	router.POST("/create?person=:person")

	router.PUT("/update?person=:person")
	router.PUT("/update/:person")

	router.Run("localhost:8080")
}
