package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/hocus-tech/ada-go-example/internal/awswrapper"
	"github.com/hocus-tech/ada-go-example/internal/handler"
)

const streamNameEnvKey = "KINESIS_STREAM_NAME"

func main() {
	streamName := os.Getenv(streamNameEnvKey)
	dynamodbTableName := "go-aws-table"
	bucketName := os.Getenv("S3" + "_BUCKET")
	queueUrl := os.Getenv("SQS_URL")
	ctx := context.Background()
	wrapper, err := awswrapper.New(ctx, streamName, dynamodbTableName, bucketName, queueUrl)
	if err != nil {
		return
	}

	h := handler.NewHandler(wrapper)

	if err = http.ListenAndServe(os.Getenv("PORT"), h); err != nil {
		log.Fatal(err)
	}
}
