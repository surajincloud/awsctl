package pkg

import (
	"context"

	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GetBucketSize(Name string, ctx context.Context) (*s3.ListObjectsV2Output, error) {

	client := S3Clinet(ctx)
	size, err := client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(Name),
	})
	if err != nil {
		log.Fatal("No Object found \n", err)
	}

	return size, err

}

func S3Clinet(ctx context.Context) *s3.Client {

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal("Check Access Key and Secret", err)
	}
	client := s3.NewFromConfig(cfg)

	
	return client
}
