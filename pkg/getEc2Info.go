package pkg

import
(	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)


type EC2DesAPI interface {
	DescribeInstances(ctx context.Context,
		params *ec2.DescribeInstancesInput,
		optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error)
}

// Returns info about ec2 instance
func GetEc2Info(ctx context.Context, input *ec2.DescribeInstancesInput, api EC2DesAPI) (*ec2.DescribeInstancesOutput, error) {

	return api.DescribeInstances(ctx, input)

}