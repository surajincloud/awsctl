package s3API

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Bucket struct {
	Name      string
	CreatedAt *time.Time
	Size      int
}

// Returns Info in the form []Bucket
func GetBucket() []Bucket {

	var s3Bucket []Bucket

	ctx := context.TODO()
	client := S3Client(ctx)

	s3List, err := client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		log.Fatal("Unable to get S3 List", err)
	}
	for _, i := range s3List.Buckets {
		size, _ := GetBucketSize(*i.Name, ctx)
		for _, j := range size.Contents {
			sizeByKB := j.Size / 1024
			s3Bucket = append(s3Bucket, Bucket{
				Name:      *aws.String(*i.Name),
				CreatedAt: i.CreationDate,
				Size:      int(sizeByKB),
			})
		}
	}

	return s3Bucket

}

// Returns Bucket Size in Bytes
func GetBucketSize(Name string, ctx context.Context) (*s3.ListObjectsV2Output, error) {

	client := S3Client(ctx)
	size, err := client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(Name),
	})
	if err != nil {
		log.Fatal("No Object found \n", err)
	}
	return size, nil
}

// returns a s3client
func S3Client(ctx context.Context) *s3.Client {

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal("Check Access Key and Secret", err)
	}
	client := s3.NewFromConfig(cfg)

	return client
}
