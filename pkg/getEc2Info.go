package pkg

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)


type EC2Instance struct {

	InstanceName 		string
	InstanceID	 		string
	InstanceType 		string
	InstancePrivateIP 	string

}



func GetEC2Instance()([]EC2Instance){
	var ec2Instance [] EC2Instance

	ctx := context.TODO()
	cfg,err:=config.LoadDefaultConfig(ctx)

	if err!=nil{
		log.Fatal("Config Error Occured",err)
	}

	client := ec2.NewFromConfig(cfg)
	
	input := &ec2.DescribeInstancesInput{}
	info, err := client.DescribeInstances(ctx,input)
	if err != nil {
		log.Fatal("Error Occured while retrieving Information", err)
	}
	
	instanceName:="-"
	for _, r := range info.Reservations {
		for _, i := range r.Instances {
			for _,j:= range i.Tags{
				if *j.Key == "Name"{
					instanceName=aws.ToString(j.Value)
				}
			}
			ec2Instance = append(ec2Instance, EC2Instance{
				InstanceName: 		instanceName,
				InstanceID:			aws.ToString(i.InstanceId) ,
				InstanceType: 		*aws.String(string(*&i.InstanceType)),
				InstancePrivateIP: 	aws.ToString(*&i.PrivateIpAddress),
			})
		}
	}

	return ec2Instance

}
