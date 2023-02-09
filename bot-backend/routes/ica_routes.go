package routes

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/translate-bot-dcube/bot-backend/utils"
)

type QueryStatusRequestBody struct {
	Name       string
	Identifier string
}

type SGArrivalCardRequestBody struct {
	Name            string
	Identifier      string
	Email           string
	Mobile          string
	Address         string
	DateOfArrival   string
	FullyVaccinated bool
	// healthDeclaration
}

var statuses = []string{"successful", "pending", "rejected"}

func QueryICAResponse(body QueryStatusRequestBody, service string) string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	// generate random status
	status := statuses[r.Intn(len(statuses))]

	return "Dear " + body.Name + " of " + body.Identifier +
		", your " + service + " application is " + status + "."
}

func QueryICAStatus(c *gin.Context) {
	var body QueryStatusRequestBody

	if err := c.BindJSON(&body); err != nil {
		utils.MakeBadRequestResponse(c, "ERR001", "User information not found!")
		return
	}

	fmt.Println("name and nric")
	fmt.Println(body.Identifier, body.Name)

	service := c.Param("service")

	switch service {
	case "entry-visa", "short-term-pass", "long-term-pass":
		utils.MakeOkResponse(c, QueryICAResponse(body, service))
	default:
		utils.MakeBadRequestResponse(c, "ERR002", "Service not found!")
	}
}

func GenerateSGArrivalCard(c *gin.Context) {
	var requestBody SGArrivalCardRequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		utils.MakeBadRequestResponse(c, "ERR001", "User information not found!")
	} else {
		utils.MakeOkResponse(c, "Your SG Arrival Card has been submittted successfully!")
	}
}
