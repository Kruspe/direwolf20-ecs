package main

import (
	"context"
	_ "embed"
	"encoding/base64"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

//go:embed img.png
var file []byte

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, _ events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	logger := log.New()
	level, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		log.Error("Could not get log level from environment variable", err)
	}
	logger.SetLevel(level)
	logger.SetFormatter(&log.JSONFormatter{})

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Error("Config could not be loaded", err)
		return events.LambdaFunctionURLResponse{StatusCode: http.StatusInternalServerError}, err
	}

	ecsClient := ecs.NewFromConfig(cfg)
	ecsClusterArn := os.Getenv("ECS_CLUSTER_ARN")
	ecsServiceArn := os.Getenv("ECS_SERVICE_ARN")
	_, err = ecsClient.UpdateService(ctx, &ecs.UpdateServiceInput{
		Cluster:      aws.String(ecsClusterArn),
		Service:      aws.String(ecsServiceArn),
		DesiredCount: aws.Int32(1),
	})
	if err != nil {
		log.Error("Could not scale up ecs service", err)
		return events.LambdaFunctionURLResponse{StatusCode: http.StatusInternalServerError}, err
	}

	return events.LambdaFunctionURLResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "image/png",
		},
		Body:            base64.StdEncoding.EncodeToString(file),
		IsBase64Encoded: true,
	}, nil
}
