package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type urlStruct struct {
	URL string `json:"url"`
}

func main() {
	router := gin.Default()
	router.POST("/scrapeProduct", scrapeRequest)

	router.Run(":8080")
}

func scrapeRequest(c *gin.Context) {
	var newUrl urlStruct

	// Call BindJSON to bind the received JSON
	if err := c.BindJSON(&newUrl); err != nil {
		return
	}

	Response, err := scrapeAmazonUrl(&newUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, Response)
	}

}
