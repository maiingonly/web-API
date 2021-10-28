package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TopUPInput struct {
	// validation JSON required or not
	//if you want learn more about tag go https://pkg.go.dev/github.com/go-playground/validator/v10@v10.9.0#section-readme
	Name     string `json:"Name" binding:"required"`
	Amount   int    `json:"Amount" binding:"required,number"`
	WalletID string `json:"WalletID" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", pong)
	r.GET("/signup/:age", signupHandler)
	r.GET("/query", queryHandler)
	r.POST("/top-up", topupHandler)

	r.Run(":8888")
}

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
		"pang":    "pang",
		"pung":    "pung",
	})
}

func signupHandler(c *gin.Context) {
	age := c.Param("age")
	c.JSON(http.StatusOK, gin.H{
		"name":  "maiing",
		"email": "maiing@gmail.com",
		"age":   age,
	})
}

func queryHandler(c *gin.Context) {
	//query string parameters are parsed using the existing underlying request object.
	// the request responds to a url matching : /query?firstname=ucup&lastname=surucup
	firstname := c.Query("firstname")
	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")\

	c.JSON(http.StatusOK, gin.H{"firstname": firstname, "lastname": lastname})
}

func topupHandler(c *gin.Context) {
	var topUpInput TopUPInput

	err := c.ShouldBindJSON(&topUpInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Println(err)
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"WalletID": topUpInput.WalletID,
		"Name":     topUpInput.Name,
		"Amount":   topUpInput.Amount,
	})
}
