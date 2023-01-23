package vpc

import (
	ec2 "github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type VPC struct {
	scope *constructs.Construct
	vpc   ec2.Vpc
}

func NewVPC(scope constructs.Construct) *VPC {
	return &VPC{
		scope: &scope,
	}
}

func (v *VPC) MakeID() *string {
	return jsii.String("VPC")
}

func (v *VPC) Make() {
	v.vpc = ec2.NewVpc(*v.scope, v.MakeID(), &ec2.VpcProps{
		IpAddresses: ec2.IpAddresses_Cidr(jsii.String("10.0.0.0/16")),
		SubnetConfiguration: &[]*ec2.SubnetConfiguration{
			{
				CidrMask:   jsii.Number(24),
				Name:       jsii.String("Public"),
				SubnetType: ec2.SubnetType_PUBLIC,
			},
		},
	})
}

func (v *VPC) GetVPC() ec2.Vpc {
	return v.vpc
}
