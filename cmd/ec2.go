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

	input := &ec2.DescribeInstancesInput{} // Res is DryRunOperation if everything is OK

	info, err := GetEc2Info(ctx, input, client)
	if err != nil {
		log.Fatal("Error Occured while retrieving Information", err)
	}

	for _, r := range info.Reservations {
		fmt.Printf("\n Reservations ID : %s ", *r.ReservationId)
		fmt.Printf("\n Owner ID : %s ", *r.OwnerId)
		// fmt.Printf("\n Requester ID : %s \n", *r.RequesterId) // Returing nil 

		for _, i := range r.Instances {

			fmt.Printf("\n Instances Launch Time: %s Instance ID: %s \n", i.LaunchTime, *i.InstanceId)
		}
	}

}

func init() {
	getCmd.AddCommand(ec2Cmd)
}
