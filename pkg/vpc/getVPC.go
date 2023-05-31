package vpc

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type Vpc struct {
	Name     *string
	State    *string
	Ipv4Cidr *string
	Default  bool
	VpcID    *string
}

func GetVPC() ([]Vpc, error) {

	var vpcList []Vpc

	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	client := ec2.NewFromConfig(cfg)

	info, err := client.DescribeVpcs(ctx, &ec2.DescribeVpcsInput{})
	if err != nil {
		return nil, err
	}
	for _, i := range info.Vpcs {
		for _, j := range i.Tags {
			if *j.Key == "Name" {
				vpcList = append(vpcList, Vpc{
					Name:     aws.String(*j.Value),
					State:    aws.String(string(i.State)),
					Ipv4Cidr: aws.String(*i.CidrBlock),
					Default:  *i.IsDefault,
					VpcID:    aws.String(*i.VpcId),
				})
			}

		}
	}

	return vpcList, err

}
