/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/spf13/cobra"
)

// ec2Cmd represents the ec2 command
var ec2Cmd = &cobra.Command{
	Use:   "ec2",
	Short: "Print ec2 related information",
	Long: `For example,
		awsctl get ec2`,
	Run: getEC2command,
}

type EC2DesAPI interface {
	DescribeInstances(ctx context.Context,
		params *ec2.DescribeInstancesInput,
		optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error)
}

// Returns info about ec2 instance
func GetEc2Info(ctx context.Context, input *ec2.DescribeInstancesInput, api EC2DesAPI) (*ec2.DescribeInstancesOutput, error) {

	return api.DescribeInstances(ctx, input)

}

func getEC2command(cmd *cobra.Command, args []string) {

	ctx := context.TODO() // Unclear Context
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal("Config Error Occured", err)
	}

	client := ec2.NewFromConfig(cfg)
	

	input := &ec2.DescribeInstancesInput{}

	info, err := GetEc2Info(ctx, input, client)
	if err != nil {
		log.Fatal("Error Occured while retrieving Information", err)
	}

	//Print Instance Details
	instanceName:=""
	fmt.Println("NAME \t INSTANCE_ID \t    INSTANCE_TYPE \t PRIVATE IP")
	for _, r := range info.Reservations {
		
		for _, i := range r.Instances {
			for _,j:= range i.Tags{
				if *j.Key == "Name"{
					instanceName=*j.Value
				}
			}
			fmt.Println(instanceName,"\t",*i.InstanceId,"\t",*&i.InstanceType,"\t",*i.PrivateIpAddress)	
		}
	}

}

func init() {
	getCmd.AddCommand(ec2Cmd)
}
