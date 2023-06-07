package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func main() {
	lambda.Start(handler)
	log.Println("Surf Is Up Dude!")
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("surf is still up!!!!")

	var response = events.APIGatewayProxyResponse{
		StatusCode: 200,
	}
	return response, nil
}
