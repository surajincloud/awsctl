package EC2

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type EC2Instance struct {
	InstanceName      string
	InstanceID        string
	InstanceType      string
	InstancePrivateIP string
}

func GetEC2Instance() []EC2Instance {

	var ec2Instance []EC2Instance
	info, err := DescribeInstance()
	if err != nil {
		log.Fatal("Error Occured while retrieving Information", err)
	}
	instanceName := "-"
	for _, r := range info.Reservations {
		for _, i := range r.Instances {
			for _, j := range i.Tags {
				if *j.Key == "Name" {
					instanceName = aws.ToString(j.Value)
				}
			}
			ec2Instance = append(ec2Instance, EC2Instance{
				InstanceName:      instanceName,
				InstanceID:        aws.ToString(i.InstanceId),
				InstanceType:      *aws.String(string(*&i.InstanceType)),
				InstancePrivateIP: aws.ToString(*&i.PrivateIpAddress),
			})
		}
	}

	return ec2Instance
}

func GetInstanceWithKeyOnly(tagName *string) []EC2Instance {

	var ec2Instance []EC2Instance
	info, err := DescribeInstance()
	if err != nil {
		log.Fatal("Unable to get info from Aws Try Again")
	}
	for _, r := range info.Reservations {
		for _, i := range r.Instances {
			for _, j := range i.Tags {
				if *j.Key == *tagName {
					ec2Instance = append(ec2Instance, EC2Instance{
						InstanceName:      *aws.String(*j.Value),
						InstanceID:        *aws.String(*i.InstanceId),
						InstanceType:      *aws.String(string(i.InstanceType)),
						InstancePrivateIP: *aws.String(*i.PrivateIpAddress),
					})
				}
			}
		}
	}

	return ec2Instance
}

func GetInstanceWithKeyValue(tagName *string, tagValue []string) []EC2Instance {

	var ec2Instance []EC2Instance
	info, err := DescribeInstance()
	if err != nil {
		log.Fatal("Unable to get info from Aws Try Again")
	}
	for _, r := range info.Reservations {
		for _, i := range r.Instances {
			for _, j := range i.Tags {
				if *j.Key == *tagName {
					for _, value := range tagValue {
						if *j.Value == string(value) {
							ec2Instance = append(ec2Instance, EC2Instance{
								InstanceName:      *aws.String(*j.Value),
								InstanceID:        *aws.String(*i.InstanceId),
								InstanceType:      *aws.String(string(i.InstanceType)),
								InstancePrivateIP: *aws.String(*i.PrivateIpAddress),
							})
						}

					}
				}
			}
		}
	}

	return ec2Instance

}

func DescribeInstance() (*ec2.DescribeInstancesOutput, error) {

	input := &ec2.DescribeInstancesInput{}
	client, _ := Ec2Client()
	info, err := client.DescribeInstances(context.TODO(), input)
	return info, err
}

func Ec2Client() (*ec2.Client, context.Context) {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal("Config Error Occured", err)
	}
	Ec2client := ec2.NewFromConfig(cfg)
	return Ec2client, ctx
}
