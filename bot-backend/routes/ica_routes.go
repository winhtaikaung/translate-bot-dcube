package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/translate-bot-dcube/bot-backend/responses"
	"github.com/translate-bot-dcube/bot-backend/utils"
)

func QueryICAStatus(c *gin.Context) {
	// var body QueryStatusRequestBody

	// if err := c.BindJSON(&body); err != nil {
	// 	utils.MakeBadRequestResponse(c, "ERR001", "User information not found!")
	// 	return
	// }

	// fmt.Println("name and nric")
	// fmt.Println(body.Identifier, body.Name)

	// service := c.Param("service")

	// switch service {
	// case "entry-visa", "short-term-pass", "long-term-pass":
	// 	utils.MakeOkResponse(c, QueryICAResponse(body, service))
	// default:
	// 	utils.MakeBadRequestResponse(c, "ERR002", "Service not found!")
	// }
}

func GenerateSGArrivalCard(c *gin.Context) {
	var requestBody responses.SGArrivalCardRequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		utils.MakeBadRequestResponse(c, "ERR001", "User information not found!")
	} else {
		utils.MakeOkResponse(c, "Your SG Arrival Card has been submittted successfully!")
	}
}
