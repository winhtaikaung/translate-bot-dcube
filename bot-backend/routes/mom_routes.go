package routes

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/translate-bot-dcube/bot-backend/responses"
	"github.com/translate-bot-dcube/bot-backend/utils"
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

func QueryMOMResponse(body responses.QueryStatusRequestBody, service string) string {
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
	var requestBody responses.QueryStatusRequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		utils.MakeBadRequestResponse(c, "ERR001", "User information not found!")
		return
	}

	service := c.Param("service")

	switch service {
	case "check-pass":
		utils.MakeOkResponse(c, QueryMOMResponse(requestBody, service))
	default:
		utils.MakeBadRequestResponse(c, "ERR002", "Service not found!")
	}
}
