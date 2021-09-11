package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type product struct {
	Name         string `json:"name"`
	ImageURL     string `json:"imageURL"`
	Description  string `json:"description"`
	Price        string `json:"price"`
	TotalRatings string `json:"totalRatings"`
}

type productDocument struct {
	URL       string  `json:"url"`
	Timestamp string  `json:"timestamp"`
	Product   product `json:"product"`
}

func main() {
	router := gin.Default()
	router.POST("/persistProduct", persistRequest)
	router.GET("/products", getProducts)

	router.Run(":8081")
}

func persistRequest(c *gin.Context) {
	var Request productDocument

	// Call BindJSON to bind the received JSON
	if err := c.BindJSON(&Request); err != nil {
		return
	}
	persistData(&Request)
	c.JSON(http.StatusOK, "Success")
}

func getProducts(c *gin.Context) {
	Products := fetchProducts()
	fmt.Println("Products : ", Products)
	c.JSON(http.StatusOK, Products)
}
