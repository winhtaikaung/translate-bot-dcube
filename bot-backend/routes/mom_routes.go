package routes

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var validity = []string{"expired", "valid"}
var passes = []string{"Employment Pass", "S Pass", "Work Permit", "Dependant's Pass", "Long-Term Visit Pass"}

func generateRandom(arr []string) string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return arr[r.Intn(len(arr))]
}

func generateRandomDate() string {
	min := time.Date(2023, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2028, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	date := time.Unix(sec, 0).String()
	return date[:10]
}

func QueryMOMResponse(body QueryStatusRequestBody, service string) string {
	var status string

	// generate random status
	pass := generateRandom(passes)
	validity := generateRandom(validity)

	if validity == "valid" {
		status = " is valid and expires on " + generateRandomDate() + "."
	} else {
		status = " has expired."
	}

	return "Dear " + body.Name + " of " + body.Identifier + ", your " + pass + status
}

func QueryMOMStatus(c *gin.Context) {
	var requestBody QueryStatusRequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, "User information not found!")
	}

	service := c.Param("service")

	switch service {
	case "check-pass":
		c.String(http.StatusOK, QueryMOMResponse(requestBody, service))
	default:
		c.String(http.StatusBadRequest, "Service not found!")
	}
}
