package responses

import (
	"math/rand"
	"time"
)

func QueryMomResponse(body *QueryStatusRequestBody, service string) string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	// generate random status
	status := statuses[r.Intn(len(statuses))]

	return "Dear " + body.Name + ", your " + service + " application is " + status + "."
}
