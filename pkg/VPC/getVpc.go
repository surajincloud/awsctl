package vpc

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/surajincloud/awsctl/pkg"
)


func GetVpc(){

	
	client,ctx:= pkg.Ec2Client()

	input:= &ec2.DescribeVpcsInput{}

	info,err:= client.DescribeVpcs(ctx,input)
	_=err
	
	for _,i:=range info.Vpcs{
		fmt.Println(i.CidrBlock)
	for _,j:=range i.Tags{
		if(*j.Key=="Name"){
			fmt.Println(j.Key)
		}
	}
	}

}
