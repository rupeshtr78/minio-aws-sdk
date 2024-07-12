package minioaws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// ListMinioBuckets lists the Minio buckets available in the S3 client.
//
// client is the S3 client to use for the operation.
//
// Returns the list of Minio buckets, or an error if the operation fails.
func ListMinioBuckets(client *s3.S3) (*s3.ListBucketsOutput, error) {

	// Get the list of buckets
	result, err := client.ListBuckets(nil)
	if err != nil {
		return nil, fmt.Errorf("error listing buckets: %v", err)
	}

	return result, nil

}

func ListMinioBucketsWithContext(client *s3.S3, ctx aws.Context) (*s3.ListBucketsOutput, error) {

	// Get the list of buckets
	result, err := client.ListBucketsWithContext(ctx, &s3.ListBucketsInput{})
	if err != nil || ctx.Err() != nil {
		return nil, fmt.Errorf("error listing buckets: %v", err)
	}

	return result, nil

}

// CreateMinioBucket creates a new Minio bucket with the specified name.
//
// client is the S3 client to use for the operation.
// bucketName is the name of the Minio bucket to create.
//
// Returns an error if the bucket creation fails.
func CreateMinioBucket(ctx context.Context, client *s3.S3, bucketName string) error {

	// Check to see if the bucket already exists
	_, err := client.HeadBucketWithContext(ctx, &s3.HeadBucketInput{
		Bucket: &bucketName,
	})
	if err == nil {
		return nil
	}

	// Create a new bucket
	_, err = client.CreateBucketWithContext(ctx, &s3.CreateBucketInput{
		Bucket: &bucketName,
	})
	if err != nil {
		return fmt.Errorf("error creating bucket: %v", err)
	}

	return nil

}

// DeleteMinioBucket deletes the specified Minio bucket.
//
// client is the S3 client to use for the operation.
// bucketName is the name of the Minio bucket to delete.
//
// Returns an error if the bucket deletion fails.
func DeleteMinioBucket(client *s3.S3, bucketName string) error {
	// Delete the bucket
	_, err := client.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: &bucketName,
	})
	if err != nil {
		return fmt.Errorf("error deleting bucket: %v", err)
	}
	return nil
}

// ListMinioObjects lists the objects in the specified Minio bucket.
//
// client is the S3 client to use for the operation.
// bucketName is the name of the Minio bucket to list objects from.
//
// Returns the list of objects in the bucket, or an error if the operation fails.
func ListMinioObjects(client *s3.S3, bucketName string) (*s3.ListObjectsOutput, error) {
	// Get the list of objects
	result, err := client.ListObjects(&s3.ListObjectsInput{
		Bucket: &bucketName,
	})
	if err != nil {
		return nil, fmt.Errorf("error listing objects: %v", err)
	}
	return result, nil
}
