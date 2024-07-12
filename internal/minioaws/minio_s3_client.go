package minioaws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// GetMinioS3Client creates a new AWS S3 client using the specified profile name.
// The client is configured to use the Minio S3 endpoint at the specified URL,
// with the specified AWS region and shared credentials.
func GetMinioS3Client(ctx context.Context, profileName string, bucket string) (*s3.S3, error) {

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

	// Create a new S3 service client
	client := s3.New(sess)

	// verify the connection
	_, err = client.HeadBucketWithContext(ctx, &s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}
