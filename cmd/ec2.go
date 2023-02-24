/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/spf13/cobra"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// ec2Cmd represents the ec2 command
var ec2Cmd = &cobra.Command{
	Use:   "ec2",
	Short: "Print ec2 related information",
	Long: `For example,
		awsctl get ec2`,
	Run: ec2FunDescribe,
}


type EC2DesAPI interface {
	DescribeInstances(ctx context.Context,
		params *ec2.DescribeInstancesInput,
		optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error)
}


// Returns info about ec2 instance 

func GetEc2Info(ctx context.Context,input *ec2.DescribeInstancesInput,api EC2DesAPI) (*ec2.DescribeInstancesOutput,error) {

	return api.DescribeInstances(ctx,input)

}

func ec2FunDescribe (cmd * cobra.Command,args [] string){

	ctx:= context.TODO() // Unclear Context
	cfg,err:=config.LoadDefaultConfig(ctx)
	if err!=nil{
		log.Fatal("Config Error Occured",err)
	}

	client:= ec2.NewFromConfig(cfg)

	input:=&ec2.DescribeInstancesInput{} // Res is DryRunOperation if everything is OK

	info,err:=GetEc2Info(ctx,input,client)
	if err!= nil{
		log.Fatal("Error Occured while retrieving Information",err)
	}

	for _ , r:= range (info.Reservations){
		fmt.Printf("\n Reservations ID : %s ",*r.ReservationId)
		fmt.Printf("\n Reservations ID : %s ",*r.OwnerId)
		fmt.Printf("\n Reservations ID : %s \n",*r.RequesterId)
		
		for _ ,i:= range r.Instances{

			fmt.Printf("\n Instances ID: %s %s",i.LaunchTime,*i.InstanceId)
		}
	}



} 

func init() {
	getCmd.AddCommand(ec2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ec2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ec2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
