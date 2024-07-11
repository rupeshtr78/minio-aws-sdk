package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	ListBuckets()
}

func ListBuckets() {
	// Create a new session for minio
	// use profile credentials using endpoint url
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "minio",
		Config: aws.Config{
			Region:           aws.String("us-west-1"),
			Endpoint:         aws.String("http://10.0.0.213:9000"),
			S3ForcePathStyle: aws.Bool(true),
			Credentials:      credentials.NewSharedCredentials("", "minio"),
		},
	})
	if err != nil {
		fmt.Printf("Error creating session: %v", err)
		return
	}

	// Create S3 service client
	svc := s3.New(sess)

	// Get the list of buckets
	result, err := svc.ListBuckets(nil)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Success", result)
}
