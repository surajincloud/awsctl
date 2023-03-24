package rds

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

type RDSInstance struct {
	DBInstanceID string
	Engine       string
	Status       string
	Version      string
	Endpoint     string
}

func GetRDSInstances() ([]RDSInstance, error) {
	var rdsInstances []RDSInstance

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, errors.New("failed to load AWS config: " + err.Error())
	}

	client := rds.NewFromConfig(cfg)

	input := &rds.DescribeDBInstancesInput{}
	result, err := client.DescribeDBInstances(context.Background(), input)
	if err != nil {
		return nil, errors.New("failed to retrieve RDS instaces : " + err.Error())
	}

	for _, dbInstance := range result.DBInstances {

		endpoint := "not found"
		if dbInstance.Endpoint != nil {
			endpoint = aws.ToString(dbInstance.Endpoint.Address)
		}
		rdsInstances = append(rdsInstances, RDSInstance{
			DBInstanceID: aws.ToString(dbInstance.DBInstanceIdentifier),
			Status:       aws.ToString(dbInstance.DBInstanceStatus),
			Endpoint:     endpoint,
			Engine:       aws.ToString(dbInstance.Engine),
			Version:      aws.ToString(dbInstance.EngineVersion),
		})
	}

	return rdsInstances, nil

}
