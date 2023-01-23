package main

import (
	"fmt"
	"os"

	"github.com/Watson-Sei/golang-awscdk-handson/ec2"
	"github.com/Watson-Sei/golang-awscdk-handson/vpc"
	awscdk "github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type GolangAwscdkHandsonStackProps struct {
	awscdk.StackProps
}

func NewGolangAwscdkHandsonStack(scope constructs.Construct, id string, props *GolangAwscdkHandsonStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}

	stackName := fmt.Sprintf("GolangAwscdkHandsonStack-%s", id)
	stack := awscdk.NewStack(scope, &stackName, &sprops)

	v := vpc.NewVPC(stack)
	v.Make()

	ec := ec2.NewEC2(stack, v.GetVPC())
	ec.Make()
	// The code that defines your stack goes here

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	name := os.Getenv("STACK_NAME")
	if name == "" {
		name = "dev"
	}

	NewGolangAwscdkHandsonStack(app, name, &GolangAwscdkHandsonStackProps{
		awscdk.StackProps{
			Env: &awscdk.Environment{
				Region: jsii.String("ap-northeast-1"),
			},
		},
	})

	app.Synth(nil)
}
