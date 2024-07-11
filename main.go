package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	client, err := GetS3Client("minio")
	if err != nil {
		panic(err)
	}
	buckets, err := ListBuckets(client)
	if err != nil {
		panic(err)
	}
	fmt.Println(buckets)
}

func ListBuckets(client *s3.S3) (*s3.ListBucketsOutput, error) {

	// Get the list of buckets
	result, err := client.ListBuckets(nil)
	if err != nil {
		return nil, fmt.Errorf("error listing buckets: %v", err)
	}

	return result, nil
}

func GetS3Client(profileName string) (*s3.S3, error) {
	// Create a new session for minio
	// use profile credentials using endpoint url
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: profileName,
		Config: aws.Config{
			Region:           aws.String("us-west-1"),
			Endpoint:         aws.String("http://10.0.0.213:9000"),
			S3ForcePathStyle: aws.Bool(true),
			Credentials:      credentials.NewSharedCredentials("", profileName),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error creating session: %v", err)
	}

	return s3.New(sess), nil
}
