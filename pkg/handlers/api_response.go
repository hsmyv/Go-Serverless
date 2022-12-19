package handlers

import (
	"github.com/aws/aws-lambda-go/events"
	"encoding/json"


)

func apiResponse(status int, body interface{})(*events.APIGatewayProxyResponse, error){
	resp := event.APIGatewayProxyResponse{Headers: map[string]string["Content=Type": "application/json"]}

	stringBody, _ := json.Marshal(body)
	resp.Body = string(stringBody)
	return &resp, nil

}