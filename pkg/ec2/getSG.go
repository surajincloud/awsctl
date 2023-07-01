package ec2

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/spf13/cobra"
)

type SecurityGroup struct {
	SGName        string
	SGId          string
	SGDescription string
}

func DescribeSecurityGroup(cmd *cobra.Command, args []string) ([]SecurityGroup, error) {

	var securityGrp []SecurityGroup

	ctx, client := Ec2Client(cmd, args)
	input := &ec2.DescribeSecurityGroupsInput{}
	info, err := client.DescribeSecurityGroups(ctx, input)

	for _, i := range info.SecurityGroups {
		securityGrp = append(securityGrp, SecurityGroup{
			SGName:        aws.ToString(i.GroupName),
			SGId:          aws.ToString(i.GroupId),
			SGDescription: aws.ToString(i.Description),
		})
	}

	return securityGrp, err

}

