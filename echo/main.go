package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)
	return events.APIGatewayProxyResponse{Body: request.Body, Headers: map[string]string{"Access-Control-Allow-Origin": "*"}, StatusCode: 200}, nil
}

func factorial(number int) int {
	if number == 0 {
		return 0
	}
	return (number * factorial(number-1))
}

func main() {
	lambda.Start(Handler)
}
