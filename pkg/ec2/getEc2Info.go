package ec2

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/spf13/cobra"
)

type EC2Instance struct {
	InstanceName      string
	InstanceID        string
	InstanceType      string
	InstancePrivateIP string
}

func DescribeInstance(tagName string, tagValue []string, cmd *cobra.Command, args []string) []EC2Instance {

	var ec2Instance []EC2Instance
	input := &ec2.DescribeInstancesInput{}
	ctx, client := Ec2Client(cmd, args)
	info, err := client.DescribeInstances(ctx, input)
	if err != nil {
		log.Fatal("Error occured while retrieving information")
	}
	if tagName == "" {
		tagName = "Name"
	}
	for _, r := range info.Reservations {
		for _, i := range r.Instances {
			for _, j := range i.Tags {
				if *j.Key == tagName {
					if len(tagValue) != 0 {
						for _, value := range tagValue {
							if *j.Value == string(value) {
								ec2Instance = append(ec2Instance, EC2Instance{
									InstanceName:      aws.ToString(j.Value),
									InstanceID:        aws.ToString(i.InstanceId),
									InstanceType:      *aws.String(string(*&i.InstanceType)),
									InstancePrivateIP: aws.ToString(*&i.PrivateIpAddress),
								})
							}
						}
					} else {
						ec2Instance = append(ec2Instance, EC2Instance{
							InstanceName:      aws.ToString(j.Value),
							InstanceID:        aws.ToString(i.InstanceId),
							InstanceType:      *aws.String(string(*&i.InstanceType)),
							InstancePrivateIP: aws.ToString(*&i.PrivateIpAddress),
						})
					}
				}
			}
		}
	}
	return ec2Instance
}

func Ec2Client(cmd *cobra.Command, args []string) (context.Context, *ec2.Client) {

	region, err := cmd.Flags().GetString("region")
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal("Config Error Occured", err)
	}
	if region != "" {
		opts := func(o *ec2.Options) {
			o.Region = region
		}
		Ec2client := ec2.NewFromConfig(cfg, opts)
		return ctx, Ec2client
	}
	Ec2client := ec2.NewFromConfig(cfg)
	return ctx, Ec2client
}
