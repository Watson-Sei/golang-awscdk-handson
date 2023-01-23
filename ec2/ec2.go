package ec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type EC2 struct {
	scope *constructs.Construct

	vpc awsec2.Vpc
}

func NewEC2(scope constructs.Construct, vps awsec2.Vpc) *EC2 {
	return &EC2{
		scope: &scope,
		vpc:   vps,
	}
}

func (e *EC2) MakeID() *string {
	return jsii.String("EC2")
}

func (e *EC2) Make() {
	awsec2.NewInstance(*e.scope, e.MakeID(), &awsec2.InstanceProps{
		Vpc:          e.vpc,
		InstanceType: awsec2.InstanceType_Of(awsec2.InstanceClass_T3, awsec2.InstanceSize_MICRO),
		UserData: awsec2.UserData_Custom(jsii.String(`#!/bin/bash
			yum update -y
			yum install -y docker
			service docker start
			usermod -a -G docker ec2-user
		`)),
	})
}
