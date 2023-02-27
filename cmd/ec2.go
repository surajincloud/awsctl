/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/spf13/cobra"
	"github.com/surajincloud/awsctl/pkg"
)

// ec2Cmd represents the ec2 command
var ec2Cmd = &cobra.Command{
	Use:   "ec2",
	Short: "Print ec2 related information",
	Long: `For example,
		awsctl get ec2`,
	Run: getEC2command,
}

func getEC2command(cmd *cobra.Command, args []string) {

	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal("Config Error Occured", err)
	}

	client := ec2.NewFromConfig(cfg)
	

	input := &ec2.DescribeInstancesInput{}

	info, err := pkg.GetEc2Info(ctx, input, client)
	if err != nil {
		log.Fatal("Error Occured while retrieving Information", err)
	}

	//Instance Details
	instanceName:="-"
	fmt.Println("NAME \t INSTANCE_ID \t    INSTANCE_TYPE \t PRIVATE IP")
	for _, r := range info.Reservations {
		
		for _, i := range r.Instances {
			for _,j:= range i.Tags{
				if *j.Key == "Name"{
					instanceName=aws.ToString(j.Value)
				}
			}

			fmt.Println(instanceName,"\t",aws.ToString(i.InstanceId),"\t",*&i.InstanceType,"\t",aws.ToString(*&i.PrivateIpAddress))
		}
	}

}

func init() {
	getCmd.AddCommand(ec2Cmd)
}
