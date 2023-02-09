package responses

import (
	"math/rand"
	"time"
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
var how_to_service = []string{"short-term-visit", "long-term-visit", "entry"}

func QueryICAResponse(body *QueryStatusRequestBody, service string) string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	// generate random status
	status := statuses[r.Intn(len(statuses))]

	return "Dear " + body.Name + " of " + body.Identifier +
		", your " + service + " application is " + status + "."
}

func QueryICAHowToResponse(body *QueryStatusRequestBody, service string) string {

	// generate random status
	resp := map[string]string{
		"short-term-visit": "Hi " + body.Name + ", you may proceed to  https://www.ica.gov.sg/enter-depart/extend_short_stay",
		"long-term-visit":  "Hello " + body.Name + ", you may proceed to https://www.ica.gov.sg/reside/LTVP/apply ",
		"entry":            "Hi " + body.Name + ", you may proceed to https://www.ica.gov.sg/enter-transit-depart/entering-singapore/sg-arrival-card ",
	}
	return resp[service]

}
