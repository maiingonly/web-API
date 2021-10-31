package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TopUPInput struct {
	// validation JSON required or not
	//if you want learn more about tag go https://pkg.go.dev/github.com/go-playground/validator/v10@v10.9.0#section-readme
	Name     string `json:"Name" binding:"required"`
	WalletID string `json:"WalletID" binding:"required"`
	Amount   int    `json:"Amount" binding:"required,number,max=1000000,min=10000"`
}

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
		"pang":    "pang",
		"pung":    "pung",
	})
}

func SignupHandler(c *gin.Context) {
	age := c.Param("age")
	c.JSON(http.StatusOK, gin.H{
		"name":  "maiing",
		"email": "maiing@gmail.com",
		"age":   age,
	})
}

func QueryHandler(c *gin.Context) {
	//query string parameters are parsed using the existing underlying request object.
	// the request responds to a url matching : /query?firstname=ucup&lastname=surucup
	firstname := c.Query("firstname")
	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")\

	c.JSON(http.StatusOK, gin.H{"firstname": firstname, "lastname": lastname})
}

func TopupHandler(c *gin.Context) {

	var topUpInput TopUPInput

	err := c.ShouldBindJSON(&topUpInput)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			//from here you can create your own error messages in whatever language you wish return

			errorMessage := fmt.Sprintf("max: %s, %s", e.Param(), e.Error())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"WalletID": topUpInput.WalletID,
		"Name":     topUpInput.Name,
		"Amount":   topUpInput.Amount,
	})
}
