package main

import (
	"context"
	"fmt"
	"log"
	"minioaws/internal/minioaws"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// Sleep for 1 second to allow the context to be cancelled

	client, err := minioaws.GetMinioS3Client(ctx, "minio", "p920")
	if err != nil {
		log.Fatalf("error getting minio client: %v", err)
	}

	// buckets1, err := minioaws.ListMinioBuckets(client)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(buckets1)

	// Create a new bucket
	err = minioaws.CreateMinioBucket(ctx, client, "test-bucket")
	if err != nil {
		fmt.Println(err)
	}

	buckets, err := minioaws.ListMinioBucketsWithContext(client, ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(buckets.Buckets)

}
